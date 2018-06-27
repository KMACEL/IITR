package workingset

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗██╗███╗   ██╗ ██████╗ ███████╗███████╗████████╗        ██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗███████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝██║████╗  ██║██╔════╝ ██╔════╝██╔════╝╚══██╔══╝        ██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝██╔════╝
██║  ███╗█████╗     ██║           ██║ █╗ ██║██║   ██║██████╔╝█████╔╝ ██║██╔██╗ ██║██║  ███╗███████╗█████╗     ██║           ██║  ██║█████╗  ██║   ██║██║██║     █████╗  ███████╗
██║   ██║██╔══╝     ██║           ██║███╗██║██║   ██║██╔══██╗██╔═██╗ ██║██║╚██╗██║██║   ██║╚════██║██╔══╝     ██║           ██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝  ╚════██║
╚██████╔╝███████╗   ██║           ╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗██║██║ ╚████║╚██████╔╝███████║███████╗   ██║           ██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗███████║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝╚══════╝   ╚═╝           ╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝╚══════╝
*/

// GetWorkingsetDevices returns all devices in the workingset
func (w Workingset) GetWorkingsetDevices(workingsetKey string) []string {
	setQueryAddress := getWorkingsetDevicesLink(workingsetKey)
	query, _ := rest.Query{}.GetQuery(setQueryAddress, false)
	var deviceList []string

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &workingsetDevicesJSONVariable)
			for _, deviceID := range workingsetDevicesJSONVariable.Content {
				deviceList = append(deviceList, deviceID.DeviceID)
			}
			return deviceList
		}
		deviceList = append(deviceList, rest.ResponseNotFound)
		return deviceList
	}
	deviceList = append(deviceList, rest.ResponseNil)
	return deviceList
}
