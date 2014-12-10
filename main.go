// WebTop project main.go
package main

import "fmt"
import "time"

func main() {
	fmt.Println("Start!")
	topJson := TopJsonService{}
	listenPort := 9999
	topJson.Start(listenPort)
	time.Sleep(40000)

}
