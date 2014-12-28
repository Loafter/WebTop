package main

import "net/http"
import "syscall"
import "encoding/json"

type SystemMonitorResponse struct {
	syscall.Sysinfo_t
}

type SystemMonitorDispatcher struct {
}

func (serviceStateDispatcher *SystemMonitorDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//this is service is not need lock
	//req := request.(SystemStateRequest)
	systemInfo := SystemMonitorResponse{}
	if err := syscall.Sysinfo(&systemInfo.Sysinfo_t); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	js, err := json.Marshal(systemInfo)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(js)
	return nil
}
