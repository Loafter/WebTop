package main

import (
	"fmt"
	"io/ioutil"
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
	http.HandleFunc("/index.html", serviceState.ServePage)
	retVal := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
	return retVal
}

//parse json message
//func (serviceState TopJsonService) parseJson(request *http.Request) error {

//}

//serve http responce in different thread
func (serviceState *TopJsonService) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	serviceState.requestSelector.Dispatch(ServiceStateRequest{BasicRequest{ServiceStatus}}, responseWriter, request)
}

//serve main page request
func (serviceState *TopJsonService) ServePage(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type: text/html", "*")
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		responseWriter.Write([]byte("Can't find start page"))
		return
	}
	responseWriter.Write(content)
}
