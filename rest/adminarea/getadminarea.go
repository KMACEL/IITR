package adminarea

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗         █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗         █████╗ ██████╗ ███████╗ █████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║        ██╔══██╗██╔══██╗██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ███████║██║  ██║██╔████╔██║██║██╔██╗ ██║        ███████║██████╔╝█████╗  ███████║
██║   ██║██╔══╝     ██║           ██╔══██║██║  ██║██║╚██╔╝██║██║██║╚██╗██║        ██╔══██║██╔══██╗██╔══╝  ██╔══██║
╚██████╔╝███████╗   ██║           ██║  ██║██████╔╝██║ ╚═╝ ██║██║██║ ╚████║        ██║  ██║██║  ██║███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝        ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
*/

/*
	AdminArea{}.GetAdminArea({Your_Admin_Area_Name}).Code
*/

// GetAdminArea returns information for a given admin domain given a name.
func (a AdminArea) GetAdminArea(adminAreaName string) ResponseGetAdminAreaJSON {
	setAddress := getAdminAreaLink(adminAreaName)

	query, err := rest.Query{}.GetQuery(setAddress, false)
	errc.ErrorCenter("GetAdminArea-GetQuery", err)

	var getAllAdminAreaA ResponseGetAllAdminAreaJSON
	var getAdminArea ResponseGetAdminAreaJSON

	if query != nil {
		errJSON := json.Unmarshal(query, &getAllAdminAreaA)
		errc.ErrorCenter("GetAdminArea-JSONUnmarshal", errJSON)

		if getAllAdminAreaA != nil {
			getAdminArea = getAllAdminAreaA[0]
			return getAdminArea
		}
		return getAdminArea
	}
	return getAdminArea
}

/*
 ██████╗ ███████╗████████╗         █████╗ ██╗     ██╗              █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗         █████╗ ██████╗ ███████╗ █████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██║     ██║             ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║        ██╔══██╗██╔══██╗██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ███████║██║     ██║             ███████║██║  ██║██╔████╔██║██║██╔██╗ ██║        ███████║██████╔╝█████╗  ███████║
██║   ██║██╔══╝     ██║           ██╔══██║██║     ██║             ██╔══██║██║  ██║██║╚██╔╝██║██║██║╚██╗██║        ██╔══██║██╔══██╗██╔══╝  ██╔══██║
╚██████╔╝███████╗   ██║           ██║  ██║███████╗███████╗        ██║  ██║██████╔╝██║ ╚═╝ ██║██║██║ ╚████║        ██║  ██║██║  ██║███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝╚══════╝╚══════╝        ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝        ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
*/

/*
	AdminArea{}.GetAllAdminArea()
*/

// GetAllAdminArea returns information for all existing admin area.
func (a AdminArea) GetAllAdminArea() ResponseGetAllAdminAreaJSON {
	setAddress := getAllAdminAreaLink()

	query, err := rest.Query{}.GetQuery(setAddress, false)
	errc.ErrorCenter("GetAllAdminArea-GetQuery", err)
	var getAdminArea ResponseGetAllAdminAreaJSON

	if query != nil {
		errJSON := json.Unmarshal(query, &getAdminArea)
		errc.ErrorCenter("GetAllAdminArea-JSONUnmarshal", errJSON)
		if getAdminArea != nil {
			return getAdminArea
		}
		return nil
	}
	return nil
}
