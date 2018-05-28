package operations

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

/*
 █████╗ ██████╗ ██████╗         ██╗ ██████╗ ████████╗        ██╗      █████╗ ██████╗ ███████╗██╗
██╔══██╗██╔══██╗██╔══██╗        ██║██╔═══██╗╚══██╔══╝        ██║     ██╔══██╗██╔══██╗██╔════╝██║
███████║██║  ██║██║  ██║        ██║██║   ██║   ██║           ██║     ███████║██████╔╝█████╗  ██║
██╔══██║██║  ██║██║  ██║        ██║██║   ██║   ██║           ██║     ██╔══██║██╔══██╗██╔══╝  ██║
██║  ██║██████╔╝██████╔╝        ██║╚██████╔╝   ██║           ███████╗██║  ██║██████╔╝███████╗███████╗
╚═╝  ╚═╝╚═════╝ ╚═════╝         ╚═╝ ╚═════╝    ╚═╝           ╚══════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝
*/

//AddIOTLabel is
type AddIOTLabel struct {
	devices   device.Device
	Label     string
	LabelType string
	DeviceID  []string
}

//Start is
func (a AddIOTLabel) Start() {

	var (
		iotLabel             device.AddIOTLabelJSON
		locationJSONVariable device.LocationJSON
	)

	if a.DeviceID == nil {
		query := a.devices.LocationMap(rest.NOMarshal, rest.Invisible)

		json.Unmarshal(query, &locationJSONVariable)

		for _, deviceCoding := range locationJSONVariable.Extras {
			var iotLabelDeviceID device.AddIOTLabelDeviceIDJSON
			iotLabelDeviceID.DeviceID = deviceCoding.DeviceID
			iotLabel.DeviceNodeSensors = append(iotLabel.DeviceNodeSensors, iotLabelDeviceID)
		}
	} else {
		for _, deviceID := range a.DeviceID {
			var iotLabelDeviceID device.AddIOTLabelDeviceIDJSON
			iotLabelDeviceID.DeviceID = deviceID
			iotLabel.DeviceNodeSensors = append(iotLabel.DeviceNodeSensors, iotLabelDeviceID)
		}
	}

	if a.LabelType == "" {
		iotLabel.LabelType = "DEVICE"
	}

	if a.Label == "" {
		iotLabel.Label = "DEVICES"
	} else {
		iotLabel.Label = a.Label
	}
	a.devices.AddIOTLabel(iotLabel, true)
}
