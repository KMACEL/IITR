package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
)

// RemoveApplication is
func (d Device) RemoveApplication(setDeviceID string, setPackageName string, visualFlag bool) string {
	setAddress := removeApplicationLink(d.DeviceID2Code(setDeviceID))
	setBody := removeApplicationBody(setPackageName)

	query, removeError := q.PutQuery(setAddress, setBody, contentTypeJSON(), visualFlag)
	errc.ErrorCenter(removeApplicationErrorTag, removeError)

	if query != nil {
		json.Unmarshal(query, &responseMessageCodeJSONVariable)
		return responseMessageCodeJSONVariable.Response
	}
	return ""
}
