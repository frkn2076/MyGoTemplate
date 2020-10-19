package router

import (
	"app/MyGoTemplate/middleware"
	"app/MyGoTemplate/controllers"

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

	// grp2 := router.Group("/socket")
	// {
	// 	grp2.POST("send", message.SendMessage)
	// 	grp2.GET("show", message.ShowMessage)
	// }
	
	return router
}


