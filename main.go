package main

import (
	_ "app/MyGoTemplate/db"
	_ "app/MyGoTemplate/logger"
	"app/MyGoTemplate/router"

	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()

	r := router.SetupRouter()

	r.LoadHTMLFiles("index.html")

	r.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	r.Run(":8080")

}
