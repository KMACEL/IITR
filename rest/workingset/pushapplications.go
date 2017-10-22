package workingset

import (
	"github.com/KMACEL/IITR/rest/device"
	"fmt"
	"github.com/KMACEL/IITR/errc"
	"encoding/json"
	"github.com/KMACEL/IITR/rest"
)

// workingset.Workingset{}.PushApplications("DD76AFEA-E0A3-4B61-97CA-509B66A884E1", false, "867377020740787","867377020747089")
// PushApplications is
func (w Workingset) PushApplications(applicationCode string, notifyUser bool, deviceID ...string) bool {
	var workingsetVariables Workingset

	workingsetKey := workingsetVariables.CreateWorkingset()

	for _, devices := range deviceID {
		workingsetVariables.AddDeviceWorkingSet(workingsetKey, device.Device{}.DeviceID2Code(devices))
	}

	fmt.Println("Workingset Device List : ", w.GetWorkingsetDevices(workingsetKey))

	setQueryAddress := pushApplicationsLink(workingsetKey)
	body := pushApplicationsBody(applicationCode, notifyUser)

	query,err:=queryVariable.PostQuery(setQueryAddress, body, contentTypeJSON(), true)
	errc.ErrorCenter("Push Application :",err)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			json.Unmarshal(query, &responsePushApplicationJSONVariable)
			//todo : succes bilgisini kontrol et
			return true
		}
		return false

	}
	return false

}
