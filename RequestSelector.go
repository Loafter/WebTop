package main

type RequestSelector struct {
	selectorRequestMap map[int]RequestDispatcher
}

func (requestSelector *RequestSelector) Dispatch(request Request) bool {
	return true
}

func (requestSelector *RequestSelector) Init(request Request) {
	requestSelector.selectorRequestMap = make(map[int]RequestDispatcher)
	requestSelector.selectorRequestMap[ServiceStatus] = new(ServiceState)
}
