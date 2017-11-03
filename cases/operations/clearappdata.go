package operations

import (
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest"
)

type ClearApp struct{
	devices device.Device
	ClearPackageName string
}

func (c ClearApp) Start(devicesID ...string){
	for _,deviceID :=range devicesID {
		c.devices.ClearAppData(c.devices.DeviceID2Code(deviceID),c.ClearPackageName,rest.Visible)
	}
}