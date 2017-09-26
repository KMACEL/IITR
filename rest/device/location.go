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
