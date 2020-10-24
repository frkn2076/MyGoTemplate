package resource

import(
	"encoding/xml"

	"app/MyGoTemplate/cache"
	"app/MyGoTemplate/helper"
)

var isFirstCall bool = true

func GetValue(key string) string {
	value := cache.Get(key)
	if(value != ""){
		isFirstCall = true
		return value
	} else if(!isFirstCall){
		return ""
	} else {
		isFirstCall = false
		settingsCacheLoad()
		return GetValue(key)
	}
}

func init() {
	settingsCacheLoad()
}

//#region Helper

func settingsCacheLoad(){
	var root Root

	helper.LoadModel("resource/settings.xml", &root)

	for _, item:= range root.Items{
		cache.Set(item.Key, item.Value, -1)
	}
}

//#endregion


//#region Models

type Root struct {
    XMLName xml.Name `xml:"Root"`
    Items   []Item   `xml:"Item"`
}

type Item struct {
	XMLName xml.Name `xml:"Item"`
    Key   string `xml:"Key,attr"`
    Value string `xml:"Value,attr"`
}

//#endregion