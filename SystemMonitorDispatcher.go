package main

import "net/http"
import "encoding/json"
import "sync"
import "time"

type SystemStateRequest struct {
	BasicRequest
}

type SystemMonitorResponse struct {
	CPUUsage  CPUAverage
	MemSample MemSample
}

type SystemMonitorDispatcher struct {
	accessMutes    sync.Mutex
	lastCPUSample  CPUSample
	lastCPUAverage CPUAverage
}

func (serviceStateDispatcher *SystemMonitorDispatcher) getCPUUsage() CPUAverage {
	serviceStateDispatcher.accessMutes.Lock()
	defer serviceStateDispatcher.accessMutes.Unlock()
	lastCPUAverage := serviceStateDispatcher.lastCPUAverage
	return lastCPUAverage

}

func (serviceStateDispatcher *SystemMonitorDispatcher) mesureCPU() {
	for {
		serviceStateDispatcher.accessMutes.Lock()
		serviceStateDispatcher.lastCPUAverage = GetCPUAverage(serviceStateDispatcher.lastCPUSample, GetCPUSample())
		serviceStateDispatcher.lastCPUSample = GetCPUSample()
		serviceStateDispatcher.accessMutes.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
	defer serviceStateDispatcher.accessMutes.Unlock()
}

func (serviceStateDispatcher *SystemMonitorDispatcher) StopMesure() error {
	return nil
}
func (serviceStateDispatcher *SystemMonitorDispatcher) StartMesure() error {
	go serviceStateDispatcher.mesureCPU()
	return nil
}

func (serviceStateDispatcher *SystemMonitorDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//this is service need lock
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
