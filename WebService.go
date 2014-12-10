package main

import (
	//	"fmt"
	//	"log"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

//main service class
type TopJsonService struct {
	requestSelector RequestSelector
	mutex           sync.Mutex
}

//start and init service
func (serviceState *TopJsonService) Start(listenPort int) error {
	serviceState.requestSelector = RequestSelector{}
	serviceState.requestSelector.Init()
	http.HandleFunc("/", serviceState.ServeHTTP)
	retVal := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
	return retVal
}

//serve http responce in different thread
func (serviceState *TopJsonService) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	profile := "Test"
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(js)
}
