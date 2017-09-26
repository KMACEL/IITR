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
func (d Device) MessageControl(deviceCode string, responseMessageID string, vasualFlag bool) (string, string) {
	var setVasualFlag bool

retry:
	setQueryAdress := messageControlLink(deviceCode, responseMessageID)

	if vasualFlag == rest.Visible {
		setVasualFlag = rest.Visible
	} else if vasualFlag == rest.Invisible {
		setVasualFlag = rest.Invisible
	}

	query, _ := rest.Query{}.GetQuery(setQueryAdress, setVasualFlag)

	var counter int

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responseMessageJSONVariable)
			responseResult := responseMessageJSONVariable.ExecutionInDeviceStatus.Result

			if responseResult != "" {
				json.Unmarshal([]byte(responseResult), &responsedescriptionJSONVariable)

				responseCode := responsedescriptionJSONVariable.Code
				if responseCode == 200 {
					return strconv.Itoa(responsedescriptionJSONVariable.Code), responsedescriptionJSONVariable.State
				} else if responseCode == 500 {
					json.Unmarshal([]byte(responseResult), &responseMesageErrorJSONVariable)
					return strconv.Itoa(responseMesageErrorJSONVariable.Code), responseMesageErrorJSONVariable.Description
				}

			}
			counter++
			if counter < 3 {
				time.Sleep(1000 * time.Millisecond)
				goto retry
			} else {
				return rest.ResponseNil, responsedescriptionJSONVariable.State
			}

		}
		return rest.ResponseNotFound, ""

	}
	return rest.ResponseNil, ""
}
