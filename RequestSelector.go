package main

//import "fmt"
import "errors"
import "net/http"

type RequestSelector struct {
	selectorRequestMap map[int]RequestDispatcher
}

func (requestSelector *RequestSelector) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//I don't protect multiple read in different thread
	if selector, contains := requestSelector.selectorRequestMap[request.RequestType()]; !contains {
		return errors.New("Usupported message type")
	} else {
		return selector.Dispatch(request, responseWriter, httpRequest)
	}
}

func (requestSelector *RequestSelector) Init() error {
	//create map
	requestSelector.selectorRequestMap = make(map[int]RequestDispatcher)
	if requestSelector.selectorRequestMap == nil {
		return errors.New("error: Can't create map")
	}
	requestSelector.selectorRequestMap[ServiceStatus] = &ServiceStateDispatcher{}
	requestSelector.selectorRequestMap[SystemMonitor] = &SystemMonitorDispatcher{}

	return nil
}
