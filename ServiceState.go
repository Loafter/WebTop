package main

import "fmt"
import "sync"

type ServiceState struct {
	mutex sync.Mutex
}

func (serviceState ServiceState) Dispatch(request Request) error {
	serviceState.mutex.Lock()
	fmt.Println("Service status checked ")
	serviceState.mutex.Unlock()
	return nil
}
