// RequestDispatcher
package main

type RequestDispatcher interface {
	Dispatch(request Request) bool
	GetChan() chan Request
}
