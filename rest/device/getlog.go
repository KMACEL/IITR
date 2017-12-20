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
func (d Device) GetLogList(setDeviceCode string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := getLogListLink(setDeviceCode)
	query, _ := q.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &logListJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
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
func (d Device) GetDeviceLog(setDeviceCode string, visualFlag bool) {
	setQueryAddress := getDeviceLogLink(setDeviceCode)
	q.PostQuery(setQueryAddress, "", contentTypeJSON(), visualFlag)
}
