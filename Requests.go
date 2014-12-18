package main

const (
	ServiceStatus = iota
	SystemMonitor
)

//basic http request with writer
type Request interface {
	RequestType() int
}

type BasicRequest struct {
	Type int
}

func (basicRequest BasicRequest) RequestType() int {
	return ServiceStatus
}

type ServiceStateRequest struct {
	BasicRequest
}

type SystemStateRequest struct {
	BasicRequest
}
