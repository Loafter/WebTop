package main

import "net/http"
import "sync"

type TopResponse struct {
}

type TopDispatcher struct {
	accessMutes sync.Mutex
	top         Top
}

func (serviceStateDispatcher *TopDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	//this is service need lock
	serviceStateDispatcher.accessMutes.Lock()
	defer serviceStateDispatcher.accessMutes.Unlock()
	switch request.RequestType() {
	case TopProcess:
		return nil
	case KillProcess:
		return nil

	}
	return nil
}
