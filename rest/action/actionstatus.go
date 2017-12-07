package action

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
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
func (a Action) GetActionStatus(setDeviceCode string, setControlType string, setSize int, visualFlag bool) []byte {
	var q rest.Query
	setQueryAddress := GetActionStatusLink(setDeviceCode, setControlType, setSize)
	query, actionError := q.GetQuery(setQueryAddress, visualFlag)

	errc.ErrorCenter(getActionStatusTag, actionError)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			unMarshalError := json.Unmarshal(query, &messageJSONVariable)
			errc.ErrorCenter(getActionStatusUnMarshalTag, unMarshalError)
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}
