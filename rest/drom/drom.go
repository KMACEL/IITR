package drom

import (
	"fmt"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
██████╗ ██████╗  ██████╗ ███╗   ███╗
██╔══██╗██╔══██╗██╔═══██╗████╗ ████║
██║  ██║██████╔╝██║   ██║██╔████╔██║
██║  ██║██╔══██╗██║   ██║██║╚██╔╝██║
██████╔╝██║  ██║╚██████╔╝██║ ╚═╝ ██║
╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝
*/

// Drom is a method for remotely licensing a device.
// The important thing to note here is that the device
// must have a drom-recording. Otherwise, no licensing will be done.
type Drom struct{}

/*
███████╗███████╗███╗   ██╗██████╗         ██████╗ ██████╗  ██████╗ ███╗   ███╗
██╔════╝██╔════╝████╗  ██║██╔══██╗        ██╔══██╗██╔══██╗██╔═══██╗████╗ ████║
███████╗█████╗  ██╔██╗ ██║██║  ██║        ██║  ██║██████╔╝██║   ██║██╔████╔██║
╚════██║██╔══╝  ██║╚██╗██║██║  ██║        ██║  ██║██╔══██╗██║   ██║██║╚██╔╝██║
███████║███████╗██║ ╚████║██████╔╝        ██████╔╝██║  ██║╚██████╔╝██║ ╚═╝ ██║
╚══════╝╚══════╝╚═╝  ╚═══╝╚═════╝         ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝
*/

/*
drom.Drom{}.SendDrom(true, "{YOUR_DEVICE_ID}")
*/

// SendDrom carries out the transmission of the drom configuration to
// which it was previously registered to a device with the Device ID.
func (d Drom) SendDrom(visualFlag bool, setDeviceID string) string {

	setAddress := sendDromLink(setDeviceID)
	header := make(map[string]string)
	header["content-type"] = "application/json"

	var queryVariable rest.Query
	query, errQuery := queryVariable.PostQuery(setAddress, "", header, visualFlag)
	errc.ErrorCenter("SendDrom-PostQuery", errQuery)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if visualFlag {
				fmt.Println("Drom SEND : OK")
			}
			return "OK"
		}
		if visualFlag {
			fmt.Println("Drom SEND : ResponseNotFound")
		}
		return rest.ResponseNotFound
	}
	if visualFlag {
		fmt.Println("Drom SEND : QUERY is NIL")
	}
	return rest.ResponseNil
}
