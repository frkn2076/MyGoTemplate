package middleware

import(
	"app/MyGoTemplate/logger"

	"bytes"
    "io/ioutil"
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

func MiddlewareServiceLogger(c *gin.Context) {

    jsonData, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        // Handle error
    }
    rdr1 := ioutil.NopCloser(bytes.NewBuffer(jsonData))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(jsonData))
    jsonData2 := new(bytes.Buffer)
	jsonData2.ReadFrom(rdr1)

	s := jsonData2.String()
    logger.ServiceLog(s)

    c.Request.Body = rdr2
    c.Next()

    blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
    c.Writer = blw
    c.Next()
	// statusCode := c.Writer.Status()
	logger.ServiceLog(blw.body.String())
}