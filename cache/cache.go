package cache

import (
	"runtime/debug"
	"time"

	"app/MyGoTemplate/logger"

	"github.com/coocood/freecache"
)

var cache *freecache.Cache = loadCache()


func loadCache() *freecache.Cache {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	logger.InfoLog("Cache created with size ", cacheSize)

	debug.SetGCPercent(20)
	
	return cache
} 

//expireSeconds <= 0 means no expire
func Set(key string, value string, expireDuration int){
	cache.Set([]byte(key), []byte(value), expireDuration)
}

func Get(key string) string {
	//if cache has not key, value is empty string
	value, err := cache.Get([]byte(key))
	if err != nil {
		logger.ErrorLog("Cache couldn't find the key ", key)
	} 
	return string(value)
}

func Delete(key string) {
	cache.Del([]byte(key))
}

func Reset(){
	cache.Clear()
}

func GetAvaregeAccessTime() int64 {
	return cache.AverageAccessTime() / int64(time.Second)
}


