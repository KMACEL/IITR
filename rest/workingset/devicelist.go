package workingset

import (
	"github.com/KMACEL/IITR/rest"
	"encoding/json"
)

func (w Workingset) GetWorkingsetDevices(workingsetKey string) []string {
	setQueryAdress := getWorkingsetDevicesLink(workingsetKey)
	query, _ := rest.Query{}.GetQuery(setQueryAdress, false)
	var deviceList []string

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &workingsetDevicesJSONVariable)
			for _, deviceID := range workingsetDevicesJSONVariable.Content {
				deviceList=append(deviceList, deviceID.DeviceID)
			}
			return deviceList

		}
		deviceList=append(deviceList, rest.ResponseNotFound)
		return deviceList
	}
	deviceList=append(deviceList, rest.ResponseNil)
	return deviceList
}
