package action

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗         █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗        ███████╗████████╗ █████╗ ████████╗██╗   ██╗███████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝██║   ██║██╔════╝
██║  ███╗█████╗     ██║           ███████║██║        ██║   ██║██║   ██║██╔██╗ ██║        ███████╗   ██║   ███████║   ██║   ██║   ██║███████╗
██║   ██║██╔══╝     ██║           ██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║        ╚════██║   ██║   ██╔══██║   ██║   ██║   ██║╚════██║
╚██████╔╝███████╗   ██║           ██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║        ███████║   ██║   ██║  ██║   ██║   ╚██████╔╝███████║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚══════╝
*/

//GetActionStatus is
func (a Action) GetActionStatus(setDeviceCode string, setControlType string, setSize int, vasualFlag bool) []byte {
	setQueryAdress := GetActionStatusLink(setDeviceCode, setControlType, setSize)

	query, _ := rest.Query{}.GetQuery(setQueryAdress, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &messageJSONVariable)
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}
