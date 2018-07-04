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

// GetAdminArea is
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

// GetAllAdminArea is
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
