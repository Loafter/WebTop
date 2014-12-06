package main

const (
	ServiceStatus = iota
)

type BasicRequest struct {
	Type int
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
