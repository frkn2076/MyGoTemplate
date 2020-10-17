package logger

import (
	"log"
	"os"
	"time"
	"fmt"
)

var info *log.Logger = InitLogger("InfoLog")
var error *log.Logger = InitLogger("ErrorLog")

func InitLogger(folderName string) *log.Logger {
	dt := time.Now()
	today := dt.Format("01-Jan-2006")
	fileName := fmt.Sprintf("logger/%s/%swarning.log", folderName, today)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

	logger := log.New(file, "prefix", log.LstdFlags)
	logger.Println(fmt.Sprintf("%s has created", folderName))
	return logger
}

func ErrorLog(logText string){
	error.Println(logText)
}

func InfoLog(logText string){
	info.Println(logText)
}

