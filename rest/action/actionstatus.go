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
// action.Action{}.GetActionStatus(device.Device{}.DeviceID2Code("gopher@go"), action.PushDROM, 1000, rest.Invisible)

//GetActionStatus is
func (a Action) GetActionStatus(setDeviceCode string, setControlType string, setSize int, visualFlag bool) []byte {
	var q rest.Query
	setQueryAddress := getActionStatusLink(setDeviceCode, setControlType, setSize)
	query, actionError := q.GetQuery(setQueryAddress, visualFlag)
	errc.ErrorCenter(getActionStatusTag, actionError)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			var messageJSONVariable ResponseActionMessageJSON
			unMarshalError := json.Unmarshal(query, &messageJSONVariable)
			errc.ErrorCenter(getActionStatusUnMarshalTag, unMarshalError)
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}
