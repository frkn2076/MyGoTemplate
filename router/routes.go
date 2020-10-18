package router

import (
	"app/MyGoTemplate/controllers"

	"github.com/gin-gonic/gin"
	"app/MyGoTemplate/middleware"

	"os"
	"fmt"
	"io"
)

func SetupRouter() *gin.Engine {
	user := new(controllers.UserController)
	f, err := os.Create("logger/ServiceLog/log.log")
	if err != nil {
		fmt.Println("Open Log File Failed", err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	
	router.Use(middleware.ServiceLogMiddleware())
	
	grp1 := router.Group("/login")
	{
		grp1.POST("login", user.Login)
		// grp1.POST("register", controllers.Register)
	}
	
	return router
}


