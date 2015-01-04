package main

import "testing"
import "fmt"

func TestTop(t *testing.T) {
	top := new(Top)
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
