package logger

import (
	"fmt"
	"log"
	"os"
	"time"
	"path"

	"gorm.io/gorm/logger"
)

var info *log.Logger = initLogger("InfoLog")
var error *log.Logger = initLogger("ErrorLog")
var service *log.Logger = initLogger("ServiceLog")
var chat *log.Logger = initLogger("ChatLog")
var transaction *log.Logger = initLogger("TransactionLog")
var QueryLogger logger.Interface = initQueryLogger("TransactionLog")

func ErrorLog(logText ...interface{}) {
	error.Println(logText...)
}

func InfoLog(logText ...interface{}) {
	info.Println(logText...)
}

func ServiceLog(logText ...interface{}) {
	service.Println(logText...)
}

func ChatLog(logText ...interface{}) {
	chat.Println(logText...)
}

func TransactionLog(logText ...interface{}) {
	transaction.Println(logText...)
}

func initLogger(folderName string) *log.Logger {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	dt := time.Now()
	today := dt.Format("02-Jan-2006")
	envLogPath := os.Getenv("LoggerFilePath")
	folderPath := path.Join(wd, envLogPath, folderName)

	//check folder created before
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0700)
	}

	fileName := path.Join(folderPath, today + ".log")

	//check log file created before
	_, err = os.Stat(fileName)
	fileNotExist := os.IsNotExist(err)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "* ", log.LstdFlags)
	if fileNotExist {
		logger.Println(fmt.Sprintf("%s has created", folderName))
	}
	return logger
}


func initQueryLogger(folderName string) logger.Interface {
	transactionLogger := initLogger(folderName)
	newLogger := logger.New(
		transactionLogger,
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)
	return newLogger
}
