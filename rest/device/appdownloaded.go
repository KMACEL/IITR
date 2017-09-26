package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗        ██████╗  ██████╗ ██╗    ██╗███╗   ██╗██╗      ██████╗  █████╗ ██████╗ ███████╗██████╗         ██╗     ██╗███████╗████████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔═══██╗██║    ██║████╗  ██║██║     ██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗        ██║     ██║██╔════╝╚══██╔══╝
██║  ███╗█████╗     ██║           ██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║██║     ██║   ██║███████║██║  ██║█████╗  ██║  ██║        ██║     ██║███████╗   ██║
██║   ██║██╔══╝     ██║           ██║  ██║██║   ██║██║███╗██║██║╚██╗██║██║     ██║   ██║██╔══██║██║  ██║██╔══╝  ██║  ██║        ██║     ██║╚════██║   ██║
╚██████╔╝███████╗   ██║           ██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║███████╗╚██████╔╝██║  ██║██████╔╝███████╗██████╔╝        ███████╗██║███████║   ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═════╝         ╚══════╝╚═╝╚══════╝   ╚═╝
*/

//GetDownloadedList is
func (d Device) GetDownloadedList(setCode string, setUnMarshal bool, setVasualFlag bool) []byte {
	code := setCode
	queryLink := applicationDownloadedLink(code)
	vasualFlag := setVasualFlag
	query, _ := rest.Query{}.GetQuery(queryLink, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &downloadedApplicationListJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}
