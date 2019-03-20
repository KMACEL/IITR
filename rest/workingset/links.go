package workingset

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/KMACEL/IITR/rest"
)

/*
██╗     ██╗███╗   ██╗██╗  ██╗███████╗
██║     ██║████╗  ██║██║ ██╔╝██╔════╝
██║     ██║██╔██╗ ██║█████╔╝ ███████╗
██║     ██║██║╚██╗██║██╔═██╗ ╚════██║
███████╗██║██║ ╚████║██║  ██╗███████║
╚══════╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
*/
//This page is the part that shows the links that the queries will use.
//It is designed in such a way that the administration is easy.

const (
	workingset         = "workingset/"
	empty              = "empty"
	devices            = "/devices/"
	add                = "add"
	application        = "/application/"
	install            = "install"
	installExternalApp = "install-external-app"
	reinstall          = "reinstall"
	deviceW            = "/device"
	control            = "/control/"
	ad                 = "ad"
)

/*const (
	workingset          = "https://api.ardich.com/api/v3/workingset/"
	empty               = "empty"
	deviceAdd           = "/devices/add/"
	application         = "/application/"
	applicationInstall  = "install"
	applicationReistall = "reinstall"
	externalApp         = "install-external-app"
	devices             = "/device?page=0&size=500"
	controlAd = "/control/ad"
)*/

//createWorkingsetLink is return
//https://api.ardich.com/api/v3/workingset/empty
func createWorkingsetLink() string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + workingset + empty

	return u.String()
}

// https://api.ardich.com:443/api/v3/workingset/{WORKING_SET_CODE}/devices/add
func addDeviceWorkingSetLink(setWorkingsetCode string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + workingset + setWorkingsetCode + devices + add

	return u.String()
}

// https://api.ardich.com:443/api/v3/workingset/{WORKING_SET_CODE}/application/install
func pushApplicationsLink(setWorkingsetCode string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + workingset + setWorkingsetCode + application + install

	return u.String()
}

// https://api.ardich.com:443/api/v3/workingset/{WORKING_SET_CODE}/application/install-external-app
func pushApplicationsExternalLink(setWorkingsetCode string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + workingset + setWorkingsetCode + application + installExternalApp

	return u.String()
}

// https://api.ardich.com:443/api/v3/workingset/{WORKING_SET_CODE}/application/reinstall
func uninstallInstallApplicationLink(setWorkingsetCode string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + workingset + setWorkingsetCode + application + reinstall

	return u.String()
}

// https://api.ardich.com:443/api/v3/workingset/{WORKING_SET_CODE}/control/ad
func sendRichMessage(setWorkingsetCode string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + workingset + setWorkingsetCode + control + ad

	return u.String()
}

// https://api.ardich.com:443/api/v3/workingset/{WORKING_SET_CODE}/device
func getWorkingsetDevicesLink(setWorkingsetCode string) string {
	data := url.Values{}
	data.Add("page", "0")
	data.Add("size", "500")

	u := rest.GetAPITemplate()
	u.Path = u.Path + workingset + setWorkingsetCode + deviceW
	u.RawQuery = data.Encode()

	return u.String()
}

/*
██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██████╔╝██║   ██║██║  ██║ ╚████╔╝
██╔══██╗██║   ██║██║  ██║  ╚██╔╝
██████╔╝╚██████╔╝██████╔╝   ██║
╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// {"apps":[{"code":"XXXX-YYYY"},{"code":"ZZZZ-TTTT"}],"notifyUser":TRUE/FALSE}
func pushApplicationsBody(notifyUser bool, applicationCode ...string) string {
	var pushApplicationsBodyJSONVar pushApplicationsBodyJSON
	pushApplicationsBodyJSONVar.NotifyUser = notifyUser

	for _, appCode := range applicationCode {
		pushApplicationsBodyJSONVar.Apps = append(pushApplicationsBodyJSONVar.Apps, apps{Code: appCode})
	}
	jsonConvert, _ := json.Marshal(pushApplicationsBodyJSONVar)

	return string(jsonConvert)
}

// {"message": "string", "time": "string", "timeType": "string", "type": "string"}
func sendRichMessageBody(message string, messageType string, timeType string, time int64) string {
	var sendRichMessageBodyJSONVar sendRichMessageBodyJSON
	sendRichMessageBodyJSONVar.Message = message
	sendRichMessageBodyJSONVar.Type = messageType
	sendRichMessageBodyJSONVar.TimeType = timeType
	sendRichMessageBodyJSONVar.Time = strconv.FormatInt(time, 10)
	jsonConvert, _ := json.Marshal(sendRichMessageBodyJSONVar)

	return string(jsonConvert)
}

//
func pushApplicationsExternalBody(fileName string, url string, versionCode int, notifyUser bool, deviceID ...string) string {
	var pushExternalApplicationBodyJSONVar pushExternalApplicationBodyJSON
	pushExternalApplicationBodyJSONVar.FileName = fileName
	pushExternalApplicationBodyJSONVar.NotifyUser = notifyUser
	pushExternalApplicationBodyJSONVar.URL = url
	pushExternalApplicationBodyJSONVar.VersionCode = versionCode

	jsonConvert, _ := json.Marshal(pushExternalApplicationBodyJSONVar)
	return string(jsonConvert)
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
