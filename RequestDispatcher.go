// RequestDispatcher
package main

type RequestDispatcher interface {
	Dispatch(request Request) error
}
