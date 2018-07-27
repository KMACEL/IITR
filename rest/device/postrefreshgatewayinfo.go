package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
██████╗ ███████╗███████╗██████╗ ███████╗███████╗██╗  ██╗         ██████╗  █████╗ ████████╗███████╗██╗    ██╗ █████╗ ██╗   ██╗        ██╗███╗   ██╗███████╗ ██████╗
██╔══██╗██╔════╝██╔════╝██╔══██╗██╔════╝██╔════╝██║  ██║        ██╔════╝ ██╔══██╗╚══██╔══╝██╔════╝██║    ██║██╔══██╗╚██╗ ██╔╝        ██║████╗  ██║██╔════╝██╔═══██╗
██████╔╝█████╗  █████╗  ██████╔╝█████╗  ███████╗███████║        ██║  ███╗███████║   ██║   █████╗  ██║ █╗ ██║███████║ ╚████╔╝         ██║██╔██╗ ██║█████╗  ██║   ██║
██╔══██╗██╔══╝  ██╔══╝  ██╔══██╗██╔══╝  ╚════██║██╔══██║        ██║   ██║██╔══██║   ██║   ██╔══╝  ██║███╗██║██╔══██║  ╚██╔╝          ██║██║╚██╗██║██╔══╝  ██║   ██║
██║  ██║███████╗██║     ██║  ██║███████╗███████║██║  ██║        ╚██████╔╝██║  ██║   ██║   ███████╗╚███╔███╔╝██║  ██║   ██║           ██║██║ ╚████║██║     ╚██████╔╝
╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝         ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝ ╚══╝╚══╝ ╚═╝  ╚═╝   ╚═╝           ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

// 	device.Device{}.RefreshGatewayInfo(device.Device{}.DeviceID2Code("00:09:4c:4f:8d:53@iotigniteagent"), device.ApplicationInfo)

// RefreshGatewayInfo is contains commands that request information from the device.
// specificParameter :
//			OSProfile
//			BatteryInfo
//			ModiverseInfo
//			NetworkInfo
//			RootedInfo
//			ProcessInfo
//			StorageInfo
//			UsageInfo
//			ApplicationInfo
//			LocationInfo
//			DeviceNodeInventory
//			DeviceFlowInventory
//			DeviceConfigInventory
func (d Device) RefreshGatewayInfo(deviceCode string, specificParameter ...string) string {
	setAddress := refreshGatewayInfoLink(deviceCode, specificParameter...)
	query, _ := q.PostQuery(setAddress, "", contentTypeJSON(), true)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responseMessageCodeJSONVariable)
			return responseMessageCodeJSONVariable.Response
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil
}
