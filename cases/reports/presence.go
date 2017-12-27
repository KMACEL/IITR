package reports

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/timop"
	"github.com/KMACEL/IITR/writefile"
)

//PresenceStatus is
type PresenceStatus struct {
}

//Start is
func (p PresenceStatus) Start(devicesID ...string) {
	var (
		devices          device.Device
		file             *os.File
		presenceFileName = "Presence_" + timop.GetTimeNamesFormat() + ".xlsx"
	)

	writefile.CreateFile(presenceFileName)
	file = writefile.OpenFile(file, presenceFileName)

	for _, deviceID := range devicesID {
		var presenceJSON device.PresenceInfoJSON
		query := devices.PresenceInfo(deviceID, rest.NOMarshal, rest.Invisible)
		json.Unmarshal(query, &presenceJSON)
		fmt.Println(presenceJSON.DeviceID, presenceJSON.Data.State)
		writefile.WriteText(file, deviceID, presenceJSON.Data.State, time.Unix(0, presenceJSON.CreateDate*1000000).String())
	}

}
