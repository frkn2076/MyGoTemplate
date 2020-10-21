package router

import (
	"app/MyGoTemplate/middleware"
	"app/MyGoTemplate/controllers"
	"app/MyGoTemplate/socket"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ServiceLogMiddleware())

	user := new(controllers.LoginController)
	// message := new(controllers.MessageController)
	
	grp1 := router.Group("/login")
	{
		grp1.POST("login", user.Login)
		grp1.POST("register", user.Register)
	}

	//User interface to demo
	router.LoadHTMLFiles("socket/index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		socket.ServeWs(c.Writer, c.Request, roomId)
	})
	
	return router
}


