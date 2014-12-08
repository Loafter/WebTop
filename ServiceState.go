package main

import "fmt"
import "sync"

type ServiceState struct {
	mutex sync.Mutex
}

func (serviceState ServiceState) Dispatch(request Request) bool {
	serviceState.mutex.Lock()
	fmt.Println("Service status checked ")
	serviceState.mutex.Unlock()
	return false
}
