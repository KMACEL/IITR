package workgroup

// Posted by Mehmet Akasayan

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗         ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗         ██╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗        ██║██╔══██╗
██║  ███╗█████╗     ██║           ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝        ██║██║  ██║
██║   ██║██╔══╝     ██║           ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝         ██║██║  ██║
╚██████╔╝███████╗   ██║           ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║             ██║██████╔╝
 ╚═════╝ ╚══════╝   ╚═╝            ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝             ╚═╝╚═════╝
*/

// GetGroupID
func (w WorkGroup) GetGroupID(workGroupName string, visualFlag bool) string {
	setQueryAddress := getGroupIDLink(workGroupName)
	query, _ := rest.Query{}.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &getWorkGroupIDJSONVariable)
			return getWorkGroupIDJSONVariable.Content[0].Code
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil

}
