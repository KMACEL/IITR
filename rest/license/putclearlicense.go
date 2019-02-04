package license

import (
	"fmt"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

// ClearLicense is
func (l License) ClearLicense(deviceID string, visualFlag bool) string {
	setAddress := clearLicenseLink(deviceID)

	var queryVariable rest.Query
	query, errQuery := queryVariable.PutQuery(setAddress, "", contentTypeJSON(), visualFlag)
	errc.ErrorCenter("ClearLicense-PutQuery", errQuery)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if visualFlag {
				fmt.Println("ClearLicense : OK")
			}
			return "OK"
		}
		if visualFlag {
			fmt.Println("ClearLicense : ResponseNotFound")
		}
		return rest.ResponseNotFound
	}
	if visualFlag {
		fmt.Println("ClearLicense : QUERY is NIL")
	}
	return rest.ResponseNil
}
