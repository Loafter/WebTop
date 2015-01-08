package main

import "net/http"
import "encoding/json"

import "sync"
import "time"
import "errors"

type SystemStateRequest struct {
	BasicRequest
}

type SystemMonitorResponse struct {
	CPUUsage  CPUAverage
	MemSample MemSample
}

type SystemMonitorDispatcher struct {
	accessMutes     sync.Mutex
	lastCPUSample   CPUSample
	lastCPUAverage  CPUAverage
	mesureJob       BatchJob
	lastRequestTime time.Time
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
		secondLastRequest := time.Now().Sub(serviceStateDispatcher.lastRequestTime)
		if secondLastRequest.Seconds() > 5 {
			serviceStateDispatcher.mesureJob.Stop()
		}
	}
	defer serviceStateDispatcher.accessMutes.Unlock()
}

func (serviceStateDispatcher *SystemMonitorDispatcher) StartMesure() error {
	serviceStateDispatcher.mesureJob.Job = serviceStateDispatcher.mesureCPU
	err := serviceStateDispatcher.mesureJob.Start()
	if err != nil {
		return errors.New("error: Can't start mesure job\n" + err.Error())
	}
	return nil
}
func (serviceStateDispatcher *SystemMonitorDispatcher) StopMesure() error {
	err := serviceStateDispatcher.mesureJob.Stop()
	if err != nil {
		return errors.New("error: Can't stop mesure job\n" + err.Error())
	}
	return nil
}

func (serviceStateDispatcher *SystemMonitorDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//this is service need lock
	if !serviceStateDispatcher.mesureJob.runJob {
		serviceStateDispatcher.mesureJob.Start()
	}

	systemInfo := SystemMonitorResponse{CPUUsage: serviceStateDispatcher.getCPUUsage(), MemSample: GetMemSample()}
	js, err := json.Marshal(systemInfo)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return errors.New("error: Can't system state response \n" + err.Error())
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(js)
	serviceStateDispatcher.accessMutes.Lock()
	defer serviceStateDispatcher.accessMutes.Unlock()
	serviceStateDispatcher.lastRequestTime = time.Now()
	return nil
}
