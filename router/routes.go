package router

import (
	"os"

	"app/MyGoTemplate/api/heartbeat"
	"app/MyGoTemplate/api/login"
	"app/MyGoTemplate/middleware"
	"app/MyGoTemplate/socket"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ServiceLogMiddleware())

	loginController := login.NewController()
	heartbeatController := heartbeat.NewController()
	// message := new(controllers.MessageController)

	grp1 := router.Group("/login")
	{
		grp1.POST("login", loginController.Login)
		grp1.POST("register", loginController.Register)
		grp1.POST("validation", loginController.RegisterValidation)
	}

	//User interface to demonstrate
	router.LoadHTMLFiles(os.Getenv("SocketDemonstrateHTMLPath"))

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		socket.ServeWs(c.Writer, c.Request, roomId)
	})

	grp2 := router.Group("/manager")
	{
		grp2.GET("/reports", heartbeatController.GetAllReports)
		grp2.GET("/clearCache", heartbeatController.ClearCache)
	}

	return router
}
