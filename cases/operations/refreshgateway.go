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

/*
For use Example:
	var ref operations.RefreshGateway
	ref.RefreshParam = device.OSProfile
	ref.Start("123456", "123746")
*/

//RefreshGateway is
type RefreshGateway struct {
	DelayTime    time.Duration
	RefreshParam string
}

//Start is
func (o RefreshGateway) Start(devicesID ...string) {
	var (
		devices device.Device
	)
	for i, deviceID := range devicesID {
		fmt.Println("Device : ", deviceID)
		devices.RefreshGatewayInfo(devices.DeviceID2Code(deviceID), o.RefreshParam)
		time.Sleep(o.DelayTime * time.Second)
		if i%10 == 0 {
			fmt.Println(i+1, "/", len(devicesID))
		}
	}
	fmt.Println("Refresh Succes")
}
