package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
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

//GetDownloadedApplicationsList is
func (d Device) GetDownloadedApplicationsList(setCode string, setUnMarshal bool, setVisualFlag bool) []byte {
	var q rest.Query

	queryLink := applicationDownloadedLink(setCode)
	visualFlag := setVisualFlag
	query, errDownloadAppList := q.GetQuery(queryLink, visualFlag)
	errc.ErrorCenter(getDownloadApplicationListErrorTag+d.DeviceCode2ID(setCode), errDownloadAppList)

	if query != nil {
		if setUnMarshal {
			json.Unmarshal(query, &downloadedApplicationListJSONVariable)
		}
		return query
	}
	return nil
}

/*
 ██████╗ ███████╗████████╗        ██████╗ ██╗   ██╗██╗██╗  ████████╗              ██╗███╗   ██╗        ██╗     ██╗███████╗████████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██║   ██║██║██║  ╚══██╔══╝              ██║████╗  ██║        ██║     ██║██╔════╝╚══██╔══╝
██║  ███╗█████╗     ██║           ██████╔╝██║   ██║██║██║     ██║       █████╗    ██║██╔██╗ ██║        ██║     ██║███████╗   ██║
██║   ██║██╔══╝     ██║           ██╔══██╗██║   ██║██║██║     ██║       ╚════╝    ██║██║╚██╗██║        ██║     ██║╚════██║   ██║
╚██████╔╝███████╗   ██║           ██████╔╝╚██████╔╝██║███████╗██║                 ██║██║ ╚████║        ███████╗██║███████║   ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═════╝  ╚═════╝ ╚═╝╚══════╝╚═╝                 ╚═╝╚═╝  ╚═══╝        ╚══════╝╚═╝╚══════╝   ╚═╝
*/

//GetBuiltInApplicationsList is
func (d Device) GetBuiltInApplicationsList(setCode string, setUnMarshal bool, setVisualFlag bool) []byte {
	var q rest.Query

	queryLink := applicationBuiltInLink(setCode)
	visualFlag := setVisualFlag

	query, errBuiltInAppList := q.GetQuery(queryLink, visualFlag)
	errc.ErrorCenter(getBuiltInApplicationListErrorTag, errBuiltInAppList)

	if query != nil {
		if setUnMarshal {
			json.Unmarshal(query, &builtInApplicationListJSONVariable)
		}
		return query
	}
	return nil
}
