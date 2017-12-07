package tests

import (
	"log"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

// MapsTest is
type MapsTest struct{}

// Start is
func (m MapsTest) Start(deviceID string, setPackageName string) {
	var devices device.Device

retry:
	log.Println("Reboot : OK")
	devices.Reboot(deviceID, rest.Visible)

	log.Println("Time Sleep : 3 Minute")
	time.Sleep(3 * time.Minute)

	log.Println("Start Application : ", setPackageName)
	devices.AppSS(device.StartApp, devices.DeviceID2Code(deviceID), "com.google.android.apps.maps", rest.Visible)

	log.Println("Time Sleep : 2 Minute")
	time.Sleep(2 * time.Minute)

	log.Println("Remove Application : ", setPackageName)
	devices.RemoveApplication(deviceID, setPackageName, rest.Visible)

	log.Println("Time Sleep : 1 Minute")
	time.Sleep(1 * time.Minute)
	goto retry

}
