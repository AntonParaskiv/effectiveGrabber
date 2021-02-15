package main

import (
	"fmt"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/infrastructure/HttpClient"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/interafaces/ApiRepository"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/usecases/GrabberInteractor"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/usecases/TaskGenerator"
	"log"
	"net/url"
	"os"
	"strconv"
)

const (
	EnvApiBaseUrl = "API_BASE_URL"
	EnvRpmLimit   = "RPM_LIMIT"
)

func main() {
	// parse env base api url
	apiBaseUrlString := os.Getenv(EnvApiBaseUrl)
	if len(apiBaseUrlString) == 0 {
		log.Fatalln("need env", EnvApiBaseUrl)
	}
	apiBaseUrl, err := url.Parse(apiBaseUrlString)
	if err != nil {
		err = fmt.Errorf("parse base url failed: %w", err)
		log.Fatal(err)
		return
	}

	// parse env rmp limit
	var rpmLimit int64 = 60
	rpmLimitString := os.Getenv(EnvRpmLimit)
	if len(rpmLimitString) == 0 {
		log.Println("use default RPM limit:", rpmLimit)
	} else {
		rpmLimit, err = strconv.ParseInt(rpmLimitString, 10, 64)
		if err != nil {
			log.Fatalf("parse env %s failed: %s", EnvRpmLimit, err.Error())
		}
		if rpmLimit <= 0 {
			log.Fatalf("env %s must be > 0", EnvRpmLimit)
		}
		log.Println("use RPM limit:", rpmLimit)
	}

	// create api layers
	httpClient := HttpClient.New().SetBaseUrl(*apiBaseUrl)
	apiRepository := ApiRepository.New().SetHttpClient(httpClient)

	// create task generator
	taskOneChan := make(chan struct{}, 1)
	taskTwoChan := make(chan struct{}, 1)
	taskGenerator := TaskGenerator.New().
		SetTaskOneChan(taskOneChan).
		SetTaskTwoChan(taskTwoChan)
	taskGenerator.Start()

	// create grabber interactor
	grabberInteractor := GrabberInteractor.New().
		SetApiRepository(apiRepository).
		SetTaskOneChan(taskOneChan).
		SetTaskTwoChan(taskTwoChan).
		SetRpmLimit(rpmLimit)
	grabberInteractor.StartGrabbing()

	log.Println("finished")
}
