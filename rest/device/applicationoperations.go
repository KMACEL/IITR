package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 █████╗ ██████╗ ██████╗ ███████╗███████╗
██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝
███████║██████╔╝██████╔╝███████╗███████╗
██╔══██║██╔═══╝ ██╔═══╝ ╚════██║╚════██║
██║  ██║██║     ██║     ███████║███████║
╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚══════╝
*/

//AppSS is
func (d Device) AppSS(setOperations int, setDeviceCode string, setApplicationPackage string, vasualFlag bool) string {
	var setAddress string

	if setOperations == StartApp {
		setAddress = startAppLink(setDeviceCode)
	} else if setOperations == StopApp {
		setAddress = stopAppLink(setDeviceCode)
	}

	setBody := applicationOperationsBodyLink(setApplicationPackage)

	query, _ := queryVariable.PostQuery(setAddress, setBody, contentTypeJSON(), vasualFlag)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responseMessageCodeJSONVariable)
			return responseMessageCodeJSONVariable.Response
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}

/*
 ██████╗██╗     ███████╗ █████╗ ██████╗          █████╗ ██████╗ ██████╗         ██████╗  █████╗ ████████╗ █████╗
██╔════╝██║     ██╔════╝██╔══██╗██╔══██╗        ██╔══██╗██╔══██╗██╔══██╗        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗
██║     ██║     █████╗  ███████║██████╔╝        ███████║██████╔╝██████╔╝        ██║  ██║███████║   ██║   ███████║
██║     ██║     ██╔══╝  ██╔══██║██╔══██╗        ██╔══██║██╔═══╝ ██╔═══╝         ██║  ██║██╔══██║   ██║   ██╔══██║
╚██████╗███████╗███████╗██║  ██║██║  ██║        ██║  ██║██║     ██║             ██████╔╝██║  ██║   ██║   ██║  ██║
 ╚═════╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝     ╚═╝             ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝
*/

// ClearAppData is
func (d Device) ClearAppData(setDeviceCode string, setApplicationPackage string, vasualFlag bool) string {
	setAdres := "https://api.ardich.com/api/v3/device/" + setDeviceCode + "/apps/clearappdata"
	setBody := applicationOperationsBodyLink(setApplicationPackage)

	query, _ := queryVariable.PostQuery(setAdres, setBody, contentTypeJSON(), vasualFlag)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responseMessageCodeJSONVariable)
			return responseMessageCodeJSONVariable.Response
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}
