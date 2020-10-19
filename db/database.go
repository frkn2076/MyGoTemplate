package db

import (
	"app/MyGoTemplate/logger"
	"app/MyGoTemplate/db/entities"
	
	"gorm.io/driver/mysql"
	"database/sql"
	"gorm.io/gorm"
	"context"
	"time"
	"os"
)

var DB *sql.DB = initDB()
var GormDB *gorm.DB = initGormDB()

//#region helper/unexporteds

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db");
	
	if err != nil {
		logger.ErrorLog("An error occured while database connection is establishing ", err.Error())
		os.Exit(0)
	}

	//Ping for 2 seconds
	ctx, cancelfunc := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancelfunc()
    if err = db.PingContext(ctx); err != nil {
		logger.ErrorLog("An error occured while ping: ", err.Error())
	}

	logger.InfoLog("Database connection is opened")
	InitScripts(db)
	logger.InfoLog("Init sql script has runned")
	return db
}

func initGormDB() *gorm.DB{

	//Initialize gorm with existing db connection
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	  }), &gorm.Config{})
	
	if err != nil {
		logger.ErrorLog("An error occured while gorm driver is establishing: ", err.Error())
	}

	//Migrations
	gormDB.AutoMigrate(&entities.User{}, &entities.Login{})

	return gormDB
}

//#endregion helper/unexporteds