package drom

import (
	"github.com/KMACEL/IITR/rest"
)

// Drom is
type Drom struct{}

//SendDrom is
func (d Drom) SendDrom(visualFlag bool, setDeviceID string) string {

	setAddress := sendDromLink(setDeviceID)
	header := make(map[string]string)
	header["content-type"] = "application/json"

	query, _ := queryVariable.PostQuery(setAddress, "", header, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			//json.Unmarshal(query, &responseMessageCodeJSONVariable)
			return "OK" //responseMessageCodeJSONVariable.Response
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil

}
