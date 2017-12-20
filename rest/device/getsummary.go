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
func (d Device) Summary(setDeviceID string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := summaryLink(setDeviceID)
	query, errSummary := q.GetQuery(setQueryAddress, visualFlag)
	errc.ErrorCenter(summaryTag, errSummary)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				json.Unmarshal(query, &summaryJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
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
func (d Device) WorkingGroupControl(setDeviceID string, visualFlag bool) string {
	query := d.Summary(setDeviceID, rest.OKMarshal, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			summaryWorkingGroup := SummaryJSON{}
			json.Unmarshal(query, &summaryWorkingGroup)
			var workingGroupString string

			if summaryWorkingGroup.Content != nil {
				if summaryWorkingGroup.Content[0].Groups != nil {
					for _, workingGroup := range summaryWorkingGroup.Content[0].Groups {
						workingGroupString = workingGroup.Name + workingGroupString
					}
					return workingGroupString
				}
				return rest.ResponseNotFound
			}
			return rest.ResponseNil
		}
		return rest.ResponseNil
	}
	return rest.ResponseNil
}
