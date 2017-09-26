package workingset

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

//CreateWorkingset is
func (w Workingset) CreateWorkingset() string {
	setAdres := CreateWorkingsetLink()
	header := make(map[string]string)
	header["content-type"] = "application/json"

	query, _ := queryVariable.PostQuery(setAdres, "", header, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &workingsetJSONVariable)
			return workingsetJSONVariable.Code
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}

//AddDeviceWorkingSet is
func (w Workingset) AddDeviceWorkingSet(workingset string, deviceCode ...string) string {
	setAdres := AddDeviceWorkingSetLink(workingset)
	header := make(map[string]string)
	header["content-type"] = "application/json"
	var deviceList string
	for i, code := range deviceCode {
		if i < len(deviceCode)-1 {
			deviceList = ",\"" + code + "\"" + deviceList
		} else {
			deviceList = "\"" + code + "\"" + deviceList
		}

	}
	setBody := "[" + deviceList + "]"
	query, _ := queryVariable.PutQuery(setAdres, setBody, header, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			return string(query)
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}
