package db

import (
	"os"
	"time"
	"context"
	"database/sql"

	"app/MyGoTemplate/logger"
	"app/MyGoTemplate/db/entities"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *sql.DB
var GormDB *gorm.DB

func init(){
	DB = initDB()
	GormDB = initGormDB()
}

//#region Helper

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db?parseTime=true");
	
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
	
	return db
}

func initGormDB() *gorm.DB{

	//Initialize gorm with existing db connection
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	  }), &gorm.Config{
		SkipDefaultTransaction: true, //skipped to start own transactions in repositories to supply Unit of Work.
	  })
	
	if err != nil {
		logger.ErrorLog("An error occured while gorm driver is establishing: ", err.Error())
	}

	gormDB.AutoMigrate(&entities.User{}, &entities.Login{}, &entities.Localization{})

	InitScripts(DB)
	logger.InfoLog("Init sql script has runned")

	return gormDB
}


//#endregion