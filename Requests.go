package main

import "net/http"

const (
	ServiceStatus = iota
)

type BasicRequest struct {
	Type           int
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type Request interface {
	RequestType() int
}

type ServiceStateStatus struct {
	BasicRequest
}

func (requestSelector ServiceStateStatus) RequestType() int {
	return ServiceStatus
}
