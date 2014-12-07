package main

import "fmt"

type RequestSelector struct {
	selectorRequestMap map[int]RequestDispatcher
}

func (requestSelector *RequestSelector) Dispatch(request Request) bool {

	if selector, contains := requestSelector.selectorRequestMap[request.RequestType()]; !contains {
		fmt.Println("Usupported message type")
		return false
	} else {
		return selector.Dispatch(request)
	}

}

func (requestSelector *RequestSelector) Init() {
	requestSelector.selectorRequestMap = make(map[int]RequestDispatcher)
	requestSelector.selectorRequestMap[ServiceStatus] = ServiceState{}
}
