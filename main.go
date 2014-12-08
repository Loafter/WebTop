// WebTop project main.go
package main

import "fmt"
import "time"

func main() {
	fmt.Println("Start!")

	requestSelector := RequestSelector{}
	requestSelector.Init()

	//if err := requestSelector.Dispatch(ServiceStateStatus{}); err {
	//	fmt.Println("Dispatch error")
	//}
	time.Sleep(4000)

}
