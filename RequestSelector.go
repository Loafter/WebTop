package main

type RequestSelector struct {
	selectorRequestMap map[int]RequestDispatcher
}

func (requestSelector *RequestSelector) Dispatch(request Request) bool {
	return true
}

func (requestSelector *RequestSelector) Init(request Request) {
	requestSelector.selectorRequestMap = make(map[int]RequestDispatcher)
}

/*func (requestSelector *RequestSelector) Init(){
	requestSelector.selectorReqestMap = map[int]RequestDispatcher
	requestSelector.selectorReqestMap[0] = RequestDispatcher{}
	p.selectorReqestMap = elements
	return p
}
*/
