package main

type RequestSelector struct {
	selectorRequestMap map[int]RequestDispatcher
}

func (requestSelector *RequestSelector) Dispatch(request Request) bool {
	return false
}

func (requestSelector *RequestSelector) Init() {
	requestSelector.selectorRequestMap = make(map[int]RequestDispatcher)
	requestSelector.selectorRequestMap[ServiceStatus] = ServiceState{}
}
