package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗           ██╗██████╗     ██████╗      ██████╗ ██████╗ ██████╗ ███████╗
██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝    ██╗    ██║██╔══██╗    ╚════██╗    ██╔════╝██╔═══██╗██╔══██╗██╔════╝
██║  ██║█████╗  ██║   ██║██║██║     █████╗      ╚═╝    ██║██║  ██║     █████╔╝    ██║     ██║   ██║██║  ██║█████╗
██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝      ██╗    ██║██║  ██║    ██╔═══╝     ██║     ██║   ██║██║  ██║██╔══╝
██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗    ╚═╝    ██║██████╔╝    ███████╗    ╚██████╗╚██████╔╝██████╔╝███████╗
╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝           ╚═╝╚═════╝     ╚══════╝     ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝
*/

//DeviceID2Code is
func (d Device) DeviceID2Code(deviceID string) string {
	setQueryAddress := deviceID2CodeLink(deviceID)
	query, _ := q.GetQuery(setQueryAddress, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &locationAllJSONVariable)
			return locationAllJSONVariable[0].DeviceCode
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil
}

/*
██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗    ██╗██████╗         ██████╗          ██████╗ ██████╗ ██████╗ ███████╗    ██╗      ██████╗  ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗
██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝    ██║██╔══██╗        ╚════██╗        ██╔════╝██╔═══██╗██╔══██╗██╔════╝    ██║     ██╔═══██╗██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║
██║  ██║█████╗  ██║   ██║██║██║     █████╗      ██║██║  ██║         █████╔╝        ██║     ██║   ██║██║  ██║█████╗      ██║     ██║   ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║
██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝      ██║██║  ██║        ██╔═══╝         ██║     ██║   ██║██║  ██║██╔══╝      ██║     ██║   ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║
██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗    ██║██████╔╝        ███████╗        ╚██████╗╚██████╔╝██████╔╝███████╗    ███████╗╚██████╔╝╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║
╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝    ╚═╝╚═════╝         ╚══════╝         ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝    ╚══════╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
*/

//DeviceID2CodeLocation is
func (d Device) DeviceID2CodeLocation(deviceID string) (string, string, string) {
	setQueryAddress := deviceID2CodeLink(deviceID)
	query, _ := q.GetQuery(setQueryAddress, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &locationAllJSONVariable)
			return locationAllJSONVariable[0].DeviceCode, locationAllJSONVariable[0].Latitude, locationAllJSONVariable[0].Longitude
		}
		return rest.ResponseNotFound, rest.ResponseNotFound, rest.ResponseNotFound
	}
	return rest.ResponseNil, rest.ResponseNil, rest.ResponseNil
}

/*
██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗            ██████╗ ██████╗ ██████╗ ███████╗        ██████╗         ██╗██████╗
██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝    ██╗    ██╔════╝██╔═══██╗██╔══██╗██╔════╝        ╚════██╗        ██║██╔══██╗
██║  ██║█████╗  ██║   ██║██║██║     █████╗      ╚═╝    ██║     ██║   ██║██║  ██║█████╗           █████╔╝        ██║██║  ██║
██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝      ██╗    ██║     ██║   ██║██║  ██║██╔══╝          ██╔═══╝         ██║██║  ██║
██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗    ╚═╝    ╚██████╗╚██████╔╝██████╔╝███████╗        ███████╗        ██║██████╔╝
╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝            ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝        ╚══════╝        ╚═╝╚═════╝
*/

//DeviceCode2ID is
func (d Device) DeviceCode2ID(deviceCode string) string {
	setQueryAddress := deviceCode2IDLink(deviceCode)
	query, _ := q.GetQuery(setQueryAddress, rest.Invisible)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			var codeJSON PresenceInfoJSON
			json.Unmarshal(query, &codeJSON)
			return codeJSON.DeviceID
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil
}
