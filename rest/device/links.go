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
	builtIn                = "/apps?type=builtin"
	modePolicy             = "/current-and-active-profile"
	apps                   = "/apps"
	clearAppData           = "/clearappdata"
	startApp               = "/startapp"
	stopApp                = "/stopapp"
	control                = "/control/"
	reboot                 = "reboot"
	deviceProfile          = "/device-profile"
	presence               = "?command=Presence"
	applicationInfo        = "?command=ApplicationInfo"
	osProfileInfo          = "?command=OSProfile"
	instantApplicationInfo = "?command=InstantApplicationInfo"
	packageName            = "{\"packageName\":\""
	packageNameEnd         = "\"}"
	deviceLogs             = "deviceLogs/"
	uploadlog              = "uploadlog"
	summary                = "summary"
	deviceParam            = "?device="
	devicesParam           = "?devices="
	location               = "location"
)

//LocationMapLink is return
func locationMapLink() string {
	return device + locationMap
}

//ApplicationDownloadedLink is return
func applicationDownloadedLink(setDeviceCode string) string {
	return device + setDeviceCode + downloaded
}

//ApplicationDownloadedLink is return
func applicationBuiltInLink(setDeviceCode string) string {
	return device + setDeviceCode + builtIn
}

//ModePolicyLink is return
func modePolicyLink(setDeviceID string) string {
	return device + setDeviceID + modePolicy
}

//StartAppLink is return
func startAppLink(setDeviceCode string) string {
	return device + setDeviceCode + apps + startApp
}

//StopAppLink is return
func stopAppLink(setDeviceCode string) string {
	return device + setDeviceCode + apps + stopApp
}

func clearAppDataLink(setDeviceCode string) string {
	return device + setDeviceCode + apps + clearAppData
}

func applicationOperationsBodyLink(setApplicationPackage string) string {
	return packageName + setApplicationPackage + packageNameEnd
}

//MessageControlLink is return
func messageControlLink(setDeviceCode string, setResponseMessageID string) string {
	return device + setDeviceCode + control + setResponseMessageID
}

//RebootLink is return
func rebootLink(setDeviceCode string) string {
	return device + setDeviceCode + control + reboot
}

//PresenceInfoLink is return
func presenceInfoLink(setDeviceID string) string {
	return api + setDeviceID + deviceProfile + presence
}

//PresenceInfoLink is return
func applicationInfoLink(setDeviceID string) string {
	return api + setDeviceID + deviceProfile + applicationInfo
}

//OSProfileInfo is return
func osProfileInfoLink(setDeviceID string) string {
	return api + setDeviceID + deviceProfile + osProfileInfo
}

//instantApplicationInfoLink is return
func instantApplicationInfoLink(setDeviceID string) string {
	return api + setDeviceID + deviceProfile + instantApplicationInfo
}

func getLogListLink(setDeviceCode string) string {
	return api + deviceLogs + setDeviceCode
}

func getDeviceLogLink(setDeviceCode string) string {
	return device + setDeviceCode + control + uploadlog
}

func informationLink(setDeviceCode string) string {
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

func removeApplicationLink(setDeviceID string) string {
	return device + setDeviceID + apps
}

func removeApplicationBody(setPackageName string) string {
	return packageName + setPackageName + packageNameEnd
}

//
