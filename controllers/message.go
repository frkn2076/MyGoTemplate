package controllers

import(
	"app/MyGoTemplate/logger"
	"app/MyGoTemplate/socket"

	"net"
	"bufio"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

const (
    connHost = "localhost"
    connPort = "8081"
    connType = "tcp"
)

type MessageController struct{}

func (u *MessageController) SendMessage(c *gin.Context) {
	request, _ := ioutil.ReadAll(c.Request.Body)

	conn, _ := net.Dial(connType, connHost+":"+connPort)
	conn.Write([]byte(string(request) + "\n"))

	c.JSON(200,"OK")
}

func (u *MessageController) ShowMessage(c *gin.Context) {
	conn, err := socket.Listener.Accept()
    if err != nil {
		logger.ErrorLog("Socket connecting error: ", err.Error())
        c.JSON(400,"Error")
	}

	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
    if err != nil {
		logger.InfoLog("Client left.")
        conn.Close()
        return
	}
	
	c.JSON(200, string(buffer[:len(buffer)-1]))
}