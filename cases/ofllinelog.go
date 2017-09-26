package cases

import (
	"encoding/json"
	"log"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

// OfflineLog is
type OfflineLog struct {
	devices  device.Device
	status   map[string]string
	DeviceID []string
}

// SetOfflineLogDeviceID is
func (o OfflineLog) SetOfflineLogDeviceID(deviceID ...string) {

}

// Start is
func (o OfflineLog) Start() {
	o.status = make(map[string]string)
	for _, devicesMap := range o.DeviceID {
		o.status[devicesMap] = ""
	}
	for {
		for _, devices := range o.DeviceID {
			query := o.devices.PresenceInfo(devices, rest.NOMarshal, rest.Invisible)
			var presenceJSON device.PresenceInfoJSON
			json.Unmarshal(query, &presenceJSON)
			presence := presenceJSON.Data.State
			if presence != o.status[devices] {
				log.Println(presenceJSON.DeviceID, ",", presenceJSON.Data.State, ",", presenceJSON.CreateDate)
				o.devices.GetDeviceLog(o.devices.DeviceID2Code(devices), rest.Invisible)
				o.status[devices] = presence
			}
		}
		time.Sleep(2 * time.Minute)
	}
}
