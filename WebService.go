package main

import (
	"fmt"
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
	fmt.Println(":" + strconv.Itoa(listenPort))

	http.HandleFunc("/", serviceState.ServeHTTP)
	retVal := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
	return retVal
}

//serve http responce in different thread
func (serviceState *TopJsonService) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	serviceState.requestSelector.Dispatch(ServiceStateRequest{BasicRequest{ServiceStatus, responseWriter, request}})
}
