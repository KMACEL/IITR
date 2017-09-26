package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
██╗███╗   ██╗███████╗ ██████╗ ██████╗ ███╗   ███╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗███████╗
██║████╗  ██║██╔════╝██╔═══██╗██╔══██╗████╗ ████║██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║██╔════╝
██║██╔██╗ ██║█████╗  ██║   ██║██████╔╝██╔████╔██║███████║   ██║   ██║██║   ██║██╔██╗ ██║███████╗
██║██║╚██╗██║██╔══╝  ██║   ██║██╔══██╗██║╚██╔╝██║██╔══██║   ██║   ██║██║   ██║██║╚██╗██║╚════██║
██║██║ ╚████║██║     ╚██████╔╝██║  ██║██║ ╚═╝ ██║██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║███████║
╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝
*/

//Informations is
func (d Device) Informations(setDeviceCode string, setUnMarshal bool, vasualFlag bool) []byte {
	setQueryAdress := informationsLink(setDeviceCode)
	query, _ := queryVariable.GetQuery(setQueryAdress, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &deviceInformationJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}
