package main

import "io/ioutil"
import "testing"
import "strconv"
import "regexp"

//process info sructure
type ProcessItem struct {
	Pid     int
	User    string
	Virt    int
	Cpu     float32
	Memory  float32
	Command string
}

type Top struct {
}

//
func (top *Top) getAllPids() (*[]ProcessItem, error) {
	dirContent, err := ioutil.ReadDir("/proc/")
	if err != nil {
		return nil, err
	}
	var process []ProcessItem
	for _, element := range dirContent {
		validPID := regexp.MustCompile(`^[0-9]+$`)
		if element.IsDir() && validPID.MatchString(element.Name()) {
			pid, _ := strconv.Atoi(element.Name())
			process = append(process, ProcessItem{Pid: pid})
		}
	}
	return &process, nil
}

func (top *Top) fillProcessInfo() (*[]ProcessItem, error) {

	return nil, nil
}

//
func (top *Top) GetProcessList() ([]ProcessItem, error) {
	processItemsRef, err := top.getAllPids()
	if err != nil {
		return nil, err
	}
	processItems := *processItemsRef
	return processItems, nil
}

func (top *Top) KillProcess(pid int) error {
	//this is service is need lock
	return nil
}

func TestGetProcessList(t *testing.T) {
	top := new(Top)
	top.GetProcessList()
}
