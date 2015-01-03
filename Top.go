package main

import "io/ioutil"
import "testing"
import "strconv"
import "regexp"
import "errors"

//process info sructure
type ProcessItem struct {
	Pid    int
	Name   string
	User   string
	Cpu    float32
	Memory float32
}

type Top struct {
}

//
func (top *Top) getAllPids() ([]ProcessItem, error) {
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
	return process, nil
}

func (top *Top) fillProcessInfo(processItems []ProcessItem) {
	for i, _ := range processItems {
		statFileData, err := ioutil.ReadFile("/proc/" + strconv.Itoa(processItems[i].Pid) + "/status")
		if err == nil {
			statFileStr := string(statFileData)
			regName := regexp.MustCompile("Name:\t(.*)\n")
			processItems[i].Name = regName.FindAllStringSubmatch(statFileStr, -1)[0][1]
			regUid := regexp.MustCompile("Uid:\t(\\w+)")
			user := regUid.FindAllStringSubmatch(statFileStr, -1)[0][1]
			processItems[i].User = user
		}
	}
}

//
func (top *Top) GetProcessList() ([]ProcessItem, error) {
	processItemsSlice, err := top.getAllPids()
	if err != nil {
		return nil, errors.New("Can't read proc directory")
	}
	top.fillProcessInfo(processItemsSlice)
	return processItemsSlice, nil
}

func (top *Top) KillProcess(pid int) error {
	//this is service is need lock
	return nil
}

func TestGetProcessList(t *testing.T) {
	top := new(Top)
	top.GetProcessList()
}
