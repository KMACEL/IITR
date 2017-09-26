package cases

import (
	"log"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

// GetLogDevices is
type GetLogDevices struct {
	dev device.Device
}

//Start is
func (g GetLogDevices) Start(deviceID []string) {

	for _, getDevice := range deviceID {
		g.dev.GetDeviceLog(getDevice, rest.Visible)
		log.Println("Getlog Device ID : " + getDevice)

	}
}
