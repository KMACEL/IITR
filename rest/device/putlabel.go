package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
)

/*
███████╗███████╗████████╗        ██╗      █████╗ ██████╗ ███████╗██╗
██╔════╝██╔════╝╚══██╔══╝        ██║     ██╔══██╗██╔══██╗██╔════╝██║
███████╗█████╗     ██║           ██║     ███████║██████╔╝█████╗  ██║
╚════██║██╔══╝     ██║           ██║     ██╔══██║██╔══██╗██╔══╝  ██║
███████║███████╗   ██║           ███████╗██║  ██║██████╔╝███████╗███████╗
╚══════╝╚══════╝   ╚═╝           ╚══════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝
*/
//	device.Device{}.SetLabel("{YOUR_DEVICE_ID}", "{LABEL_NAME}", {true|false})

// SetLabel is
func (d Device) SetLabel(setDeviceID string, labelName string, visualFlag bool) string {
	setAddress := setLabelLink(setDeviceID)
	var setBodyJSON SetLabelBodyJSON
	setBodyJSON.Label = labelName
	setBody, _ := json.Marshal(setBodyJSON)

	query, removeError := q.PutQuery(setAddress, string(setBody), contentTypeJSON(), false)
	errc.ErrorCenter(removeApplicationErrorTag, removeError)

	if query != nil {
		json.Unmarshal(query, &responseMessageCodeJSONVariable)
		if visualFlag {
			return "Device ID : " + setDeviceID + " Label Name : " + labelName + "Status : " + responseMessageCodeJSONVariable.Response
		}
	} else {
		if visualFlag {
			return "Device ID : " + setDeviceID + " Label Name : " + labelName + "Status : False"
		}
	}
	return ""
}
