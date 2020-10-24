package db

import (
	"database/sql"
	"encoding/xml"

	"app/MyGoTemplate/helper"
	"app/MyGoTemplate/logger"
)

//Additional scripts like trigger, SP etc.
func InitScripts(db *sql.DB) {
	var root Root

	helper.LoadModel("db/localizationFeeder.xml", &root)

	languages := root.Languages

	transaction, err := db.Begin()
	if err != nil {
		logger.ErrorLog("An error occured while beginning transaction - InitScripts - dbInitializer.go ", err.Error())
	}
	defer transaction.Rollback()

	template := "INSERT IGNORE INTO Localization(resource, message, language) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language)"
	statement, err := transaction.Prepare(template)
	if err != nil {
		logger.ErrorLog("An error occured while preparing statements - InitScripts - dbInitializer.go ", err.Error())
	}
	defer statement.Close()

	for _, language := range languages {
		for _, resource := range language.Resources {
			_, err = statement.Exec(resource.Key, resource.Message, language.Code)
			if err != nil {
				logger.ErrorLog("An error occured while executing statements - InitScripts - dbInitializer.go ", err.Error())
			} else {
				logger.TransactionLog(template, resource.Key, resource.Message, language.Code)
			}
		}
	}

	err = transaction.Commit()
	if err != nil {
		logger.ErrorLog("An error occured while committing transaction - InitScripts - dbInitializer.go ", err.Error())
	}
}

//#region Models

type Root struct {
	XMLName   xml.Name   `xml:"Root"`
	Languages []Language `xml:"Language"`
}

type Language struct {
	XMLName   xml.Name   `xml:"Language"`
	Code      string     `xml:"Code,attr"`
	Resources []Resource `xml:"Resource"`
}

type Resource struct {
	XMLName xml.Name `xml:"Resource"`
	Key     string   `xml:"Key,attr"`
	Message string   `xml:"Message,attr"`
}

//#endregion
