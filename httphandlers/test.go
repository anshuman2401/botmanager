package httphandlers

import (
	"botmanager/cache"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

type TestObject struct{
	Key string
	Val string
}

func HelloTestSet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	RedisConnection := cache.NewClient()

	var pa TestObject
	err := json.NewDecoder(r.Body).Decode(&pa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Need to set some ttl
	currentVal := RedisConnection.Incr(pa.Key)
	fmt.Println("Current Val: "  + strconv.FormatInt(currentVal, 10))

	currentTime := time.Now()
	//currentTime := time.Now()
	//minute := currentTime.Format("2006-01-02 15:04") // Minute
	//hour := currentTime.Format("2006-01-02 15") // Hour
	//day := currentTime.Format("2006-01-02") // Day
	//month := currentTime.Format("2006-01") // Month

	fmt.Println(currentTime.String())
	fmt.Println(pa)
	w.Write([]byte("{" + pa.Key + ":" + pa.Val +"}"))
}

func HelloTestGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("I am in func")
	RedisConnection := cache.NewClient()
	msg := RedisConnection.Get(p.ByName("key"))
	w.Write([]byte("{\"msg\":" + msg))
}

