package db

import (
	"database/sql"
	"io/ioutil"
	"strings"

	"app/MyGoTemplate/logger"
)

//Additional scripts like CRUD, trigger, SP etc.
func InitScripts(db *sql.DB) {

	initScript, err := ioutil.ReadFile("db/init.sql")
	if err != nil {
		logger.ErrorLog("An error occured while reading init.sql file - InitScripts - dbInitializer.go ", err.Error())
	}

	transaction, err := db.Begin()
	if err != nil {
		logger.ErrorLog("An error occured while beginning transaction - InitScripts - dbInitializer.go ", err.Error())
	} else {
		logger.TransactionLog("Transaction began")
	}

	defer func(){
		err := transaction.Rollback()
		if err != nil {
			logger.ErrorLog("An error occured while rollbacking transaction - InitScripts - dbInitializer.go ", err.Error())
		} else {
			logger.TransactionLog("Transaction rollback")
		}
	}()

	for _, statement := range strings.Split(string(initScript), ";") {
		statement := strings.TrimSpace(statement)
		if statement == "" {
			continue
		}
		if _, err := transaction.Exec(statement); err != nil {
			logger.ErrorLog("An error occured while executing statements - InitScripts - dbInitializer.go ", err.Error())
		} else {
			logger.TransactionLog(statement)
		}
	}

	err = transaction.Commit()
	if err != nil {
		logger.ErrorLog("An error occured while committing transaction - InitScripts - dbInitializer.go ", err.Error())
	} else {
		logger.TransactionLog("Transaction committed")
	}
}