package operations

import (
	"log"
	"time"

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
	DelayTime time.Duration
}

//Start is
func (g GetLogDevices) Start(getLogDeviceID ...string) {
	var (
		devices device.Device
	)
retry:
	for _, getDevice := range getLogDeviceID {
		devices.GetDeviceLog(devices.DeviceID2Code(getDevice), rest.Visible)
		log.Println("Getlog Device ID : " + getDevice)
	}
	time.Sleep(g.DelayTime * time.Minute)
	goto retry
}
