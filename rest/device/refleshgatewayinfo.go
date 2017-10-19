package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

//POST : https://api.ardich.com/api/v3/device/cea9bbd434b04a7db1865d210f449f0e/control/status

// RefleshGatewayInfo is
func (d Device) RefleshGatewayInfo(deviceCode string) string {
	//setBody := applicationOperationsBodyLink(setApplicationPackage)
	setAdres := "https://api.ardich.com/api/v3/device/" + deviceCode + "/control/status"
	query, _ := queryVariable.PostQuery(setAdres, "setBody", contentTypeJSON(), true)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responseMessageCodeJSONVariable)
			return responseMessageCodeJSONVariable.Response
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}
