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
func (d Device) LocationMap(setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := locationMapLink()
	query, _ := q.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &locationJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}

/*
██╗      ██████╗  ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗        ██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗
██║     ██╔═══██╗██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝
██║     ██║   ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║        ██║  ██║█████╗  ██║   ██║██║██║     █████╗
██║     ██║   ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║        ██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝
███████╗╚██████╔╝╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║        ██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗
╚══════╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝
*/

//LocationDevice is
func (d Device) LocationDevice(deviceID string) (string, string) {
	setQueryAddress := deviceID2CodeLink(deviceID)
	query, _ := q.GetQuery(setQueryAddress, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &locationAllJSONVariable)
			return locationAllJSONVariable[0].Latitude, locationAllJSONVariable[0].Longitude
		}
		return rest.ResponseNotFound, rest.ResponseNotFound
	}
	return rest.ResponseNil, rest.ResponseNil
}
