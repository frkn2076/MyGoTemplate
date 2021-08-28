package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var info *log.Logger = initLogger("InfoLog")
var error *log.Logger = initLogger("ErrorLog")
var service *log.Logger = initLogger("ServiceLog")
var chat *log.Logger = initLogger("ChatLog")
var transaction *log.Logger = initLogger("TransactionLog")

func ErrorLog(logText ...interface{}) {
	error.Println(logText)
}

func InfoLog(logText ...interface{}) {
	info.Println(logText)
}

func ServiceLog(logText ...interface{}) {
	service.Println(logText)
}

func ChatLog(logText ...interface{}) {
	chat.Println(logText)
}

func TransactionLog(logText ...interface{}) {
	transaction.Println(logText)
}


//#region helper/unexporteds

func initLogger(folderName string) *log.Logger {
	dt := time.Now()
	today := dt.Format("02-Jan-2006")
	fileName := fmt.Sprintf(os.Getenv("LoggerFilePath"), folderName, today)
	
	//check log file created before
	_, err := os.Stat(fileName)
	fileNotExist := os.IsNotExist(err)
	
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "prefix: ", log.LstdFlags)
	if(fileNotExist){
		logger.Println(fmt.Sprintf("%s has created", folderName))
	}
	return logger
}

//#endregion helper/unexporteds