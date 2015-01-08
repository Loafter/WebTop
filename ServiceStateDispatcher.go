package main

import "net/http"
import "encoding/json"
import "errors"

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
		return errors.New("error: Can't service state response \n" + err.Error())
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(js)
	return nil
}
