package main

import (
	_ "app/MyGoTemplate/db"
	_ "app/MyGoTemplate/logger"
	"app/MyGoTemplate/socket"
	"app/MyGoTemplate/router"
)


func main() {

	r := router.SetupRouter()
	socket.Start()
	r.Run(":8080")
}



