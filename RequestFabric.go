package main

import "errors"
import "io/ioutil"
import "net/http"
import "encoding/json"

type JsonFabric struct {
}

//parse json message
func (fabric *JsonFabric) ProduceJsonRequest(request *http.Request) (Request, error) {
	bodyData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errors.New("Can't read request body \n " + err.Error())
	}
	defer request.Body.Close()

	var basicRequest BasicRequest
	err = json.Unmarshal(bodyData, &basicRequest)
	if err != nil {
		return nil, errors.New("error: Can't parse basic data \n " + err.Error())
	}
	switch basicRequest.Type {
	case ServiceStatus:
		var serviceStateRequest ServiceStateRequest
		err := json.Unmarshal(bodyData, &serviceStateRequest)
		if err != nil {
			return nil, errors.New("error: Can't parse service state request \n " + err.Error())
		}
		return serviceStateRequest, nil
	case SystemMonitor:
		var systemStateRequest SystemStateRequest
		err := json.Unmarshal(bodyData, &systemStateRequest)
		if err != nil {
			return nil, errors.New("error: Can't parse system state request \n " + err.Error())
		}
		return systemStateRequest, nil
	case TopProcess:
		var topRequest TopRequest
		err := json.Unmarshal(bodyData, &topRequest)
		if err != nil {
			return nil, errors.New("error: Can't parse system state request \n " + err.Error())
		}
		return topRequest, nil
	}

	return nil, errors.New("error: Unknown request type")
}
