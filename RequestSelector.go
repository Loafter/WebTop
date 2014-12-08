package main

//import "fmt"
import "errors"

type RequestSelector struct {
	selectorRequestMap map[int]RequestDispatcher
}

func (requestSelector *RequestSelector) Dispatch(request Request) error {
	//I don't protect multiple read in different thread
	if selector, contains := requestSelector.selectorRequestMap[request.RequestType()]; !contains {
		return errors.New("Usupported message type")
	} else {
		return selector.Dispatch(request)
	}

}

func (requestSelector *RequestSelector) Init() error {
	requestSelector.selectorRequestMap = make(map[int]RequestDispatcher)
	requestSelector.selectorRequestMap[ServiceStatus] = ServiceState{}
	if requestSelector.selectorRequestMap == nil {
		return errors.New("Can't create map")
	}
	return nil
}
