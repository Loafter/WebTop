package main

import "net/http"
import "errors"
import "encoding/json"

type TopResponse struct {
	ProcessItems []ProcessItem
}

type KillResponse struct {
	KilResult string
}

type KillRequest struct {
	BasicRequest
	Pid int `Pid:"Type,string,omitempty"`
}
type TopRequest struct {
	BasicRequest
}

type TopDispatcher struct {
	top Top
}

func (serviceStateDispatcher *TopDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	switch request.RequestType() {
	case TopProcess:
		topProcessList, err := serviceStateDispatcher.top.GetProcessList()
		if err != nil {
			return errors.New("error get process list ")
		}
		topResponse := TopResponse{ProcessItems: topProcessList}
		tr, err := json.Marshal(topResponse)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return err
		}
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write(tr)
	case KillProcess:
		killRequest := request.(KillRequest)
		err := serviceStateDispatcher.top.KillProcess(killRequest.Pid)
		var killResultMsg string
		if err != nil {
			killResultMsg = err.Error()
		} else {
			killResultMsg = "Send kill top process" + string(killRequest.Pid)
		}
		killResponseResponse := KillResponse{KilResult: killResultMsg}
		tr, err := json.Marshal(killResponseResponse)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return err
		}
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write(tr)
	}
	return nil
}
func (serviceStateDispatcher *TopDispatcher) StartMesure() {
	serviceStateDispatcher.top.StartCollectInfo()
}
