package main

import "io/ioutil"
import "strconv"
import "regexp"
import "errors"
import "time"
import "os"
import "log"

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
	processItems    []ProcessItem
	collectInfoJob  BatchJob
	lastRequestTime time.Time
}

func (top *Top) StartCollectInfo() error {
	top.collectInfoJob.Job = top.collectInfo
	top.lastRequestTime = time.Now()
	err := top.collectInfoJob.Start()
	if err != nil {
		stErr := "error: Can't start collect info job"
		log.Println(stErr)
		return errors.New(stErr)
	}
	return nil
}
func (top *Top) StopCollectInfo() error {
	err := top.collectInfoJob.Stop()
	if err != nil {
		stErr := "error: Can't stop  collect info job"
		log.Println(stErr)
		return errors.New(stErr)
	}
	return nil
}

func (top *Top) getTicksbyPid(pid int) (int64, error) {
	statFileData, err := ioutil.ReadFile("/proc/" + strconv.Itoa(pid) + "/stat")
	if err != nil {
		stErr := "error: problem with read proc filesystem"
		log.Println(stErr)
		return 0, errors.New(stErr)
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

func (top *Top) getTicksProcessor() (int64, error) {
	statFileData, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		stErr := "error: problem with read proc filesystem"
		log.Println(stErr)
		return 0, errors.New(stErr)
	}
	statFileStr := string(statFileData)
	cpuTimeReg := regexp.MustCompile("\\d+")
	cpuStatField := cpuTimeReg.FindAllString(statFileStr, -1)[:9]
	var sumTime int64
	for _, cpuField := range cpuStatField {
		tim, _ := strconv.Atoi(cpuField)
		sumTime += int64(tim)
	}
	return sumTime, nil
}

func (top *Top) getTicksMap(pids []int) map[int]int64 {
	ticksMap := make(map[int]int64)
	for _, element := range pids {
		ticks, err := top.getTicksbyPid(element)
		if err != nil {
			ticksMap[element] = 0
		}
		ticksMap[element] = ticks
	}
	return ticksMap
}

func (top *Top) collectInfo() {
	for {
		pids, err := top.getAllPids()
		if err != nil {
			stErr := "error: problem with read proc filesystem"
			log.Println(stErr)
			return //errors.New("error: problem with read proc filesystem")
		}
		StartTicks := top.getTicksMap(pids)
		sumOldTick, _ := top.getTicksProcessor()
		time.Sleep(1000 * time.Millisecond)
		pids, _ = top.getAllPids()
		EndTicks := top.getTicksMap(pids)
		sumNewTick, _ := top.getTicksProcessor()

		top.processItems = top.fillProcessInfo(StartTicks, EndTicks, sumNewTick-sumOldTick)
		secondLastRequest := time.Now().Sub(top.lastRequestTime)
		if secondLastRequest.Seconds() > 5 {
			log.Println("info: sleep collect top info job")
			top.collectInfoJob.Stop()
		}
	}
	return
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
				stErr := "info: can't messure cpu info for process "
				log.Println(stErr, processItem.Name)
				processItem.Cpu = 0
			}
			processItems = append(processItems, processItem)
		}
	}
	return processItems
}

//
func (top *Top) GetProcessList() ([]ProcessItem, error) {
	if !top.collectInfoJob.IsRunning() {
		log.Println("info: start collect top info job")
		top.collectInfoJob.Start()
	}
	top.lastRequestTime = time.Now()
	return top.processItems, nil
}

func (top *Top) KillProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		stErr := "info: can't find process pid=" + string(pid)
		log.Println(stErr)
		return errors.New(stErr)
	}
	err = process.Kill()
	if err != nil {
		stErr := "info: can't kill process pid=" + string(pid)
		log.Println(stErr)
		return errors.New(stErr)
	}
	return nil
}
