package drom

import (
	"github.com/KMACEL/IITR/rest"
)

// Drom is
type Drom struct{}

//SendDrom is
func (d Drom) SendDrom(vasualFlag bool, setDeviceID ...string) string {
	for _, deviceID := range setDeviceID {
		setAdres := sendDromLink(deviceID)
		header := make(map[string]string)
		header["content-type"] = "application/json"

		query, _ := queryVariable.PostQuery(setAdres, "", header, vasualFlag)

		if query != nil {
			if string(query) != rest.ResponseNotFound {
				//json.Unmarshal(query, &responseMessageCodeJSONVariable)
				return "OK" //responseMessageCodeJSONVariable.Response
			}
			return rest.ResponseNotFound

		}
		return rest.ResponseNil
	}
	return "OK"
}
