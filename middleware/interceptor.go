package middleware

import(
	"github.com/gin-gonic/gin"

	"fmt"
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
		body, _ := ioutil.ReadAll(c.Request.Body)
		println(string(body))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()
		responseBody := bodyLogWriter.body.String()
		fmt.Println(responseBody)
    }
}