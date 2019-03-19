package workingset

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗██╗███╗   ██╗ ██████╗ ███████╗███████╗████████╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝██║████╗  ██║██╔════╝ ██╔════╝██╔════╝╚══██╔══╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██║ █╗ ██║██║   ██║██████╔╝█████╔╝ ██║██╔██╗ ██║██║  ███╗███████╗█████╗     ██║
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██║███╗██║██║   ██║██╔══██╗██╔═██╗ ██║██║╚██╗██║██║   ██║╚════██║██╔══╝     ██║
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗██║██║ ╚████║╚██████╔╝███████║███████╗   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝╚══════╝   ╚═╝
*/

//CreateWorkingset is
func (w Workingset) CreateWorkingset() string {
	setAddress := createWorkingsetLink()
	query, _ := queryVariable.PostQuery(setAddress, "", contentTypeJSON(), rest.Invisible)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &workingsetJSONVariable)
			return workingsetJSONVariable.Code
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil
}

/*
 █████╗ ██████╗ ██████╗         ██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗        ██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗██╗███╗   ██╗ ██████╗ ███████╗███████╗████████╗
██╔══██╗██╔══██╗██╔══██╗        ██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝        ██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝██║████╗  ██║██╔════╝ ██╔════╝██╔════╝╚══██╔══╝
███████║██║  ██║██║  ██║        ██║  ██║█████╗  ██║   ██║██║██║     █████╗          ██║ █╗ ██║██║   ██║██████╔╝█████╔╝ ██║██╔██╗ ██║██║  ███╗███████╗█████╗     ██║
██╔══██║██║  ██║██║  ██║        ██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝          ██║███╗██║██║   ██║██╔══██╗██╔═██╗ ██║██║╚██╗██║██║   ██║╚════██║██╔══╝     ██║
██║  ██║██████╔╝██████╔╝        ██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗        ╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗██║██║ ╚████║╚██████╔╝███████║███████╗   ██║
╚═╝  ╚═╝╚═════╝ ╚═════╝         ╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝         ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝╚══════╝   ╚═╝
*/

//AddDeviceWorkingSet is
func (w Workingset) AddDeviceWorkingSet(workingsetKey string, deviceCode ...string) string {
	setAdres := addDeviceWorkingSetLink(workingsetKey)
	var addDeviceWorkingSetBodyVar addDeviceWorkingSetBody
	addDeviceWorkingSetBodyVar.DeviceList = deviceCode
	setBody, _ := json.Marshal(addDeviceWorkingSetBodyVar.DeviceList)

	query, _ := queryVariable.PutQuery(setAdres, string(setBody), contentTypeJSON(), rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			return string(query)
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}
