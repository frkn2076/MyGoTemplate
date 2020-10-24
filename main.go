package main

import (
	_ "app/MyGoTemplate/db"
	_ "app/MyGoTemplate/logger"
	"app/MyGoTemplate/router"
	"app/MyGoTemplate/socket"

)

func main() {
	go socket.H.Run()

	r := router.SetupRouter()
	r.Run(":8080")
}
