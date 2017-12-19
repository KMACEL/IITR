package operations

import (
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

//ClearApp is
type ClearApp struct {
	devices          device.Device
	ClearPackageName string
}

//Start is
func (c ClearApp) Start(devicesID ...string) {
	for _, deviceID := range devicesID {
		c.devices.ClearAppData(c.devices.DeviceID2Code(deviceID), c.ClearPackageName, rest.Visible)
	}
}
