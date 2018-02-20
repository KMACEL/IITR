package operations

import (
	"fmt"
	"time"

	"github.com/KMACEL/IITR/rest/device"
)

/*
██████╗ ███████╗███████╗██████╗ ███████╗███████╗██╗  ██╗         ██████╗  █████╗ ████████╗███████╗██╗    ██╗ █████╗ ██╗   ██╗
██╔══██╗██╔════╝██╔════╝██╔══██╗██╔════╝██╔════╝██║  ██║        ██╔════╝ ██╔══██╗╚══██╔══╝██╔════╝██║    ██║██╔══██╗╚██╗ ██╔╝
██████╔╝█████╗  █████╗  ██████╔╝█████╗  ███████╗███████║        ██║  ███╗███████║   ██║   █████╗  ██║ █╗ ██║███████║ ╚████╔╝
██╔══██╗██╔══╝  ██╔══╝  ██╔══██╗██╔══╝  ╚════██║██╔══██║        ██║   ██║██╔══██║   ██║   ██╔══╝  ██║███╗██║██╔══██║  ╚██╔╝
██║  ██║███████╗██║     ██║  ██║███████╗███████║██║  ██║        ╚██████╔╝██║  ██║   ██║   ███████╗╚███╔███╔╝██║  ██║   ██║
╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝         ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝ ╚══╝╚══╝ ╚═╝  ╚═╝   ╚═╝
*/

//RefreshGateway is
type RefreshGateway struct {
	DelayTime time.Duration
}

//Start is
func (o RefreshGateway) Start(devicesID ...string) {
	var (
		devices device.Device
	)
	for _, deviceID := range devicesID {
		fmt.Println("Device : ", deviceID)
		devices.RefreshGatewayInfo(devices.DeviceID2Code(deviceID))
		time.Sleep(o.DelayTime * time.Second)
	}
}
