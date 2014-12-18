// RequestDispatcher
package main

import "net/http"

type RequestDispatcher interface {
	Dispatch(request Request, responseWriter http.ResponseWriter, httpRequest *http.Request) error
}
