package main

import "fmt"

type ServiceState struct {
}

func (requestSelector ServiceState) Dispatch(request Request) bool {
	fmt.Println("Service status checked")
	return false
}
