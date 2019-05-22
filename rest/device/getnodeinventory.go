package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗        ███╗   ██╗ ██████╗ ██████╗ ███████╗        ██╗███╗   ██╗██╗   ██╗███████╗███╗   ██╗████████╗ ██████╗ ██████╗ ██╗   ██╗
██╔════╝ ██╔════╝╚══██╔══╝        ████╗  ██║██╔═══██╗██╔══██╗██╔════╝        ██║████╗  ██║██║   ██║██╔════╝████╗  ██║╚══██╔══╝██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║  ███╗█████╗     ██║           ██╔██╗ ██║██║   ██║██║  ██║█████╗          ██║██╔██╗ ██║██║   ██║█████╗  ██╔██╗ ██║   ██║   ██║   ██║██████╔╝ ╚████╔╝
██║   ██║██╔══╝     ██║           ██║╚██╗██║██║   ██║██║  ██║██╔══╝          ██║██║╚██╗██║╚██╗ ██╔╝██╔══╝  ██║╚██╗██║   ██║   ██║   ██║██╔══██╗  ╚██╔╝
╚██████╔╝███████╗   ██║           ██║ ╚████║╚██████╔╝██████╔╝███████╗        ██║██║ ╚████║ ╚████╔╝ ███████╗██║ ╚████║   ██║   ╚██████╔╝██║  ██║   ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═══╝ ╚═════╝ ╚═════╝ ╚══════╝        ╚═╝╚═╝  ╚═══╝  ╚═══╝  ╚══════╝╚═╝  ╚═══╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝   ╚═╝
*/

// GetNodeInventory is
func (d Device) GetNodeInventory(setDeviceID string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := getNodeInventoryLink(setDeviceID)
	query, _ := q.GetQuery(setQueryAddress, visualFlag)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &nodeInventoryJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}
