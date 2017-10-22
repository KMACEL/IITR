package device

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
	api                    = "https://api.ardich.com/api/v3/"
	device                 = "https://api.ardich.com:443/api/v3/device/"
	locationMap            = "device-location-map"
	downloaded             = "/apps?type=downloaded"
	modePolicy             = "/current-and-active-profile"
	stopApp                = "/apps/stopapp"
	startApp               = "/apps/startapp"
	control                = "/control/"
	reboot                 = "reboot"
	presence               = "/device-profile?command=Presence"
	applicationInfo        = "/device-profile?command=ApplicationInfo"
	osProfileInfo          = "/device-profile?command=OSProfile"
	instantApplicationInfo = "/device-profile?command=InstantApplicationInfo"
	packageName            = "{\"packageName\":\""
	packageNameEnd         = "\"}\""
	deviceLogs             = "deviceLogs/"
	uploadlog              = "uploadlog"
	summary                = "summary"
	deviceParam            = "?device="
	devicesParam           = "?devices="
	location               = "location"
)

//LocationMapLink is retrun
func locationMapLink() string {
	return device + locationMap
}

//ApplicationDownloadedLink is retrun
func applicationDownloadedLink(setDeviceCode string) string {
	return device + setDeviceCode + downloaded
}

//ModePolicyLink is retrun
func modePolicyLink(setDeviceID string) string {
	return device + setDeviceID + modePolicy
}

//StartAppLink is retrun
func startAppLink(setDeviceCode string) string {
	return device + setDeviceCode + startApp
}

//StopAppLink is retrun
func stopAppLink(setDeviceCode string) string {
	return device + setDeviceCode + stopApp
}

func applicationOperationsBodyLink(setApplicationPackage string) string {
	return packageName + setApplicationPackage + packageNameEnd
}

//MessageControlLink is retrun
func messageControlLink(setDeviceCode string, setResponseMessageID string) string {
	return device + setDeviceCode + control + setResponseMessageID
}

//RebootLink is retrun
func rebootLink(setDeviceCode string) string {
	return device + setDeviceCode + control + reboot
}

//PresenceInfoLink is retrun
func presenceInfoLink(setDeviceID string) string {
	return api + setDeviceID + presence
}

//PresenceInfoLink is retrun
func applicationInfoLink(setDeviceID string) string {
	return api + setDeviceID + applicationInfo
}

//OSProfileInfo is retrun
func osProfileInfoLink(setDeviceID string) string {
	return api + setDeviceID + osProfileInfo
}

//instantApplicationInfoLink is retrun
func instantApplicationInfoLink(setDeviceID string) string {
	return api + setDeviceID + instantApplicationInfo
}

func getLogListLink(setDeviceCode string) string {
	return api + deviceLogs + setDeviceCode
}

func getDeviceLogLink(setDeviceCode string) string {
	return device + setDeviceCode + control + uploadlog
}

func informationsLink(setDeviceCode string) string {
	return device + setDeviceCode
}

func summaryLink(deviceID string) string {
	return device + summary + deviceParam + deviceID
}

func deviceID2CodeLink(setDeviceID string) string {
	return device + location + devicesParam + setDeviceID
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
