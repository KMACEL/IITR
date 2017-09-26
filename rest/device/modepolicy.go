package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 █████╗  ██████╗████████╗██╗██╗   ██╗███████╗        ██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗        ██████╗  ██████╗ ██╗     ██╗ ██████╗██╗   ██╗
██╔══██╗██╔════╝╚══██╔══╝██║██║   ██║██╔════╝        ██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝        ██╔══██╗██╔═══██╗██║     ██║██╔════╝╚██╗ ██╔╝
███████║██║        ██║   ██║██║   ██║█████╗          ██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗          ██████╔╝██║   ██║██║     ██║██║      ╚████╔╝
██╔══██║██║        ██║   ██║╚██╗ ██╔╝██╔══╝          ██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝          ██╔═══╝ ██║   ██║██║     ██║██║       ╚██╔╝
██║  ██║╚██████╗   ██║   ██║ ╚████╔╝ ███████╗        ██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗        ██║     ╚██████╔╝███████╗██║╚██████╗   ██║
╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝  ╚═══╝  ╚══════╝        ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝        ╚═╝      ╚═════╝ ╚══════╝╚═╝ ╚═════╝   ╚═╝
*/

//ActiveProfilePolicy is
func (d Device) ActiveProfilePolicy(deviceID string, setUnMarshal bool, vasualFlag bool) []byte {
	var query []byte

	queryLink := modePolicyLink(deviceID)
	query, _ = rest.Query{}.GetQuery(queryLink, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &activeProfilePolicyJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}
