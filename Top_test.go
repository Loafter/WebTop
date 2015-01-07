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
func TestTopTicks(t *testing.T) {
	top := new(Top)
	top.StartCollectInfo()
	for i := 0; i < 5; i++ {
		Ticks, err := top.getTicksProcessor()
		time.Sleep(500 * time.Millisecond)
		if err != nil {
			t.Errorf("")
			return
		} else {
			fmt.Printf("Ticks: %v \n", Ticks)
			fmt.Println("_______________________________________________________________________________________________________________________________________________________________________")
		}
	}

}

func ThreadTest() {
	fmt.Println("Do some job")
	time.Sleep(1000 * time.Millisecond)
}
func TestBatchJob(t *testing.T) {
	batchJob := BatchJob{}
	err := batchJob.Start()
	if err == nil {
		t.Errorf("error: Test start not inited job failed ")
	}

	err = batchJob.Stop()
	if err == nil {
		t.Errorf("error: Test stop not inited job failed ")
	}
	batchJob.Job = ThreadTest
	batchJob.Start()
	time.Sleep(500 * time.Millisecond)
	err = batchJob.Stop()
	if err != nil {
		t.Errorf("error: Test stop failed ")
	}
}
