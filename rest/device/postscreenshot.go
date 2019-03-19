package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

// ScreenShotBodyJSON is
type ScreenShotBodyJSON struct {
	DestinationPath string `json:"destinationPath"`
	FileName        string `json:"fileName"`
	StorageType     string `json:"storageType"`
	Upload          bool   `json:"upload"`
}

//https://api.ardich.com/api/v3/device/928f500d47e14545b6b235d1bff7083a/control/take-screenshot
func screenShotLink(deviceCode string) string {
	return device + deviceCode + control + takescreenshot
}

/*{
  "destinationPath": "test",
  "fileName": "test1",
  "storageType": "internal",
  "upload": true
}*/

// ScreenShot is
func (d Device) ScreenShot(deviceCode string, screenShotBody ScreenShotBodyJSON) string {
	setAddress := screenShotLink(deviceCode)
	jsonConvert, _ := json.Marshal(screenShotBody)

	query, _ := q.PostQuery(setAddress, string(jsonConvert), contentTypeJSON(), true)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responseMessageCodeJSONVariable)
			return responseMessageCodeJSONVariable.Response
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil
}
