package main

import "time"
import "net/http"
import "encoding/json"

type ServiceStateResponse struct {
	CurrentTime  time.Time
	ServiceState bool
}

type ServiceStateDispatcher struct {
}

func (serviceStateDispatcher *ServiceStateDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//this is service is not need lock
	serviceState := ServiceStateResponse{time.Now(), true}
	js, err := json.Marshal(serviceState)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(js)
	return nil
}
