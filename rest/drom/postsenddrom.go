package drom

import (
	"fmt"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

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
func (d Drom) SendDrom(setDeviceID string, visualFlag bool) string {

	setAddress := sendDromLink(setDeviceID)

	var queryVariable rest.Query
	query, errQuery := queryVariable.PostQuery(setAddress, "", contentTypeJSON(), visualFlag)
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
