package operations

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

/*
███████╗████████╗ █████╗ ██████╗ ████████╗         █████╗ ██████╗ ██████╗
██╔════╝╚══██╔══╝██╔══██╗██╔══██╗╚══██╔══╝        ██╔══██╗██╔══██╗██╔══██╗
███████╗   ██║   ███████║██████╔╝   ██║           ███████║██████╔╝██████╔╝
╚════██║   ██║   ██╔══██║██╔══██╗   ██║           ██╔══██║██╔═══╝ ██╔═══╝
███████║   ██║   ██║  ██║██║  ██║   ██║           ██║  ██║██║     ██║
╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝           ╚═╝  ╚═╝╚═╝     ╚═╝
*/
// For use Example:
//     var startApp util.StartApp
//     startApp.StartPackageName="com.estoty.game2048"
//     startApp.Start("867377020740787")

//StartApp is
type StartApp struct {
	devices          device.Device
	StartPackageName string
	DelayTime        time.Duration
}

//Start is
func (s StartApp) Start(devicesID ...string) {

	if devicesID != nil {
		for i, deviceID := range devicesID {
			s.devices.AppSS(device.StartApp, s.devices.DeviceID2Code(deviceID), s.StartPackageName, rest.Visible)
			fmt.Println("Device : ", deviceID, " | ", i+1, " - ", len(devicesID))
			if s.DelayTime != 0 {
				time.Sleep(s.DelayTime * time.Second)
			}
		}
	} else {
		query := s.devices.LocationMap(rest.NOMarshal, rest.Invisible)
		var locationJSONVariable device.LocationJSON
		json.Unmarshal(query, &locationJSONVariable)

		for _, deviceCoding := range locationJSONVariable.Extras {
			s.devices.AppSS(device.StartApp, deviceCoding.DeviceCode, s.StartPackageName, rest.Visible)
		}
	}
	fmt.Println("Success")
}
