package main

import "net/http"
import "errors"
import "encoding/json"

type TopResponse struct {
	ProcessItems []ProcessItem
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
		return nil
	}
	return nil
}
func (serviceStateDispatcher *TopDispatcher) StartMesure() {
	serviceStateDispatcher.top.StartCollectInfo()
}
