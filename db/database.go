package db

import (
    "database/sql"
	"app/MyGoTemplate/logger"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
    var err error
    
	DB, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db")

    if err != nil {
		logger.ErrorLog("An error occured while database connection is establishing")
        return err
    }
    if err = DB.Ping(); err != nil {
        return err
	}
	logger.InfoLog("Database connection is opened")
    Init()
    logger.InfoLog("Init sql script has runned")
    return nil
}
