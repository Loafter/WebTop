package main

import "net/http"
import "encoding/json"
import "sync"

type SystemMonitorResponse struct {
	CPUUsage  CPUAverage
	MemSample MemSample
}

type SystemMonitorDispatcher struct {
	accessMutes   sync.Mutex
	lastCPUSample CPUSample
}

func (serviceStateDispatcher *SystemMonitorDispatcher) getCPUUsage() CPUAverage {
	cpuAveragePer := GetCPUAverage(serviceStateDispatcher.lastCPUSample, GetCPUSample())
	return cpuAveragePer
}

func (serviceStateDispatcher *SystemMonitorDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//this is service need lock
	serviceStateDispatcher.accessMutes.Lock()
	defer serviceStateDispatcher.accessMutes.Unlock()
	systemInfo := SystemMonitorResponse{CPUUsage: serviceStateDispatcher.getCPUUsage(), MemSample: GetMemSample()}
	js, err := json.Marshal(systemInfo)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(js)
	return nil
}
