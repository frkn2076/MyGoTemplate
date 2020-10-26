package middleware

import(
	"bytes"
	"io/ioutil"
	"net/http"
	"fmt"

	"app/MyGoTemplate/logger"
	"app/MyGoTemplate/controllers/models/response"
	"app/MyGoTemplate/cache"
	s "app/MyGoTemplate/session"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}
func (w bodyLogWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

func ServiceLogMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		//#region RequestLog

		body, _ := ioutil.ReadAll(c.Request.Body)
		logger.ServiceLog("Request: ", c.Request.RequestURI, " IP: ", c.ClientIP(), " " , string(body))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		//#endregion


		//#region ResponseLog

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		
		c.Writer = bodyLogWriter
		
		c.Next() // < the rest of handlers in the chain are executed here!
		
		errorHandler(c)

		responseBody := bodyLogWriter.body.String()
		logger.ServiceLog("Response: ", string(responseBody))

		//#endregion
    }
}

func errorHandler(c *gin.Context) {
	lang, err := s.SessionGet(c, "language")
	if err != nil {
		logger.ErrorLog("An error occured while using ", err.Error())
	}
	language := fmt.Sprintf("%v", lang)
	
	if(len(c.Errors) > 0){
		key := c.Errors[0].Error() + language
		errorMessage := cache.Get(key)
		c.AbortWithStatusJSON(http.StatusInternalServerError, &response.BaseResponse{IsSuccess: false,	Message: errorMessage})
		return
	} else if(c.Writer.Status() < http.StatusOK || c.Writer.Status() > http.StatusIMUsed) {
		c.AbortWithStatusJSON(c.Writer.Status(), &response.BaseResponse{IsSuccess: false,	Message: cache.Get("GlobalErrorMessage" + language)})
		return
	}
}