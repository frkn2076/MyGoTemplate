package main

import (
	_ "app/MyGoTemplate/db"
	_ "app/MyGoTemplate/logger"
	"app/MyGoTemplate/router"

)


func main() {

	r := router.SetupRouter()
	r.Run()
}



