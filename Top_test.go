package main

import "testing"
import "time"
import "fmt"

func TestTopList(t *testing.T) {
	top := new(Top)
	top.StartCollectInfo()
	for i := 0; i < 10; i++ {
		processList, err := top.GetProcessList()
		time.Sleep(500 * time.Millisecond)
		if processList == nil {
			t.Errorf("GetProcessList return null ")
			return
		} else if err != nil {
			t.Errorf("")
			return
		}
	}

}

func TestTopCpu(t *testing.T) {
	top := new(Top)
	top.StartCollectInfo()
	for i := 0; i < 5; i++ {
		processList, err := top.GetProcessList()
		time.Sleep(500 * time.Millisecond)
		if processList == nil {
			t.Errorf("GetProcessList return null ")
			return
		} else if err != nil {
			t.Errorf("")
			return
		} else {
			for _, element := range processList {
				if element.Cpu != 0 {
					fmt.Printf("Name: %v Usage: %v \n", element.Name, int32(element.Cpu*100))
				}
			}
			fmt.Println("_______________________________________________________________________________________________________________________________________________________________________")
		}
	}

}
