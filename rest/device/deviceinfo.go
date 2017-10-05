package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
██████╗ ██████╗ ███████╗███████╗███████╗███╗   ██╗ ██████╗███████╗        ██╗███╗   ██╗███████╗ ██████╗
██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝████╗  ██║██╔════╝██╔════╝        ██║████╗  ██║██╔════╝██╔═══██╗
██████╔╝██████╔╝█████╗  ███████╗█████╗  ██╔██╗ ██║██║     █████╗          ██║██╔██╗ ██║█████╗  ██║   ██║
██╔═══╝ ██╔══██╗██╔══╝  ╚════██║██╔══╝  ██║╚██╗██║██║     ██╔══╝          ██║██║╚██╗██║██╔══╝  ██║   ██║
██║     ██║  ██║███████╗███████║███████╗██║ ╚████║╚██████╗███████╗        ██║██║ ╚████║██║     ╚██████╔╝
╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═══╝ ╚═════╝╚══════╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

//PresenceInfo is
func (d Device) PresenceInfo(setDeviceID string, setUnMarshal bool, vasualFlag bool) []byte {
	setQueryAdress := presenceInfoLink(setDeviceID)

	query, _ := rest.Query{}.GetQuery(setQueryAdress, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &presenceInfoJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}

/*
 █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗        ██╗███╗   ██╗███████╗ ██████╗
██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║        ██║████╗  ██║██╔════╝██╔═══██╗
███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║        ██║██╔██╗ ██║█████╗  ██║   ██║
██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║        ██║██║╚██╗██║██╔══╝  ██║   ██║
██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║        ██║██║ ╚████║██║     ╚██████╔╝
╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

//ApplicationInfo is
func (d Device) ApplicationInfo(setDeviceID string, setUnMarshal bool, vasualFlag bool) []byte {
	setQueryAdress := applicationInfoLink(setDeviceID)

	query, _ := rest.Query{}.GetQuery(setQueryAdress, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &applicationInfoJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}

/*
 ██████╗ ███████╗        ██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗        ██╗███╗   ██╗███████╗ ██████╗
██╔═══██╗██╔════╝        ██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝        ██║████╗  ██║██╔════╝██╔═══██╗
██║   ██║███████╗        ██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗          ██║██╔██╗ ██║█████╗  ██║   ██║
██║   ██║╚════██║        ██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝          ██║██║╚██╗██║██╔══╝  ██║   ██║
╚██████╔╝███████║        ██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗        ██║██║ ╚████║██║     ╚██████╔╝
 ╚═════╝ ╚══════╝        ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

// OSProfileInfo is
func (d Device) OSProfileInfo(setDeviceID string, setUnMarshal bool, vasualFlag bool) []byte {
	setQueryAdress := osProfileInfoLink(setDeviceID)

	query, _ := rest.Query{}.GetQuery(setQueryAdress, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &osProfileInfoJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}
