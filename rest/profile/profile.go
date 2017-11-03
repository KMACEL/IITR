package profile

import (
	"encoding/json"

	"fmt"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/workingset"
)

//PushMode is
func (p Profile) PushMode(workingset string, setMode string, setPolicy string) string {

	setAddress := pushProfileLink(setMode, workingset)
	setBody := "{\"defaultPolicy\":{\"code\": \"" + setPolicy + "\"}}"

	header := make(map[string]string)
	header["content-type"] = "application/json"

	query, _ := queryVariable.PostQuery(setAddress, setBody, header, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &modeResponseJSONVariable)
			return modeResponseJSONVariable[0].Status
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}

//PushMode is
func (p Profile) PushModeAuto(setMode string, setPolicy string, devicesID ...string) string {
	var (
		workingsets workingset.Workingset
		devices     device.Device
	)

	workingsetKey := workingsets.CreateWorkingset()

	for _, deviceID := range devicesID {
		workingsets.AddDeviceWorkingSet(workingsetKey, devices.DeviceID2Code(deviceID))
	}

	fmt.Println("Workingset Device List : ", workingsets.GetWorkingsetDevices(workingsetKey))

	setAddress := pushProfileLink(setMode, workingsetKey)

	query, _ := queryVariable.PostQuery(setAddress, pushProfileBody(setPolicy), contentTypeJSON(), rest.Visible)
	fmt.Println("query : ", string(query))
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &modeResponseJSONVariable)
			return modeResponseJSONVariable[0].Status
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}
