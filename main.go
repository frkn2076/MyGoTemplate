package main

import (
	_ "app/MyGoTemplate/db"
	_ "app/MyGoTemplate/logger"
	"app/MyGoTemplate/router"
	"app/MyGoTemplate/socket"
	"app/MyGoTemplate/resource"

	"fmt"
)

func main() {
	go socket.H.Run()
	fmt.Println(resource.GetValue("furkan"))
	fmt.Println(resource.GetValue("batu"))
	
	r := router.SetupRouter()
	r.Run(":8080")
}
