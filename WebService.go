package main

import (
	//	"fmt"
	//	"log"
	//"encoding/json"
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
func (serviceState *TopJsonService) Start(port int) error {
	serviceState.requestSelector = RequestSelector{}
	serviceState.requestSelector.Init()
	http.HandleFunc("/", serviceState.ServeHTTP)
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

//serve http responce in different thread
func (serviceState *TopJsonService) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Server", "A Go Web Server")
	responseWriter.WriteHeader(200)

}
