package operations

import (
	"log"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

/*
 ██████╗ ███████╗████████╗            ██╗      ██████╗  ██████╗             ██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗███████╗
██╔════╝ ██╔════╝╚══██╔══╝            ██║     ██╔═══██╗██╔════╝             ██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝██╔════╝
██║  ███╗█████╗     ██║               ██║     ██║   ██║██║  ███╗            ██║  ██║█████╗  ██║   ██║██║██║     █████╗  ███████╗
██║   ██║██╔══╝     ██║               ██║     ██║   ██║██║   ██║            ██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝  ╚════██║
╚██████╔╝███████╗   ██║               ███████╗╚██████╔╝╚██████╔╝            ██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗███████║
 ╚═════╝ ╚══════╝   ╚═╝               ╚══════╝ ╚═════╝  ╚═════╝             ╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝╚══════╝
*/
// For use Example:
//	   var getLogOperations util.GetLogDevices
//     getLogOperations.Start("867377020740787")

// GetLogDevices is
type GetLogDevices struct {
}

//Start is
func (g GetLogDevices) Start(getLogDeviceID ...string) {
	var (
		devices device.Device
	)

	for _, getDevice := range getLogDeviceID {
		devices.GetDeviceLog(devices.DeviceID2Code(getDevice), rest.Visible)
		log.Println("Getlog Device ID : " + getDevice)
	}
}
