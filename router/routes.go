package router

import (
	"app/MyGoTemplate/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	grp1 := router.Group("/login")
	{
		grp1.POST("login", controllers.Login)
		grp1.POST("register", controllers.Register)
	}

	return router
}
