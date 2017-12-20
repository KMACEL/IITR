package reports

import (
	"encoding/json"
	"fmt"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

//PresenceStatus is
type PresenceStatus struct {
}

//Start is
func (p PresenceStatus) Start(devicesID ...string) {
	var devices device.Device

	for _, deviceID := range devicesID {
		var presenceJSON device.PresenceInfoJSON
		query := devices.PresenceInfo(deviceID, rest.NOMarshal, rest.Invisible)
		json.Unmarshal(query, &presenceJSON)
		fmt.Println(presenceJSON.DeviceID, presenceJSON.Data.State)
	}

}
