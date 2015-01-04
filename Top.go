package main

import "io/ioutil"
import "testing"
import "strconv"
import "regexp"
import "errors"
import "time"

//import "fmt"
import "sync"

//process info sructure
type ProcessItem struct {
	Pid    int
	Name   string
	User   string
	Cpu    float32
	Memory int
}

type Top struct {
	processItems []ProcessItem
	accessMutes  sync.Mutex
}

func (top *Top) StartCollectInfo() error {
	top.processItems = []ProcessItem{}
	go top.collectInfo()
	return nil
}
func (top *Top) StopCollectInfo() error {
	return nil
}

func (top *Top) getTicksbyPid(pid int) (int64, error) {
	return 0, nil
}

func (top *Top) getTicksMap(pids []int) map[int]int64 {
	ticksMap := make(map[int]int64)
	for _, element := range pids {
		ticks, err := top.getTicksbyPid(element)
		if err == nil {
			ticksMap[element] = 0
		}
		ticksMap[element] = ticks
	}
	return ticksMap
}
func (top *Top) collectInfo() error {
	for {
		pids, err := top.getAllPids()
		if err != nil {
			return errors.New("error: problem read proc filesystem")
		}
		StartTicks := top.getTicksMap(pids)
		time.Sleep(500 * time.Millisecond)
		pids, _ = top.getAllPids()
		EndTicks := top.getTicksMap(pids)
		top.accessMutes.Lock()
		top.processItems = top.fillProcessInfo(StartTicks, EndTicks)
		top.accessMutes.Unlock()
	}
	defer top.accessMutes.Unlock()
	return nil
}

//
func (top *Top) getAllPids() ([]int, error) {
	dirContent, err := ioutil.ReadDir("/proc/")
	if err != nil {
		return nil, err
	}
	var process []int
	for _, element := range dirContent {
		validPID := regexp.MustCompile(`^[0-9]+$`)
		if element.IsDir() && validPID.MatchString(element.Name()) {
			pid, _ := strconv.Atoi(element.Name())
			process = append(process, pid)
		}
	}
	return process, nil
}

func (top *Top) fillProcessInfo(oldTicks map[int]int64, newTicks map[int]int64) []ProcessItem {
	processItems := []ProcessItem{}
	for i, _ := range newTicks {
		statFileData, err := ioutil.ReadFile("/proc/" + strconv.Itoa(i) + "/status")
		if err == nil {
			processItem := ProcessItem{}
			processItem.Pid = i
			statFileStr := string(statFileData)
			regName := regexp.MustCompile("Name:\t(.*)\n")
			processItem.Name = regName.FindAllStringSubmatch(statFileStr, -1)[0][1]
			regUid := regexp.MustCompile("Uid:\t(\\w+)")
			processItem.User = regUid.FindAllStringSubmatch(statFileStr, -1)[0][1]
			processItems = append(processItems, processItem)
			//regMem := regexp.MustCompile("VmRSS:\\s(\\d+)")
			//if regMem.MatchString(statFileStr) == true {
			//	fmt.Printf("Match ")
			//} else {
			//	fmt.Printf("No match ")
			//}
			//processItems[i].User = regMem.FindAllStringSubmatch(statFileStr, -1)[0][0]
		}
	}
	return processItems
}

//
func (top *Top) GetProcessList() ([]ProcessItem, error) {
	top.accessMutes.Lock()
	defer top.accessMutes.Unlock()
	return top.processItems, nil
}

func (top *Top) KillProcess(pid int) error {
	//this is service is need lock
	return nil
}

func TestGetProcessList(t *testing.T) {
	top := new(Top)
	top.GetProcessList()
}
