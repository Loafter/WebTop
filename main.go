// WebTop project main.go
package main

import "log"

func main() {
	topJson := new(TopJsonService)
	listenPort := 9977
	if err := topJson.Start(listenPort); err != nil {
		log.Println(err)
	}

}
