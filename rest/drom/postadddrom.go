package drom

import (
	"github.com/KMACEL/IITR/rest"
)

/*
 █████╗ ██████╗ ██████╗         ██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗
██╔══██╗██╔══██╗██╔══██╗        ██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝
███████║██║  ██║██║  ██║        ██║  ██║█████╗  ██║   ██║██║██║     █████╗
██╔══██║██║  ██║██║  ██║        ██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝
██║  ██║██████╔╝██████╔╝        ██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗
╚═╝  ╚═╝╚═════╝ ╚═════╝         ╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝
*/

/*
Example :
    drom.Drom{}.AddDevice("{DeviceID}", "{ModeName}")
*/

// AddDevice is
func (d Drom) AddDevice(deviceID string, configurationName string) string {

	setAddress := addDeviceLink()
	query, _ := rest.Query{}.PostQuery(setAddress, addDeviceBody(d.GetDromConfiguration(configurationName, false, false), deviceID), contentTypeJSON(), true)
	if query != nil {

		if string(query) != rest.ResponseNotFound {
			//json.Unmarshal(query, &responseMessageCodeJSONVariable)
			return "ok"
		}
		return "rest.ResponseNotFound"
	}
	return rest.ResponseNil
}
