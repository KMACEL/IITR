package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
)

// SetLabel is
func (d Device) SetLabel(setDeviceID string, labelName string, visualFlag bool) string {
	setAddress := "https://api.ardich.com/api/v3/device/" + setDeviceID + "/label"
	var setBodyJSON SetLabelBodyJSON
	setBodyJSON.Label = labelName
	setBody, _ := json.Marshal(setBodyJSON)

	query, removeError := q.PutQuery(setAddress, string(setBody), contentTypeJSON(), visualFlag)
	errc.ErrorCenter(removeApplicationErrorTag, removeError)

	if query != nil {
		json.Unmarshal(query, &responseMessageCodeJSONVariable)
		return responseMessageCodeJSONVariable.Response
	}
	return ""
}

// SetLabelBodyJSON is
type SetLabelBodyJSON struct {
	Label string `json:"label"`
}
