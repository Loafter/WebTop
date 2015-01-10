package main

import "errors"
import "io/ioutil"
import "net/http"
import "encoding/json"
import "log"

type JsonFabric struct {
}

//parse json message
func (fabric *JsonFabric) ProduceJsonRequest(request *http.Request) (Request, error) {
	bodyData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		stErr := "error: Can't read request body"
		log.Println(stErr)
		return nil, errors.New(stErr)
	}
	defer request.Body.Close()

	var basicRequest BasicRequest
	err = json.Unmarshal(bodyData, &basicRequest)
	if err != nil {
		stErr := "error: Can't parse basic data"
		log.Println(stErr)
		return nil, errors.New(stErr)
	}
	switch basicRequest.Type {
	case ServiceStatus:
		var serviceStateRequest ServiceStateRequest
		err := json.Unmarshal(bodyData, &serviceStateRequest)
		if err != nil {
			stErr := "error: Can't parse service state request"
			log.Println(stErr)
			return nil, errors.New("error: Can't parse service state request")
		}
		return serviceStateRequest, nil
	case SystemMonitor:
		var systemStateRequest SystemStateRequest
		err := json.Unmarshal(bodyData, &systemStateRequest)
		if err != nil {
			stErr := "error: Can't parse system state request"
			log.Println(stErr)
			return nil, errors.New(stErr)
		}
		return systemStateRequest, nil
	case TopProcess:
		var topRequest TopRequest
		err := json.Unmarshal(bodyData, &topRequest)
		if err != nil {
			stErr := "error: Can't parse top request"
			log.Println(stErr)
			return nil, errors.New(stErr)

		}
		return topRequest, nil
	case KillProcess:
		var killRequest KillRequest
		err := json.Unmarshal(bodyData, &killRequest)
		if err != nil {
			stErr := "error: Can't parse top request"
			log.Println(stErr)
			return nil, errors.New(stErr)
		}
		return killRequest, nil

	}

	return nil, errors.New("error: Unknown request type")
}
