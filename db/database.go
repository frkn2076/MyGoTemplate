package db

import (
	"app/MyGoTemplate/logger"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB = initDB()

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db")

	if err != nil {
		logger.ErrorLog("An error occured while database connection is establishing")
		logger.ErrorLog(err.Error())
	}
	if err = db.Ping(); err != nil {
		logger.ErrorLog("An error occured while ping")
		logger.ErrorLog(err.Error())
	}
	logger.InfoLog("Database connection is opened")
	InitScripts(db)
	logger.InfoLog("Init sql script has runned")
	return db
}
