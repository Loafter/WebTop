package main

import "net/http"
import "encoding/json"
import "errors"
import "log"

type ServiceStateRequest struct {
	BasicRequest
}

type ServiceStateResponse struct {
	UpTime       bool
	WorkFromRoot bool
	ProcStat     bool
	MemInfo      bool
}

type ServiceStateDispatcher struct {
}

func (serviceStateDispatcher *ServiceStateDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//this is service is not need lock
	serviceState := ServiceStateResponse{true, true, true, true}
	js, err := json.Marshal(serviceState)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		stErr := "error: Can't service state response"
		log.Println(stErr)
		responseWriter.Write(js)
		return errors.New(stErr)
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(js)
	return nil
}
