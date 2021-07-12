package main

import (
	"fmt"
	"github.com/ltinyho/ltcache/ltcache"
	"log"
	"net/http"
	"time"
)

var db = map[string]string{
	"Lt":  "630",
	"Sql": "589",
	"Zzl": "567",
}

func main() {
	ltcache.NewGroup("scores", 2<<10, ltcache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		time.Sleep(time.Second*3)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))


	addr:="localhost:9999"
	peers:=ltcache.NewHTTPPool(addr)
	log.Println("ltcache is running at",addr)
	log.Fatal(http.ListenAndServe(addr,peers))
}
