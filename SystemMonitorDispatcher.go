package main

import (
	"net/http"
	"syscall"
)
import "encoding/json"

type SystemMonitorResponse struct {
}

type SystemMonitorDispatcher struct {
	syscall.Sysinfo_t
}

func (serviceStateDispatcher *SystemMonitorDispatcher) Dispatch(request Request) error {
	//this is service is not need lock
	req := request.(ServiceStateRequest)
	systemInfo := SystemMonitorResponse{}
	js, err := json.Marshal(systemInfo)
	if err != nil {
		http.Error(req.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	req.ResponseWriter.Header().Set("Content-Type", "application/json")
	req.ResponseWriter.Write(js)
	return nil
}
