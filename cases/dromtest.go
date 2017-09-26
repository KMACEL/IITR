package cases

import (
	"fmt"
	"log"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/drom"
	"github.com/KMACEL/IITR/timop"
)

//DromTest is
type DromTest struct {
	devices device.Device
	dromes  drom.Drom
}

//Start is
func (d DromTest) Start(deviceID string) {
	var random int
	for {
		random = timop.Random(0, 2)
		fmt.Println(random)
		if random == 0 {
			log.Println("Send Drom...")
			d.sendDrom(deviceID)
			time.Sleep(time.Duration(timop.Random(0, 7)) * time.Minute)
		} else if random == 1 {
			log.Println("Send Reboot...")
			d.reboot(deviceID)
			time.Sleep(time.Duration(timop.Random(0, 7)) * time.Minute)
		}
	}
}

func (d DromTest) sendDrom(deviceID string) {
	d.dromes.SendDrom(rest.Invisible, deviceID)
}

func (d DromTest) reboot(deviceID string) {
	d.devices.Reboot(d.devices.DeviceID2Code(deviceID), rest.Visible)
}
