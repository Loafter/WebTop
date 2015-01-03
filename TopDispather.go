package main

import "net/http"

type TopResponse struct {
}

type TopDispatcher struct {
	top Top
}

func (serviceStateDispatcher *TopDispatcher) Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error {
	switch request.RequestType() {
	case TopProcess:
		return nil
	case KillProcess:
		return nil

	}
	return nil
}
