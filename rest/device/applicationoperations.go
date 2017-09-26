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
	var setAdres string

	if setOperations == StartApp {
		setAdres = startAppLink(setDeviceCode)
	} else if setOperations == StopApp {
		setAdres = stopAppLink(setDeviceCode)
	}

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
