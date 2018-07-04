package adminarea

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
███╗   ███╗ ██████╗ ██╗   ██╗███████╗         █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗         █████╗ ██████╗ ███████╗ █████╗
████╗ ████║██╔═══██╗██║   ██║██╔════╝        ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║        ██╔══██╗██╔══██╗██╔════╝██╔══██╗
██╔████╔██║██║   ██║██║   ██║█████╗          ███████║██║  ██║██╔████╔██║██║██╔██╗ ██║        ███████║██████╔╝█████╗  ███████║
██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██╔══╝          ██╔══██║██║  ██║██║╚██╔╝██║██║██║╚██╗██║        ██╔══██║██╔══██╗██╔══╝  ██╔══██║
██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ███████╗        ██║  ██║██████╔╝██║ ╚═╝ ██║██║██║ ╚████║        ██║  ██║██║  ██║███████╗██║  ██║
╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚══════╝        ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝        ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
*/

// MoveAdminArea is
func (a AdminArea) MoveAdminArea(adr QueryRequirements) string {
	setAddress := moveAdminAreaLink()

	var adminAreaBodyJSON QueryBodyJSON

	for _, addDevice := range adr.AddToAdminAreaDeviceCode {
		adminAreaBodyJSON.Devices = append(adminAreaBodyJSON.Devices, CodeJSON{Code: addDevice})
	}
	adminAreaBodyJSON.Code = AdminArea{}.GetAdminArea(adr.AdminAreaName).Code

	jsonConvert, _ := json.Marshal(adminAreaBodyJSON)
	setBody := string(jsonConvert)

	query, errMoveAdminArea := rest.Query{}.PutQuery(setAddress, setBody, contentTypeJSON(), rest.Invisible)
	errc.ErrorCenter("MoveAdminArea-PutQuery", errMoveAdminArea)

	if query != nil {
		return string(query)
	}
	return ""
}
