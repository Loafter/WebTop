package main

import "testing"
import "fmt"
import "time"

func TestTop(t *testing.T) {
	top := new(Top)
	top.StartCollectInfo()
	time.Sleep(600 * time.Millisecond)
	//top.getAllPids()
	processList, err := top.GetProcessList()

	if processList == nil {
		t.Errorf("GetProcessList return null ")
		return
	} else if err != nil {
		t.Errorf("")
		return
	}
	fmt.Println("_____________________________________________________________Process list_____________________________________________________________")
	fmt.Println(processList)
	fmt.Println("_____________________________________________________________Process list_____________________________________________________________")
}
