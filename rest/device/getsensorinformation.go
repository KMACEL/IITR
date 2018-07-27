package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗        ███████╗███████╗███╗   ██╗███████╗ ██████╗ ██████╗         ██████╗  █████╗ ████████╗ █████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔════╝██╔════╝████╗  ██║██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗
██║  ███╗█████╗     ██║           ███████╗█████╗  ██╔██╗ ██║███████╗██║   ██║██████╔╝        ██║  ██║███████║   ██║   ███████║
██║   ██║██╔══╝     ██║           ╚════██║██╔══╝  ██║╚██╗██║╚════██║██║   ██║██╔══██╗        ██║  ██║██╔══██║   ██║   ██╔══██║
╚██████╔╝███████╗   ██║           ███████║███████╗██║ ╚████║███████║╚██████╔╝██║  ██║        ██████╔╝██║  ██║   ██║   ██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚══════╝╚══════╝╚═╝  ╚═══╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝
*/

// GetSensorData is
func (d Device) GetSensorData(setDeviceID string, setNodeName string, setThingName string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := getSensorDataLink(setDeviceID, setNodeName, setThingName)
	query, _ := q.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &sensorDataJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}
