package main

import "errors"
import "net/http"
import "log"

type RequestSelector struct {
	selectorRequestMap map[int]RequestDispatcher
}

func (requestSelector *RequestSelector) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//I don't protect multiple read in different thread
	if selector, contains := requestSelector.selectorRequestMap[request.RequestType()]; !contains {
		stErr := "error: Usupported message type"
		log.Println(stErr)
		return errors.New(stErr)
	} else {
		return selector.Dispatch(request, responseWriter, httpRequest)
	}
}

func (requestSelector *RequestSelector) Init() error {
	//create map
	requestSelector.selectorRequestMap = make(map[int]RequestDispatcher)
	if requestSelector.selectorRequestMap == nil {
		stErr := "error: Can't create map"
		log.Println(stErr)
		return errors.New(stErr)
	}
	requestSelector.selectorRequestMap[ServiceStatus] = new(ServiceStateDispatcher)
	systemMonitorDispatcher := new(SystemMonitorDispatcher)
	systemMonitorDispatcher.StartMesure()
	requestSelector.selectorRequestMap[SystemMonitor] = systemMonitorDispatcher
	//same dispatcher for two message
	topProcessDispatcher := new(TopDispatcher)
	topProcessDispatcher.StartMesure()
	requestSelector.selectorRequestMap[TopProcess] = topProcessDispatcher
	requestSelector.selectorRequestMap[KillProcess] = topProcessDispatcher
	return nil
}
