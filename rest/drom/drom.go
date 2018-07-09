package drom

import (
	"fmt"

	"github.com/KMACEL/IITR/rest"
)

// Drom is
type Drom struct{}

//SendDrom is
func (d Drom) SendDrom(visualFlag bool, setDeviceID string) string {

	setAddress := sendDromLink(setDeviceID)
	header := make(map[string]string)
	header["content-type"] = "application/json"

	query, _ := queryVariable.PostQuery(setAddress, "", header, visualFlag)

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
