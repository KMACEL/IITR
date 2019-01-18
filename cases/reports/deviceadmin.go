package reports

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/timop"
	"github.com/KMACEL/IITR/writefile"
)

/*
██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗         █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗
██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝        ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║
██║  ██║█████╗  ██║   ██║██║██║     █████╗          ███████║██║  ██║██╔████╔██║██║██╔██╗ ██║
██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝          ██╔══██║██║  ██║██║╚██╔╝██║██║██║╚██╗██║
██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗        ██║  ██║██████╔╝██║ ╚═╝ ██║██║██║ ╚████║
╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝        ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝
*/

/*
For use Example:
	reports.DeviceAdmin{}.Start()
*/

// DeviceAdmin is
type DeviceAdmin struct{}

// Start is
func (d DeviceAdmin) Start() {

	var f *os.File
	fileName := "DeviceAdmin_" + timop.GetTimeNamesFormat() + ".csv"
	writefile.CreateFile(fileName)
	f = writefile.OpenFile(f, fileName)
	writefile.SplitCharacter = ";"

	query := device.Device{}.LocationMap(rest.NOMarshal, rest.Invisible)
	deviceCode := device.LocationJSON{}
	json.Unmarshal(query, &deviceCode)

	var a int
	for i, deviceID := range deviceCode.Extras {
		var dev device.ModiverseInfoJSON
		json.Unmarshal(device.Device{}.ModiverseInfo(deviceID.DeviceID, false, false), &dev)

		if dev.Data.DeviceAdmin == false && dev.DeviceID != "" {
			a++
			fmt.Println(a, " - ", deviceID.DeviceID, " - ", dev.Data.DeviceAdmin, " : ", dev.DeviceID)
			writefile.WriteText(f, deviceID.DeviceID, strconv.FormatBool(dev.Data.DeviceAdmin))
		}

		if i%100 == 0 {
			fmt.Println(i, " / ", len(deviceCode.Extras))

		}
	}
}
