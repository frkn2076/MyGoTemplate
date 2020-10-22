package resource

import(
	"os"
	"io/ioutil"
	"encoding/xml"

	"app/MyGoTemplate/logger"
	_ "app/MyGoTemplate/cache"
)

var Items []Item = loadItems()

type Root struct {
    XMLName xml.Name `xml:"Root"`
    Items   []Item   `xml:"Item"`
}

type Item struct {
	XMLName xml.Name `xml:"Item"`
    Key   string `xml:"Key,attr"`
    Value string `xml:"Value,attr"`
}

func GetValue(key string) string {
	for _, item := range Items {
		if(item.Key == key){
			return item.Value
		}
	}
	return ""
}

func loadItems() []Item {
	xmlFile, err := os.Open("resource/settings.xml")
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

	var root Root

	xml.Unmarshal(byteValue, &root)

	return root.Items
}