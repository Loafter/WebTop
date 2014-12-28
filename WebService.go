package main

import "errors"
import "fmt"
import "io/ioutil"
import "net/http"
import "strconv"
import "encoding/json"

type test_struct struct {
	Test string
}

//main service class
type TopJsonService struct {
	requestSelector RequestSelector
}

//start and init service
func (serviceState *TopJsonService) Start(listenPort int) error {
	serviceState.requestSelector = RequestSelector{}
	serviceState.requestSelector.Init()
	fmt.Println(":" + strconv.Itoa(listenPort))
	http.HandleFunc("/webtop", serviceState.ServeHTTP)
	http.HandleFunc("/index.html", serviceState.ServePage)
	http.HandleFunc("/samplejson", serviceState.ReturnDummyReq)
	retVal := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
	return retVal
}

//parse json message
func parseJsonRequest(request *http.Request) (Request, error) {
	bodyData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errors.New("Can't read request body \n " + err.Error())
	}

	var basicRequest BasicRequest
	err = json.Unmarshal(bodyData, &basicRequest)
	if err != nil {
		return nil, errors.New("error: Can't parse basic data \n " + err.Error())
	}
	switch basicRequest.Type {
	case ServiceStatus:
		var serviceStateRequest ServiceStateRequest
		err := json.Unmarshal(bodyData, &serviceStateRequest)
		if err != nil {
			return nil, errors.New("error: Can't parse service state request \n " + err.Error())
		}
		return serviceStateRequest, nil
	case SystemMonitor:
		var systemStateRequest SystemStateRequest
		err := json.Unmarshal(bodyData, &systemStateRequest)
		if err != nil {
			return nil, errors.New("error: Can't parse system state request \n " + err.Error())
		}
		return systemStateRequest, nil
	}
	return nil, errors.New("error: Unknown request type")
}

//serve http responce in different thread
func (service *TopJsonService) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if reqData, err := parseJsonRequest(request); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	} else {
		service.requestSelector.Dispatch(reqData, responseWriter, request)
	}

}

//serve main page request
func (service *TopJsonService) ServePage(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type: text/html", "*")
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		responseWriter.Write([]byte("error: Can't find start page \n " + err.Error()))
		return
	}
	responseWriter.Write(content)
}

func (service *TopJsonService) ReturnDummyReq(responseWriter http.ResponseWriter, request *http.Request) {
	//this is service is not need lock
	serviceState := ServiceStateRequest{BasicRequest{1}}
	js, err := json.Marshal(serviceState)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(js)
	return

}
