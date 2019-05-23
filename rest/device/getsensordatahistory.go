package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

func (d Device) GetSensorDataHistory(setDeviceID string, setNodeName string, setThingName string, lastDataSize int, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := getSensorDataHistoryLink(setDeviceID, setNodeName, setThingName, lastDataSize)
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
