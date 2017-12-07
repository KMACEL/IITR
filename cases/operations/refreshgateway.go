package operations

import "github.com/KMACEL/IITR/rest/device"

type RefreshGateway struct{}

func (o RefreshGateway) Start (devicesID ...string) {
	var (
		devices device.Device
	)
	for _, deviceID := range devicesID {
		devices.RefreshGatewayInfo(devices.DeviceID2Code(deviceID))
	}
}