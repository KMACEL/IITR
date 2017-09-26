package cases

import (
	"encoding/json"
	"strconv"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

// BlockedAppList is
type BlockedAppList struct {
	writeCsvArray []string
	devices       device.Device
}

// Start is
func (b BlockedAppList) Start() {
	query := b.devices.LocationMap(rest.NOMarshal, rest.Invisible)
	if query != nil {
		if string(query) != rest.ResponseNotFound {
			deviceCode := device.LocationJSON{}
			json.Unmarshal(query, &deviceCode)

			for _, deviceCoding := range deviceCode.Extras {

				applicationQuery := b.devices.ApplicationInfo(deviceCoding.DeviceID, rest.NOMarshal, rest.Invisible)

				applications := device.ApplicationInfoJSON{}
				json.Unmarshal(applicationQuery, &applications)

				for _, apllicationStatus := range applications.Data {
					b.writeCsvArray = append(b.writeCsvArray, strconv.FormatBool(apllicationStatus.Running), strconv.Itoa(apllicationStatus.Blocked), "\n")
				}
			}
		}
	}
}
