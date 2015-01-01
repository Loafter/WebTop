// WebTop project main.go
package main

import "fmt"

func main() {
	top := new(Top)

	fmt.Println(top.GetProcessList())

	/*topJson := new(TopJsonService)
	listenPort := 9977
	if err := topJson.Start(listenPort); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Web service start success")
	time.Sleep(400000)*/

}
