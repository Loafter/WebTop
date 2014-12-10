package main

import "net/http"

const (
	ServiceStatus = iota
)

//basic http request with writer
type Request interface {
	RequestType() int
}

type BasicRequest struct {
	Type           int
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type ServiceStateRequest struct {
	BasicRequest
}

func (basicRequest BasicRequest) RequestType() int {
	return ServiceStatus
}
