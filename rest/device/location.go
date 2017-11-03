package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
██╗      ██████╗  ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗        ███╗   ███╗ █████╗ ██████╗
██║     ██╔═══██╗██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║        ████╗ ████║██╔══██╗██╔══██╗
██║     ██║   ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║        ██╔████╔██║███████║██████╔╝
██║     ██║   ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║        ██║╚██╔╝██║██╔══██║██╔═══╝
███████╗╚██████╔╝╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║        ██║ ╚═╝ ██║██║  ██║██║
╚══════╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝
*/

//LocationMap is
func (d Device) LocationMap(setUnMarshal bool, vasualFlag bool) []byte {
	setQueryAdress := locationMapLink()
	query, _ := rest.Query{}.GetQuery(setQueryAdress, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &locationJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}

//LocationDevice is
func (d Device) LocationDevice(deviceID string) (string,string) {
	setQueryAdress := deviceID2CodeLink(deviceID)
	query, _ := rest.Query{}.GetQuery(setQueryAdress, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &locationAllJSONVariable)
			return locationAllJSONVariable[0].Latitude,locationAllJSONVariable[0].Longitude
		}
		return rest.ResponseNotFound,rest.ResponseNotFound
	}
	return rest.ResponseNil,rest.ResponseNil
}