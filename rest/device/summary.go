package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
███████╗██╗   ██╗███╗   ███╗███╗   ███╗ █████╗ ██████╗ ██╗   ██╗
██╔════╝██║   ██║████╗ ████║████╗ ████║██╔══██╗██╔══██╗╚██╗ ██╔╝
███████╗██║   ██║██╔████╔██║██╔████╔██║███████║██████╔╝ ╚████╔╝
╚════██║██║   ██║██║╚██╔╝██║██║╚██╔╝██║██╔══██║██╔══██╗  ╚██╔╝
███████║╚██████╔╝██║ ╚═╝ ██║██║ ╚═╝ ██║██║  ██║██║  ██║   ██║
╚══════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝
*/

// Summary is
func (d Device) Summary(setDeviceID string, setUnMarshal bool, vasualFlag bool) []byte {
	setQueryAdress := summaryLink(setDeviceID)
	query, errSummary := queryVariable.GetQuery(setQueryAdress, vasualFlag)
	errc.ErrorCenter(errc.Summary, errSummary)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &summaryJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return []byte(rest.ResponseNil)
}

/*
██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗██╗███╗   ██╗ ██████╗          ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗          ██████╗ ██████╗ ███╗   ██╗████████╗██████╗  ██████╗ ██╗
██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝██║████╗  ██║██╔════╝         ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗        ██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██╔═══██╗██║
██║ █╗ ██║██║   ██║██████╔╝█████╔╝ ██║██╔██╗ ██║██║  ███╗        ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝        ██║     ██║   ██║██╔██╗ ██║   ██║   ██████╔╝██║   ██║██║
██║███╗██║██║   ██║██╔══██╗██╔═██╗ ██║██║╚██╗██║██║   ██║        ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝         ██║     ██║   ██║██║╚██╗██║   ██║   ██╔══██╗██║   ██║██║
╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗██║██║ ╚████║╚██████╔╝        ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║             ╚██████╗╚██████╔╝██║ ╚████║   ██║   ██║  ██║╚██████╔╝███████╗
 ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝          ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝              ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝
*/

// WorkingGroupControl is
func (d Device) WorkingGroupControl(setDeviceID string, vasualFlag bool) string {
	query := d.Summary(setDeviceID, rest.OKMarshal, vasualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			summaryWorkingGroup := SummaryJSON{}
			json.Unmarshal(query, &summaryWorkingGroup)
			var workingGroupString string
			for _, workingGroup := range summaryWorkingGroup.Content[0].Groups {
				workingGroupString = workingGroup.Name + workingGroupString
			}

			return workingGroupString
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil
}
