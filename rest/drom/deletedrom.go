package drom

import (
	"fmt"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
██████╗ ███████╗██╗     ███████╗████████╗███████╗        ██████╗ ██████╗  ██████╗ ███╗   ███╗
██╔══██╗██╔════╝██║     ██╔════╝╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗██╔═══██╗████╗ ████║
██║  ██║█████╗  ██║     █████╗     ██║   █████╗          ██║  ██║██████╔╝██║   ██║██╔████╔██║
██║  ██║██╔══╝  ██║     ██╔══╝     ██║   ██╔══╝          ██║  ██║██╔══██╗██║   ██║██║╚██╔╝██║
██████╔╝███████╗███████╗███████╗   ██║   ███████╗        ██████╔╝██║  ██║╚██████╔╝██║ ╚═╝ ██║
╚═════╝ ╚══════╝╚══════╝╚══════╝   ╚═╝   ╚══════╝        ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝
*/

//DeleteDrom is
func (d Drom) DeleteDrom(setDeviceID string, visualFlag bool) string {

	setAddress := deleteDromLink(setDeviceID)

	var queryVariable rest.Query
	query, errQuery := queryVariable.DeleteQuery(setAddress, "", contentTypeJSON(), visualFlag)
	errc.ErrorCenter("DeleteDrom-PostQuery", errQuery)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if visualFlag {
				fmt.Println("DeleteDrom : OK")
			}
			return "OK"
		}
		if visualFlag {
			fmt.Println("DeleteDrom : ResponseNotFound")
		}
		return rest.ResponseNotFound
	}
	if visualFlag {
		fmt.Println("DeleteDrom : QUERY is NIL")
	}
	return rest.ResponseNil
}
