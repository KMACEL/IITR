package device

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/KMACEL/IITR/rest"
)

/*
███╗   ███╗███████╗███████╗███████╗ █████╗  ██████╗ ███████╗         ██████╗ ██████╗ ███╗   ██╗████████╗██████╗  ██████╗ ██╗
████╗ ████║██╔════╝██╔════╝██╔════╝██╔══██╗██╔════╝ ██╔════╝        ██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██╔═══██╗██║
██╔████╔██║█████╗  ███████╗███████╗███████║██║  ███╗█████╗          ██║     ██║   ██║██╔██╗ ██║   ██║   ██████╔╝██║   ██║██║
██║╚██╔╝██║██╔══╝  ╚════██║╚════██║██╔══██║██║   ██║██╔══╝          ██║     ██║   ██║██║╚██╗██║   ██║   ██╔══██╗██║   ██║██║
██║ ╚═╝ ██║███████╗███████║███████║██║  ██║╚██████╔╝███████╗        ╚██████╗╚██████╔╝██║ ╚████║   ██║   ██║  ██║╚██████╔╝███████╗
╚═╝     ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝         ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝
*/

//MessageControl is
func (d Device) MessageControl(deviceCode string, responseMessageID string, visualFlag bool) (string, string) {
	var setVisualFlag bool

retry:
	setQueryAddress := messageControlLink(deviceCode, responseMessageID)

	if visualFlag == rest.Visible {
		setVisualFlag = rest.Visible
	} else if visualFlag == rest.Invisible {
		setVisualFlag = rest.Invisible
	}

	query, _ := q.GetQuery(setQueryAddress, setVisualFlag)

	var counter int

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responseMessageJSONVariable)
			responseResult := responseMessageJSONVariable.ExecutionInDeviceStatus.Result

			if responseResult != "" {
				json.Unmarshal([]byte(responseResult), &responseDescriptionJSONVariable)

				responseCode := responseDescriptionJSONVariable.Code
				if responseCode == 200 {
					return strconv.Itoa(responseDescriptionJSONVariable.Code), responseDescriptionJSONVariable.State
				} else if responseCode == 500 {
					json.Unmarshal([]byte(responseResult), &responseMessageErrorJSONVariable)
					return strconv.Itoa(responseMessageErrorJSONVariable.Code), responseMessageErrorJSONVariable.Description
				}
			}
			counter++
			if counter < 3 {
				time.Sleep(1000 * time.Millisecond)
				goto retry
			} else {
				return rest.ResponseNil, responseDescriptionJSONVariable.State
			}

		}
		return rest.ResponseNotFound, ""

	}
	return rest.ResponseNil, ""
}
