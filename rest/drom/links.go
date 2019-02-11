package drom

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
██╗     ██╗███╗   ██╗██╗  ██╗███████╗
██║     ██║████╗  ██║██║ ██╔╝██╔════╝
██║     ██║██╔██╗ ██║█████╔╝ ███████╗
██║     ██║██║╚██╗██║██╔═██╗ ╚════██║
███████╗██║██║ ╚████║██║  ██╗███████║
╚══════╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
*/
//This page is the part that shows the links that the queries will use.
//It is designed in such a way that the administration is easy.

var (
	dromdeviceconfiguration = "dromdeviceconfiguration/"
	push                    = "push/"
	delete                  = "delete/"
	list                    = "list"
	add                     = "add"
)

// https://api.ardich.com/api/v3/dromdeviceconfiguration/push/{YOUR_DEVICE_ID}
func sendDromLink(setDeviceID string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + dromdeviceconfiguration + push + setDeviceID

	return u.String()
}

// https://api.ardich.com/api/v3/dromdeviceconfiguration/delete/{YOUR_DEVICE_ID}
func deleteDromLink(setDeviceID string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + dromdeviceconfiguration + delete + setDeviceID

	return u.String()
}

// https://api.ardich.com:443/api/v3/dromconfiguration/list
func configurationListLink() string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + dromdeviceconfiguration + list

	return u.String()
}

// https://api.ardich.com:443/api/v3/dromdeviceconfiguration/add
func addDeviceLink() string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + dromdeviceconfiguration + add

	return u.String()
}

// { "configurationId": "xxxx", "deviceId": "yyyy"}
func addDeviceBody(configurationID string, deviceID string) string {
	var addDeviceBodyVar addDeviceBodyJSON
	addDeviceBodyVar.ConfigurationID = configurationID
	addDeviceBodyVar.DeviceID = deviceID
	jsonConvert, _ := json.Marshal(addDeviceBodyVar)

	return string(jsonConvert)
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
