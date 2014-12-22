// WebTop project main.go
package main

import "fmt"
import "time"

func main() {
	topJson := TopJsonService{}
	listenPort := 9977
	if err := topJson.Start(listenPort); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Web service start success")
	time.Sleep(400000)

}
