package helper

import(
	"os"
	"io/ioutil"
	"encoding/xml"

	"app/MyGoTemplate/logger"
)

func LoadModel(path string, model interface{}) {
	xmlFile, err := os.Open(path)
	if err != nil {
		logger.ErrorLog(err)
		os.Exit(0)
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		logger.ErrorLog(err)
		os.Exit(0)
	}

	xml.Unmarshal(byteValue, &model)
}