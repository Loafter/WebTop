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
	Type int `json:"Type,string,omitempty"`
}

func (basicRequest BasicRequest) RequestType() int {
	return basicRequest.Type
}

type ServiceStateRequest struct {
	BasicRequest
}

type SystemStateRequest struct {
	BasicRequest
}
