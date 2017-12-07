package reports

import (
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest"
	"encoding/json"
	"fmt"
)

type PresenceStatus struct{
}

func (p PresenceStatus) Start (devicesID ...string){
	var devices device.Device

	for _,deviceId:=range devicesID{
		var presenceJSON device.PresenceInfoJSON
		query:=devices.PresenceInfo(deviceId,rest.NOMarshal,rest.Invisible)
		json.Unmarshal(query, &presenceJSON)
		fmt.Println(presenceJSON.DeviceID,presenceJSON.Data.State)
	}

}
