package main

import (
	_ "app/MyGoTemplate/infra/environments"
	_ "app/MyGoTemplate/infra/db"
	_ "app/MyGoTemplate/infra/logger"
	_ "app/MyGoTemplate/infra/resource"
	"app/MyGoTemplate/router"
	"app/MyGoTemplate/socket"
)

func main() {
	go socket.H.Run()

	r := router.SetupRouter()
	r.Run(":8080")
}
