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
	workingset         = "https://api.ardich.com/api/v3/workingset/"
	empty              = "empty"
	deviceAdd          = "/devices/add/"
	applicationInstall = "/application/install"
	devices            = "/device"
)

//createWorkingsetLink is retrun
func createWorkingsetLink() string {
	return workingset + empty
}

//addDeviceWorkingSetLink is
func addDeviceWorkingSetLink(setWorkingset string) string {
	return workingset + setWorkingset + deviceAdd
}

func pushApplicationsLink(workingsetKey string) string {
	return workingset + workingsetKey + applicationInstall
}

func pushApplicationsBody(applicationCode string, notifyUser bool) string {
	return `{"apps": [{"code": "` + applicationCode + `"}],"notifyUser":` + strconv.FormatBool(notifyUser) + `}`
}

func getWorkingsetDevicesLink(workingsetKey string) string {
	return  workingset + workingsetKey + devices
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
