package middleware

import(
	"app/MyGoTemplate/logger"

	"github.com/gin-gonic/gin"

	"bytes"
	"io/ioutil"
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
		logger.ServiceLog("Request: ", c.Request.RequestURI, " ", string(body))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		//#endregion RequestLog


		//#region ResponseLog

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()
		responseBody := bodyLogWriter.body.String()
		logger.ServiceLog("Response: ", string(responseBody))

		//#endregion ResponseLog
    }
}