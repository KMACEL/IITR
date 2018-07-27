package device

import (
	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
██╗    ██╗██╗██████╗ ███████╗
██║    ██║██║██╔══██╗██╔════╝
██║ █╗ ██║██║██████╔╝█████╗
██║███╗██║██║██╔═══╝ ██╔══╝
╚███╔███╔╝██║██║     ███████╗
 ╚══╝╚══╝ ╚═╝╚═╝     ╚══════╝
*/
// 	device.Device{}.Wipe(device.Device{}.DeviceID2Code("{YOUR_DEVICE_ID}"))

//Wipe is
func (d Device) Wipe(deviceCode string) {
	setQueryAddress := wipeLink(deviceCode)
	_, errWipe := q.PostQuery(setQueryAddress, "", contentTypeJSON(), rest.Invisible)
	errc.ErrorCenter("Wipe", errWipe)
}
