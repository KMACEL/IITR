package device

import "github.com/KMACEL/IITR/errc"

/*
██████╗ ███████╗██████╗  ██████╗  ██████╗ ████████╗
██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝
██████╔╝█████╗  ██████╔╝██║   ██║██║   ██║   ██║
██╔══██╗██╔══╝  ██╔══██╗██║   ██║██║   ██║   ██║
██║  ██║███████╗██████╔╝╚██████╔╝╚██████╔╝   ██║
╚═╝  ╚═╝╚══════╝╚═════╝  ╚═════╝  ╚═════╝    ╚═╝
*/

//Reboot is
func (d Device) Reboot(setDeviceID string, visualFlag bool) {
	setQueryAddress := rebootLink(d.DeviceID2Code(setDeviceID))
	_, errReboot := q.PostQuery(setQueryAddress, "", contentTypeJSON(), visualFlag)
	errc.ErrorCenter("Reboot", errReboot)
}
