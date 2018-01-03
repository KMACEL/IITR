package operations

import (
	"encoding/json"

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
}

//Start is
func (s StartApp) Start(devicesID ...string) {

	if devicesID != nil {
		for _, deviceID := range devicesID {
			s.devices.AppSS(device.StartApp, s.devices.DeviceID2Code(deviceID), s.StartPackageName, rest.Visible)
		}
	} else {
		query := s.devices.LocationMap(rest.NOMarshal, rest.Invisible)
		var locationJSONVariable device.LocationJSON
		json.Unmarshal(query, &locationJSONVariable)

		for _, deviceCoding := range locationJSONVariable.Extras {
			s.devices.AppSS(device.StartApp, deviceCoding.DeviceCode, s.StartPackageName, rest.Visible)
		}
	}

}
