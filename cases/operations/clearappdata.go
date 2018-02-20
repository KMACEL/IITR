package operations

import (
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

/*
 ██████╗██╗     ███████╗ █████╗ ██████╗          █████╗ ██████╗ ██████╗
██╔════╝██║     ██╔════╝██╔══██╗██╔══██╗        ██╔══██╗██╔══██╗██╔══██╗
██║     ██║     █████╗  ███████║██████╔╝        ███████║██████╔╝██████╔╝
██║     ██║     ██╔══╝  ██╔══██║██╔══██╗        ██╔══██║██╔═══╝ ██╔═══╝
╚██████╗███████╗███████╗██║  ██║██║  ██║        ██║  ██║██║     ██║
 ╚═════╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝     ╚═╝
*/

// ClearApp is trying to delete the data of the desired application, found in the specified "deviceID" s.
type ClearApp struct {
	devices          device.Device
	ClearPackageName string
}

// Start is the function of ClearApp. If you serialize from outside, you get "deviceID"
func (c ClearApp) Start(devicesID ...string) {
	for _, deviceID := range devicesID {
		c.devices.ClearAppData(c.devices.DeviceID2Code(deviceID), c.ClearPackageName, rest.Visible)
	}
}
