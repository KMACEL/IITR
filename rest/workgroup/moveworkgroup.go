package workgroup

import (
	"encoding/json"
	"fmt"

	"github.com/KMACEL/IITR/rest"
)

/*
███╗   ███╗ ██████╗ ██╗   ██╗███████╗        ██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗ ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗
████╗ ████║██╔═══██╗██║   ██║██╔════╝        ██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗
██╔████╔██║██║   ██║██║   ██║█████╗          ██║ █╗ ██║██║   ██║██████╔╝█████╔╝ ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝
██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██╔══╝          ██║███╗██║██║   ██║██╔══██╗██╔═██╗ ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝
██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ███████╗        ╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║
╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚══════╝         ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝
*/

/*
	var w workgroup.WorkGroupRequirements
	w.AddToWorkGroupDeviceCode = []string{"{YOUR_DEVICE}", "{YOUR_DEVICE}"}
	w.WorkGroupName = "{NEW_GROUP_NAME}" // new group
	w.GroupCodes = []string{workgroup.WorkGroup{}.GetGroupID("{GROUP_NAME}", false)}
	workgroup.WorkGroup{}.MoveWorkGroup(w)
*/

// MoveWorkGroup is
func (a WorkGroup) MoveWorkGroup(wgr WorkGroupRequirements) string {
	setAddress := moveWorkGroupLink()
	var workGroupBodyJSON WorkGroupBodyJSON

	for _, addDevice := range wgr.AddToWorkGroupDeviceCode {
		workGroupBodyJSON.Devices = append(workGroupBodyJSON.Devices, CodeJSON{Code: addDevice})
	}
	workGroupBodyJSON.Name = wgr.WorkGroupName
	workGroupBodyJSON.GroupCodes = wgr.GroupCodes

	jsonConvert, _ := json.Marshal(workGroupBodyJSON)
	fmt.Println(string(jsonConvert))

	setBody := string(jsonConvert)

	query, err := rest.Query{}.PostQuery(setAddress, setBody, contentTypeJSON(), true)
	if err != nil {
		panic(err)
	}
	fmt.Println("PostQuery Result: ", string(query))
	//	errc.ErrorCenter(removeApplicationErrorTag, removeError)

	/*if query != nil {
		json.Unmarshal(query, &responseMessageCodeJSONVariable)
		return responseMessageCodeJSONVariable.Response
	}*/
	return ""
}
