package reports

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/timop"
	"github.com/KMACEL/IITR/writefile"
)

//PresenceHistory is
type PresenceHistory struct{}

//Start is
func (p PresenceHistory) Start(devicesID ...string) {
	var (
		devices             device.Device
		presenceHistoryJSON device.PresenceHistroyJSON
		f                   *os.File
		presenceFileName    = "PresenceHistory_" + timop.GetTimeNamesFormat() + ".xlsx"
	)

	writefile.CreateFile(presenceFileName)
	f = writefile.OpenFile(f, presenceFileName)

	writefile.WriteText(f, "Device ID", "Online", "Offline")

	for _, deviceID := range devicesID {

		query := devices.PresenceHistory(deviceID, false, false)
		json.Unmarshal(query, &presenceHistoryJSON)

		for _, presenceH := range presenceHistoryJSON.List {
			fmt.Println(presenceH.Data.State, " - ", time.Unix(0, presenceH.CreateDate*1000000).String())
			if presenceH.Data.State == "ONLINE" {
				writefile.WriteText(f, deviceID, time.Unix(0, presenceH.CreateDate*1000000).String())
			} else if presenceH.Data.State == "OFFLINE" {
				writefile.WriteText(f, deviceID, "", time.Unix(0, presenceH.CreateDate*1000000).String())
			} else {
				writefile.WriteText(f, deviceID, "N/A", "N/A")
			}
		}
	}
}
