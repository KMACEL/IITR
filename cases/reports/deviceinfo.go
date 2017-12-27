package reports

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"os"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/timop"
	"github.com/KMACEL/IITR/writefile"
)

//Delete Packet

//DeviceInformation is
type DeviceInformation struct{}

//TODO: sadece firmware değil bütün deviceindo bilgilerini bas

//Start is
func (d DeviceInformation) Start(deviceID ...string) {
	var (
		counter   int
		devices   device.Device
		file      *os.File
		fileError *os.File
		fileAll   *os.File
	)
	writefile.CreateFile("UpdateYes_" + timop.GetTimeNamesFormat() + ".xlsx")
	file = writefile.OpenFile(file, "UpdateYes_"+timop.GetTimeNamesFormat()+".xlsx")

	writefile.CreateFile("UpdateNotYet_" + timop.GetTimeNamesFormat() + ".xlsx")
	fileError = writefile.OpenFile(fileError, "UpdateNotYet_"+timop.GetTimeNamesFormat()+".xlsx")

	writefile.CreateFile("UpdateAll_" + timop.GetTimeNamesFormat() + ".xlsx")
	fileAll = writefile.OpenFile(fileAll, "UpdateAll_"+timop.GetTimeNamesFormat()+".xlsx")

	writefile.WriteText(file, "Device ID", "Firmware", "Modiverse Version", "State")
	writefile.WriteText(fileError, "Device ID", "Firmware", "Modiverse Version", "State")
	writefile.WriteText(fileAll, "Device ID", "Firmware", "Modiverse Version", "State")

	query := device.Device{}.LocationMap(rest.NOMarshal, rest.Invisible)
	if query != nil {
		if string(query) != rest.ResponseNotFound {

			// This assignment is aimed at resetting the variable
			deviceCode := device.LocationJSON{}
			json.Unmarshal(query, &deviceCode)

			for _, deviceCodings := range deviceID {
				deviceCoding := deviceCodings
				queryInformation := devices.DeviceInformation(devices.DeviceID2Code(deviceCoding), rest.NOMarshal, rest.Invisible)

				if queryInformation != nil {
					if string(queryInformation) != rest.ResponseNotFound {

						deviceInformation := device.InformationJSON{}
						json.Unmarshal(queryInformation, &deviceInformation)
						writefile.WriteText(fileAll, deviceInformation.DeviceID, deviceInformation.OsProfile.Display, deviceInformation.ModeAppVersion, deviceInformation.Presence.State, time.Unix(0, deviceInformation.DetailLastModifiedDate*1000000).String(), "Yes")

						if deviceInformation.OsProfile.Display == "rkpx2-eng 4.4.4 KTU84Q eng.turkey.20171027.143708 test-keys" {
							fmt.Println(deviceInformation.DeviceID, ",", deviceInformation.OsProfile.Display, ",", deviceInformation.Presence.State, ",", time.Unix(0, deviceInformation.DetailLastModifiedDate*1000000).String(), ",")
							writefile.WriteText(file, deviceInformation.DeviceID, deviceInformation.OsProfile.Display, deviceInformation.ModeAppVersion, deviceInformation.Presence.State, time.Unix(0, deviceInformation.DetailLastModifiedDate*1000000).String(), "Yes")
							//writefile.WriteText(fileAll, deviceInformation.DeviceID, deviceInformation.OsProfile.Display, deviceInformation.ModeAppVersion, deviceInformation.Presence.State, time.Unix(0, deviceInformation.DetailLastModifiedDate*1000000).String(), "Yes")

						} else {
							writefile.WriteText(fileError, deviceInformation.DeviceID, deviceInformation.OsProfile.Display, deviceInformation.ModeAppVersion, deviceInformation.Presence.State, time.Unix(0, deviceInformation.DetailLastModifiedDate*1000000).String(), "No")
							//writefile.WriteText(fileAll, deviceInformation.DeviceID, deviceInformation.OsProfile.Display, deviceInformation.ModeAppVersion, deviceInformation.Presence.State, time.Unix(0, deviceInformation.DetailLastModifiedDate*1000000).String(), "No")

						}
						counter++
					}
				}
				if counter%10 == 0 {
					fmt.Println(counter, " / ", len(deviceID))
				}
			}
		}
	}
	log.Println("Total : ", counter)
}
