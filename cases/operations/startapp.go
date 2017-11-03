package operations

import (
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest"
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


type StartApp struct{
	devices device.Device
	StartPackageName string
}

func (s StartApp) Start(devicesID ...string){
	for _,deviceID :=range devicesID {
		s.devices.AppSS(device.StartApp,s.devices.DeviceID2Code(deviceID),s.StartPackageName,rest.Visible)
	}
}