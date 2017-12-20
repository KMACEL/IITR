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
func (d Device) ActiveProfilePolicy(deviceID string, setUnMarshal bool, visualFlag bool) []byte {
	queryLink := modePolicyLink(deviceID)
	query, _ := q.GetQuery(queryLink, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &activeProfilePolicyJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}
