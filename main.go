package main

import (
	"app/MyGoTemplate/infra/environments"
	"app/MyGoTemplate/infra/db"
	_ "app/MyGoTemplate/infra/logger"
	_ "app/MyGoTemplate/infra/resource"
	"app/MyGoTemplate/router"
	"app/MyGoTemplate/socket"
)

func main() {
	environments.Load()

	db.ConnectDatabases()
	db.MigrateTables()
	db.InitScripts()

	go socket.H.Run()

	r := router.SetupRouter()
	r.Run(":8080")
}
