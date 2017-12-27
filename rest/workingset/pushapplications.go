package workingset

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

/*
██████╗ ██╗   ██╗███████╗██╗  ██╗             █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗███████╗
██╔══██╗██║   ██║██╔════╝██║  ██║            ██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║██╔════╝
██████╔╝██║   ██║███████╗███████║            ███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║███████╗
██╔═══╝ ██║   ██║╚════██║██╔══██║            ██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║╚════██║
██║     ╚██████╔╝███████║██║  ██║            ██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║███████║
╚═╝      ╚═════╝ ╚══════╝╚═╝  ╚═╝            ╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝
*/
// Example :
//     workingset.Workingset{}.PushApplications("DD76AFEA-E0A3-4B61-97CA-509B66A884E1", false, "867377020740787","867377020747089")

// PushApplications is
func (w Workingset) PushApplications(applicationCode string, notifyUser bool, deviceID ...string) bool {
	var workingsetVariables Workingset

	workingsetKey := workingsetVariables.CreateWorkingset()

	for _, devices := range deviceID {
		workingsetVariables.AddDeviceWorkingSet(workingsetKey, device.Device{}.DeviceID2Code(devices))
	}
	// todo workingsete array olarak ver
	fmt.Println("Workingset Device List : ", w.GetWorkingsetDevices(workingsetKey))

	setQueryAddress := pushApplicationsLink(workingsetKey)
	body := pushApplicationsBody(applicationCode, notifyUser)

	query, err := queryVariable.PostQuery(setQueryAddress, body, contentTypeJSON(), true)
	errc.ErrorCenter("Push Application :", err)

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

/*
██████╗ ██╗   ██╗███████╗██╗  ██╗         █████╗ ██████╗ ██████╗         ███████╗██╗  ██╗████████╗███████╗██████╗ ███╗   ██╗ █████╗ ██╗
██╔══██╗██║   ██║██╔════╝██║  ██║        ██╔══██╗██╔══██╗██╔══██╗        ██╔════╝╚██╗██╔╝╚══██╔══╝██╔════╝██╔══██╗████╗  ██║██╔══██╗██║
██████╔╝██║   ██║███████╗███████║        ███████║██████╔╝██████╔╝        █████╗   ╚███╔╝    ██║   █████╗  ██████╔╝██╔██╗ ██║███████║██║
██╔═══╝ ██║   ██║╚════██║██╔══██║        ██╔══██║██╔═══╝ ██╔═══╝         ██╔══╝   ██╔██╗    ██║   ██╔══╝  ██╔══██╗██║╚██╗██║██╔══██║██║
██║     ╚██████╔╝███████║██║  ██║        ██║  ██║██║     ██║             ███████╗██╔╝ ██╗   ██║   ███████╗██║  ██║██║ ╚████║██║  ██║███████╗
╚═╝      ╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝     ╚═╝             ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
*/

// PushApplicationsExternal is
func (w Workingset) PushApplicationsExternal(fileName string, url string, notifyUser bool, deviceID ...string) bool {

	var workingsetVariables Workingset

	workingsetKey := workingsetVariables.CreateWorkingset()

	for _, devices := range deviceID {
		workingsetVariables.AddDeviceWorkingSet(workingsetKey, device.Device{}.DeviceID2Code(devices))
	}
	// todo workingsete array olarak ver
	fmt.Println("Workingset Device List : ", w.GetWorkingsetDevices(workingsetKey))

	setQueryAddress := pushApplicationsExternalLink(workingsetKey)

	// todo json olarak veriyi al
	body := `{
	  "deviceIds"	:[],
	"expireDate":	0,
	"fileName":	"` + fileName + `",
	"notifyUser":` + strconv.FormatBool(notifyUser) + `,
	"packageName":""	,
	"token":""	,
	"url":	"` + url + `",
	"versionCode":	"0"
	}`

	query, err := queryVariable.PostQuery(setQueryAddress, body, contentTypeJSON(), true)
	errc.ErrorCenter("Push Application :", err)

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

/*
██╗   ██╗███╗   ██╗██╗███╗   ██╗███████╗████████╗ █████╗ ██╗     ██╗                   ██╗███╗   ██╗███████╗████████╗ █████╗ ██╗     ██╗          █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗
██║   ██║████╗  ██║██║████╗  ██║██╔════╝╚══██╔══╝██╔══██╗██║     ██║                   ██║████╗  ██║██╔════╝╚══██╔══╝██╔══██╗██║     ██║         ██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║
██║   ██║██╔██╗ ██║██║██╔██╗ ██║███████╗   ██║   ███████║██║     ██║         █████╗    ██║██╔██╗ ██║███████╗   ██║   ███████║██║     ██║         ███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║
██║   ██║██║╚██╗██║██║██║╚██╗██║╚════██║   ██║   ██╔══██║██║     ██║         ╚════╝    ██║██║╚██╗██║╚════██║   ██║   ██╔══██║██║     ██║         ██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║
╚██████╔╝██║ ╚████║██║██║ ╚████║███████║   ██║   ██║  ██║███████╗███████╗              ██║██║ ╚████║███████║   ██║   ██║  ██║███████╗███████╗    ██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║
 ╚═════╝ ╚═╝  ╚═══╝╚═╝╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚══════╝              ╚═╝╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚══════╝    ╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
*/

//UninstallInstallApplication is
func (w Workingset) UninstallInstallApplication(applicationCode string, notifyUser bool, deviceID ...string) bool {
	var workingsetVariables Workingset
	workingsetKey := workingsetVariables.CreateWorkingset()
	for _, devices := range deviceID {
		workingsetVariables.AddDeviceWorkingSet(workingsetKey, device.Device{}.DeviceID2Code(devices))
	}
	// todo workingsete array olarak ver
	fmt.Println("Workingset Device List : ", w.GetWorkingsetDevices(workingsetKey))

	setQueryAddress := uninstallInstallApplicationLink(workingsetKey)

	body := `{
		"notifyUser":` + strconv.FormatBool(notifyUser) + `,
		"apps": [
			{
				"code": "` + applicationCode + `"
			}
		]
	}`

	query, err := queryVariable.PostQuery(setQueryAddress, body, contentTypeJSON(), true)
	errc.ErrorCenter("Push Application :", err)

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
