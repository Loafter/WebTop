// WebTop project main.go
package main

import "fmt"

func main() {
	topJson := new(TopJsonService)
	listenPort := 9977
	if err := topJson.Start(listenPort); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Web service start success")

}
