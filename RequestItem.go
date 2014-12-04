package main

const (
	ServiceStateRequest = iota
	ProcessorUsageRequest
	MemoryUsage
)

type RequestItem interface {
	GetType() int
}
