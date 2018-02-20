package workgroup

// Posted by Mehmet Akasayan

import (
	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗         ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗         ██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗███████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗        ██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝██╔════╝
██║  ███╗█████╗     ██║           ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝        ██║  ██║█████╗  ██║   ██║██║██║     █████╗  ███████╗
██║   ██║██╔══╝     ██║           ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝         ██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝  ╚════██║
╚██████╔╝███████╗   ██║           ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║             ██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗███████║
 ╚═════╝ ╚══════╝   ╚═╝            ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝             ╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝╚══════╝
*/

// GetGroupDevices is
func (w WorkGroup) GetGroupDevices(workGroupID string, visualFlag bool) []byte {
	setQueryAddress := getGroupDevicesLink(workGroupID)
	query, _ := rest.Query{}.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			//			json.Unmarshal(query, &getGroupDevicesJSONVariable)
			return query
		}
		return nil
	}
	return nil
}
