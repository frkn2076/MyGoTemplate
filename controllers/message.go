package controllers

// import (
// "app/MyGoTemplate/logger"
// "app/MyGoTemplate/socket"

// 	"bufio"
// 	"io/ioutil"
// 	"net"

// 	"github.com/gin-gonic/gin"
//  )

// const (
//     connHost = "localhost"
//     connPort = "8080"
//     connType = "tcp"
// )

// type MessageController struct{}

// func (u *MessageController) SendMessage(c *gin.Context) {
// 	request, _ := ioutil.ReadAll(c.Request.Body)

// 	conn, err := net.Dial(connType, connHost+":"+connPort)
// 	if err != nil {
// 		logger.ErrorLog(err)
// 	}
// 	conn.Write([]byte(string(request) + "\n"))

// 	c.JSON(200,"OK")
// }

// func (u *MessageController) ShowMessage(c *gin.Context) {
// 	// l, _ := net.Listen("tcp", "127.0.0.1:8080")
// 	conn, err := socket.Listener.Accept()
//     if err != nil {
// 		logger.ErrorLog("Socket connecting error: ", err.Error())
//         c.JSON(400,"Error")
// 	}

// 	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
//     if err != nil {
// 		logger.InfoLog("Client left.")
//         conn.Close()
//         return
// 	}
	
// 	c.JSON(200, string(buffer[:len(buffer)-1]))
// }