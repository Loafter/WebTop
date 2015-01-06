package main

import "io/ioutil"
import "strconv"
import "regexp"
import "errors"
import "time"
import "sync"

//process info sructure
type ProcessItem struct {
	Pid    int
	Name   string
	User   string
	Cpu    float32
	Memory int
}

type Timeval struct {
	Sec  int32
	Usec int32
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
	statFileData, err := ioutil.ReadFile("/proc/" + strconv.Itoa(pid) + "/stat")
	if err != nil {
		return 0, errors.New("error: problem with read proc filesystem \n " + err.Error())
	}
	statFileStr := string(statFileData)
	cpuTimeReg := regexp.MustCompile("\\d+")
	cpuStatField := cpuTimeReg.FindAllString(statFileStr, -1)
	utime, _ := strconv.Atoi(cpuStatField[11])
	stime, _ := strconv.Atoi(cpuStatField[12])
	cutime, _ := strconv.Atoi(cpuStatField[13])
	cstime, _ := strconv.Atoi(cpuStatField[14])
	sumTime := int64(utime + stime + cutime + cstime)
	return sumTime, nil
}

func (top *Top) getTicksMap(pids []int) (map[int]int64, int64) {
	var sum int64
	ticksMap := make(map[int]int64)
	for _, element := range pids {
		ticks, err := top.getTicksbyPid(element)
		if err != nil {
			ticksMap[element] = 0
		}
		ticksMap[element] = ticks
		sum += ticks
	}
	return ticksMap, sum
}
func (top *Top) collectInfo() error {
	for {
		pids, err := top.getAllPids()
		if err != nil {
			return errors.New("error: problem with read proc filesystem \n " + err.Error())
		}
		StartTicks, sumOldTick := top.getTicksMap(pids)
		time.Sleep(1000 * time.Millisecond)
		pids, _ = top.getAllPids()
		EndTicks, sumNewTick := top.getTicksMap(pids)
		top.accessMutes.Lock()
		top.processItems = top.fillProcessInfo(StartTicks, EndTicks, sumNewTick-sumOldTick)
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

func (top *Top) fillProcessInfo(oldTicks map[int]int64, newTicks map[int]int64, sumTicks int64) []ProcessItem {
	processItems := []ProcessItem{}
	for i, newTickVal := range newTicks {
		statFileData, err := ioutil.ReadFile("/proc/" + strconv.Itoa(i) + "/status")
		if err == nil {
			processItem := ProcessItem{}
			processItem.Pid = i
			statFileStr := string(statFileData)
			regName := regexp.MustCompile("Name:\t(.*)\n")
			processItem.Name = regName.FindAllStringSubmatch(statFileStr, -1)[0][1]
			regUid := regexp.MustCompile("Uid:\t(\\w+)")
			processItem.User = regUid.FindAllStringSubmatch(statFileStr, -1)[0][1]
			regMem := regexp.MustCompile("VmRSS:\\s+(\\d+)")
			if regMem.MatchString(statFileStr) {
				intMem, _ := strconv.Atoi(regMem.FindAllStringSubmatch(statFileStr, -1)[0][1])
				processItem.Memory = intMem
			}
			oldTickVal, exist := oldTicks[i]
			if exist {

				processItem.Cpu = float32(float64(newTickVal-oldTickVal) / float64(sumTicks))

			} else {
				processItem.Cpu = 0
			}
			if processItem.Cpu != 0 {
				processItems = append(processItems, processItem)
			}
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
