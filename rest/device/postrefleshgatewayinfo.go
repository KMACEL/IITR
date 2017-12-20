package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

//POST : https://api.ardich.com/api/v3/device/cea9bbd434b04a7db1865d210f449f0e/control/status

// RefreshGatewayInfo is
func (d Device) RefreshGatewayInfo(deviceCode string) string {
	//setBody := applicationOperationsBodyLink(setApplicationPackage)
	setAddress := "https://api.ardich.com/api/v3/device/" + deviceCode + "/control/status"
	query, _ := q.PostQuery(setAddress, "setBody", contentTypeJSON(), true)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responseMessageCodeJSONVariable)
			return responseMessageCodeJSONVariable.Response
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}
