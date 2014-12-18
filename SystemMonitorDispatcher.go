package main

import (
	"net/http"
	"syscall"
)
import "encoding/json"

type SystemMonitorResponse struct {
	syscall.Sysinfo_t
}

type SystemMonitorDispatcher struct {
}

func (serviceStateDispatcher *SystemMonitorDispatcher) Dispatch(request Request) error {
	//this is service is not need lock
	req := request.(SystemStateRequest)
	systemInfo := SystemMonitorResponse{}
	if err := syscall.Sysinfo(&systemInfo.Sysinfo_t); err == nil {
		http.Error(req.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	js, err := json.Marshal(systemInfo)
	if err != nil {
		http.Error(req.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	req.ResponseWriter.Header().Set("Content-Type", "application/json")
	req.ResponseWriter.Write(js)
	return nil
}
