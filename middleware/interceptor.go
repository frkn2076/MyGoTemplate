package middleware

import(
	"bytes"
	"io/ioutil"
	"net/http"

	"app/MyGoTemplate/logger"
	"app/MyGoTemplate/controllers/models/response"

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
		responseBody := bodyLogWriter.body.String()
		logger.ServiceLog("Response: ", string(responseBody))

		//#endregion


		//#region ExceptionHandle

		// base := r.BaseResponse{
		// 	IsSuccess: false,
		// 	Message: "asdasd",
		// }

		if(c.Writer.Status() == 500){
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.SingletonBaseResponseInstance)
			return
		}

		//#endregion
    }
}