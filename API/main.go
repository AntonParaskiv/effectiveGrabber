package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	EnvHttpListen = "HTTP_LISTEN"
	EnvRpmLimit   = "RPM_LIMIT"
)

var (
	CurrentRpm      int64
	RpmDecreaseChan chan time.Time
	RpmLimit        int64 = 60
)

func main() {
	// parse env http listen
	httpListen := os.Getenv(EnvHttpListen)
	if len(httpListen) == 0 {
		httpListen = ":4000"
	}

	// parse env rmp limit
	rpmLimitString := os.Getenv(EnvRpmLimit)
	if len(rpmLimitString) == 0 {
		log.Println("use default RPM limit:", RpmLimit)
	} else {
		rpmLimit, err := strconv.ParseInt(rpmLimitString, 10, 32)
		if err != nil {
			log.Fatalf("parse env %s failed: %s", EnvRpmLimit, err.Error())
		}
		if rpmLimit <= 0 {
			log.Fatalf("env %s must be > 0", EnvRpmLimit)
		}
		RpmLimit = rpmLimit
		log.Println("use RPM limit:", RpmLimit)
	}

	// define handlers
	http.HandleFunc("/api/one", handlerOne)
	http.HandleFunc("/api/two", handlerTwo)

	// run rpm count decreaser
	RpmDecreaseChan = make(chan time.Time, 100)
	go runRpmCountDecreaser()

	log.Println("service API starting at", httpListen)
	log.Fatal(http.ListenAndServe(httpListen, nil))
}

func handlerOne(w http.ResponseWriter, r *http.Request) {
	handleWithSleep(500, w, r)
}

func handlerTwo(w http.ResponseWriter, r *http.Request) {
	handleWithSleep(2000, w, r)
}

func handleWithSleep(sleep int, w http.ResponseWriter, r *http.Request) {
	increaseRpmCount()

	currentRpm := atomic.LoadInt64(&CurrentRpm)
	if currentRpm > RpmLimit {
		log.Println(r.URL, "request limit is exceeded. current rpm:", currentRpm)
		io.WriteString(w, `{ "result": 0 }`)
		return
	}

	log.Println(r.URL, "sleep:", sleep, "current rpm:", currentRpm)
	<-time.After(time.Duration(sleep) * time.Millisecond)
	io.WriteString(w, `{ "result": 1 }`)
	return
}

func increaseRpmCount() {
	// increase counter
	atomic.AddInt64(&CurrentRpm, 1)

	// send decrease time
	decreaseTime := time.Now().Add(time.Minute)
	RpmDecreaseChan <- decreaseTime
}

func runRpmCountDecreaser() {
	for {
		<-time.After(time.Until(<-RpmDecreaseChan))
		atomic.AddInt64(&CurrentRpm, -1)
	}
}
