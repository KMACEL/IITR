package adminarea

// Posted by Mehmet Akasayan

import (
	"encoding/json"
	"fmt"

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

//MoveAdminArea
func (a AdminArea) MoveAdminArea(adr AdminAreaRequirements) string {
	setAddress := moveAdminAreaLink()

	var adminAreaBodyJSON AdminAreaBodyJSON

	for _, addDevice := range adr.AddToAdminAreaDeviceCode {
		adminAreaBodyJSON.Devices = append(adminAreaBodyJSON.Devices, CodeJSON{Code: addDevice})
	}
	adminAreaBodyJSON.Name = adr.AdminAreaName

	jsonConvert, _ := json.Marshal(adminAreaBodyJSON)
	fmt.Println(string(jsonConvert))

	setBody := string(jsonConvert)

	query, _ := rest.Query{}.PutQuery(setAddress, setBody, contentTypeJSON(), true)
	fmt.Println(string(query))
	//	errc.ErrorCenter(removeApplicationErrorTag, removeError)

	/*if query != nil {
		json.Unmarshal(query, &responseMessageCodeJSONVariable)
		return responseMessageCodeJSONVariable.Response
	}*/
	return ""
}
