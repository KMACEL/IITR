package workingset

import (
	"encoding/json"
	"fmt"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

/*
███████╗███████╗███╗   ██╗██████╗         ██████╗ ██╗ ██████╗██╗  ██╗        ███╗   ███╗███████╗███████╗███████╗ █████╗  ██████╗ ███████╗
██╔════╝██╔════╝████╗  ██║██╔══██╗        ██╔══██╗██║██╔════╝██║  ██║        ████╗ ████║██╔════╝██╔════╝██╔════╝██╔══██╗██╔════╝ ██╔════╝
███████╗█████╗  ██╔██╗ ██║██║  ██║        ██████╔╝██║██║     ███████║        ██╔████╔██║█████╗  ███████╗███████╗███████║██║  ███╗█████╗
╚════██║██╔══╝  ██║╚██╗██║██║  ██║        ██╔══██╗██║██║     ██╔══██║        ██║╚██╔╝██║██╔══╝  ╚════██║╚════██║██╔══██║██║   ██║██╔══╝
███████║███████╗██║ ╚████║██████╔╝        ██║  ██║██║╚██████╗██║  ██║        ██║ ╚═╝ ██║███████╗███████║███████║██║  ██║╚██████╔╝███████╗
╚══════╝╚══════╝╚═╝  ╚═══╝╚═════╝         ╚═╝  ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝        ╚═╝     ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝
*/

//SendRichMessage is
func (w Workingset) SendRichMessage(message string, messageType string, timeType string, time int64, workingSetKeyExternal string, deviceID ...string) bool {
	var (
		workingsetVariables Workingset
		workingSetKey       string
	)

	if deviceID != nil && workingSetKeyExternal == "" {
		workingSetKey = workingsetVariables.CreateWorkingset()
		for _, devices := range deviceID {
			workingsetVariables.AddDeviceWorkingSet(workingSetKey, device.Device{}.DeviceID2Code(devices))
		}
		fmt.Println("WorkingsetKey : ", workingSetKey, "\nWorkingset Device List : ", w.GetWorkingsetDevices(workingSetKey))
	} else {
		fmt.Println("WorkingsetKey : ", workingSetKeyExternal, "\nWorkingset Device List : ", w.GetWorkingsetDevices(workingSetKeyExternal))
		workingSetKey = workingSetKeyExternal
	}

	setQueryAddress := sendRichMessage(workingSetKey)
	// TODO: json olarak veriyi al
	body := sendRichMessageBody(message, messageType, timeType, time)

	query, err := queryVariable.PostQuery(setQueryAddress, body, contentTypeJSON(), true)
	errc.ErrorCenter("Push Application :", err)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responsePushApplicationJSONVariable)
			//TODO: succes bilgisini kontrol et
			return true
		}
		return false

	}
	return false
}
