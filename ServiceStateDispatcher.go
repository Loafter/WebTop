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

func (serviceStateDispatcher ServiceStateDispatcher) Dispatch(request Request) error {
	//this is service is not need lock
	req := request.(ServiceStateRequest)
	serviceState := ServiceStateResponse{time.Now(), true}
	js, err := json.Marshal(serviceState)
	if err != nil {
		http.Error(req.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	req.ResponseWriter.Header().Set("Content-Type", "application/json")
	req.ResponseWriter.Write(js)
	return nil
}
