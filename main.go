package main

import (
	_ "app/MyGoTemplate/db"
	_ "app/MyGoTemplate/logger"
	"app/MyGoTemplate/router"
	"app/MyGoTemplate/middleware"

)

func main() {

	r := router.SetupRouter()
	r.Use(middleware.MiddlewareServiceLogger)
	r.Run()

}
