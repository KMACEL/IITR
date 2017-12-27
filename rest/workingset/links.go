package workingset

import "strconv"

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
	workingset          = "https://api.ardich.com/api/v3/workingset/"
	empty               = "empty"
	deviceAdd           = "/devices/add/"
	application         = "/application/"
	applicationInstall  = "install"
	applicationReistall = "reinstall"
	externalApp         = "install-external-app"
	devices             = "/device?page=0&size=500"
	controlAd           = "/control/ad"
)

//createWorkingsetLink is return
func createWorkingsetLink() string {
	return workingset + empty
}

//addDeviceWorkingSetLink is
func addDeviceWorkingSetLink(setWorkingset string) string {
	return workingset + setWorkingset + deviceAdd
}

func pushApplicationsLink(workingsetKey string) string {
	return workingset + workingsetKey + application + applicationInstall
}

func pushApplicationsBody(applicationCode string, notifyUser bool) string {
	return `{"apps": [{"code": "` + applicationCode + `"}],"notifyUser":` + strconv.FormatBool(notifyUser) + `}`
}

func pushApplicationsExternalLink(workingsetKey string) string {
	return workingset + workingsetKey + application + externalApp
}

func uninstallInstallApplicationLink(workingsetKey string) string {
	return workingset + workingsetKey + application + applicationReistall
}

func pushApplicationsExternalBody(applicationCode string, notifyUser bool) string {
	return `{"apps": [{"code": "` + applicationCode + `"}],"notifyUser":` + strconv.FormatBool(notifyUser) + `}`
}

func getWorkingsetDevicesLink(workingsetKey string) string {
	return workingset + workingsetKey + devices
}

func sendRichMessage(workingsetKey string) string {
	return workingset + workingsetKey + controlAd
}

func sendRichMessageBody(message string, messageType string, timeType string, time int64) string {
	return `{"message":"` + message + `","type":"` + messageType + `","timeType":"` + timeType + `","time":` + strconv.FormatInt(time, 10) + `}`
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
