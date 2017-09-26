package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗        ██╗      ██████╗  ██████╗         ██╗     ██╗███████╗████████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║     ██╔═══██╗██╔════╝         ██║     ██║██╔════╝╚══██╔══╝
██║  ███╗█████╗     ██║           ██║     ██║   ██║██║  ███╗        ██║     ██║███████╗   ██║
██║   ██║██╔══╝     ██║           ██║     ██║   ██║██║   ██║        ██║     ██║╚════██║   ██║
╚██████╔╝███████╗   ██║           ███████╗╚██████╔╝╚██████╔╝        ███████╗██║███████║   ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚══════╝ ╚═════╝  ╚═════╝         ╚══════╝╚═╝╚══════╝   ╚═╝
*/

// GetLogList is
func (d Device) GetLogList(setDeviceCode string, setUnMarshal bool, vasualFlag bool) []byte {
	setQueryAdress := getLogListLink(setDeviceCode)
	query, _ := queryVariable.GetQuery(setQueryAdress, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &logListJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}

/*
 ██████╗ ███████╗████████╗        ██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗        ██╗      ██████╗  ██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝        ██║     ██╔═══██╗██╔════╝
██║  ███╗█████╗     ██║           ██║  ██║█████╗  ██║   ██║██║██║     █████╗          ██║     ██║   ██║██║  ███╗
██║   ██║██╔══╝     ██║           ██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝          ██║     ██║   ██║██║   ██║
╚██████╔╝███████╗   ██║           ██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗        ███████╗╚██████╔╝╚██████╔╝
 ╚═════╝ ╚══════╝   ╚═╝           ╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝        ╚══════╝ ╚═════╝  ╚═════╝
*/

//GetDeviceLog is
func (d Device) GetDeviceLog(setDeviceCode string, vasualFlag bool) {
	setQueryAdress := getDeviceLogLink(setDeviceCode)
	queryVariable.PostQuery(setQueryAdress, "", contentTypeJSON(), vasualFlag)
}
