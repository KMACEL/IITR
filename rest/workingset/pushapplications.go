package workingset

import (
	"encoding/json"
	"fmt"

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
// For use Example :
//     workingset.Workingset{}.PushApplications([]string{"AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA","BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB"}, false, "1234","1234")

// PushApplications is
func (w Workingset) PushApplications(applicationCode []string, notifyUser bool, deviceID ...string) bool {
	var workingsetVariables Workingset

	workingsetKey := workingsetVariables.CreateWorkingset()
	workingsetVariables.AddDeviceWorkingSet(workingsetKey, deviceID...)
	fmt.Println("Workingset Device List : ", w.GetWorkingsetDevices(workingsetKey))

	setQueryAddress := pushApplicationsLink(workingsetKey)
	body := pushApplicationsBody(notifyUser, applicationCode...)

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
func (w Workingset) PushApplicationsExternal(fileName string, url string, versionCode int, notifyUser bool, deviceID ...string) bool {
	var workingsetVariables Workingset
	workingsetKey := workingsetVariables.CreateWorkingset()
	workingsetVariables.AddDeviceWorkingSet(workingsetKey, deviceID...)
	fmt.Println("Workingset Device List : ", w.GetWorkingsetDevices(workingsetKey))

	setQueryAddress := pushApplicationsExternalLink(workingsetKey)
	body := pushApplicationsExternalBody(fileName, url, versionCode, notifyUser)

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
func (w Workingset) UninstallInstallApplication(applicationCode []string, notifyUser bool, deviceID ...string) bool {
	var workingsetVariables Workingset
	workingsetKey := workingsetVariables.CreateWorkingset()
	for _, devices := range deviceID {
		workingsetVariables.AddDeviceWorkingSet(workingsetKey, device.Device{}.DeviceID2Code(devices))
	}
	// todo workingsete array olarak ver
	fmt.Println("Workingset Device List : ", w.GetWorkingsetDevices(workingsetKey))

	setQueryAddress := uninstallInstallApplicationLink(workingsetKey)

	body := pushApplicationsBody(notifyUser, applicationCode...)

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
