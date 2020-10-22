package cache

import(
	"fmt"
	"runtime/debug"
	"time"

	"github.com/coocood/freecache"
)

var cache *freecache.Cache = loadCache()


func loadCache() *freecache.Cache {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	debug.SetGCPercent(20)
	
	return cache
} 

func Set(key string, value string, expireDuration int){
	cache.Set([]byte(key), []byte(value), expireDuration)
}

func Get(key string) string {
	got, err := cache.Get([]byte(key))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(got))
	}
	return string(got)
}

func Delete(key string) {
	affected := cache.Del([]byte(key))
	fmt.Println("deleted key ", affected)
}

func GetAvaregeAccessTime() int64 {
	return cache.AverageAccessTime() / int64(time.Millisecond)
}


