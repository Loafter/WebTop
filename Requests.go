package main

import "net/http"

const (
	ServiceStatus = iota
	SystemMonitor
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

type SystemStateRequest struct {
	BasicRequest
}

func (basicRequest BasicRequest) RequestType() int {
	return ServiceStatus
}
