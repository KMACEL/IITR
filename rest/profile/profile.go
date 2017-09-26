package profile

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

//PushMode is
func (p Profile) PushMode(workingset string, setMode string, setPolicy string) string {

	setAdres := pushProfileLink(setMode, workingset)
	setBody := "{\"defaultPolicy\":{\"code\": \"" + setPolicy + "\"}}"

	header := make(map[string]string)
	header["content-type"] = "application/json"

	query, _ := queryVariable.PostQuery(setAdres, setBody, header, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &modeResponseJSONVariable)
			return modeResponseJSONVariable[0].Status
		}
		return rest.ResponseNotFound

	}
	return rest.ResponseNil
}
