package main

type ServiceState struct {
}

func (requestSelector ServiceState) Dispatch(request Request) bool {
	return true
}
