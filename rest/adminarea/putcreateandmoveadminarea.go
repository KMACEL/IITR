package adminarea

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗         █████╗ ███╗   ██╗██████╗         ███╗   ███╗ ██████╗ ██╗   ██╗███████╗         █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗         █████╗ ██████╗ ███████╗ █████╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██╔══██╗████╗  ██║██╔══██╗        ████╗ ████║██╔═══██╗██║   ██║██╔════╝        ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║        ██╔══██╗██╔══██╗██╔════╝██╔══██╗
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ███████║██╔██╗ ██║██║  ██║        ██╔████╔██║██║   ██║██║   ██║█████╗          ███████║██║  ██║██╔████╔██║██║██╔██╗ ██║        ███████║██████╔╝█████╗  ███████║
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██╔══██║██║╚██╗██║██║  ██║        ██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██╔══╝          ██╔══██║██║  ██║██║╚██╔╝██║██║██║╚██╗██║        ██╔══██║██╔══██╗██╔══╝  ██╔══██║
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ██║  ██║██║ ╚████║██████╔╝        ██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ███████╗        ██║  ██║██████╔╝██║ ╚═╝ ██║██║██║ ╚████║        ██║  ██║██║  ██║███████╗██║  ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝         ╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚══════╝        ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝        ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
*/

/*
	var adminArea adminarea.AdminArea
	var adminAreaReq adminarea.QueryRequirements

	adminAreaReq.AdminAreaName = "{Created_Admin_Area_Name}"
	adminAreaReq.AddToAdminAreaDeviceCode = []string{"{Device_ID}"}

	adminArea.CreateAndMoveAdminArea(adminAreaReq)
*/

// CreateAndMoveAdminArea is creates a new admin area and the given devices perform the assigned admin area.
func (a AdminArea) CreateAndMoveAdminArea(adr QueryRequirements) string {
	setAddress := moveAdminAreaLink()

	var adminAreaBodyJSON QueryBodyJSON

	for _, addDevice := range adr.AddToAdminAreaDeviceCode {
		adminAreaBodyJSON.Devices = append(adminAreaBodyJSON.Devices, CodeJSON{Code: addDevice})
	}
	adminAreaBodyJSON.Name = adr.AdminAreaName

	jsonConvert, _ := json.Marshal(adminAreaBodyJSON)
	setBody := string(jsonConvert)

	query, errCreateAndMoveAdminArea := rest.Query{}.PutQuery(setAddress, setBody, contentTypeJSON(), rest.Invisible)
	errc.ErrorCenter("errCreateAndMoveAdminArea-PutQuery", errCreateAndMoveAdminArea)

	if query != nil {
		return string(query)
	}
	return ""
}
