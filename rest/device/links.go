package device

import (
	"net/url"

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
	dev     = "device/"
	control = "/control/"
	status  = "status"
)

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
	reboot                 = "reboot"
	deviceProfile          = "/device-profile"
	deviceProfileHistory   = "/device-profile-history"
	command                = "?command="
	presence               = "Presence"
	applicationInfo        = "ApplicationInfo"
	osProfileInfo          = "OSProfile"
	instantApplicationInfo = "InstantApplicationInfo"
	packageName            = "{\"packageName\":\""
	packageNameEnd         = "\"}"
	deviceLogs             = "deviceLogs/"
	uploadlog              = "uploadlog"
	summary                = "summary"
	deviceParam            = "?device="
	devicesParam           = "?devices="
	location               = "location"
	sensorData             = "/sensor-data"
	nodeID                 = "nodeId="
	sensorID               = "sensorId="
	iotLabel               = "iotlabel/label/multi"
)

// https://api.ardich.com:443/api/v3/device/device-location-map
func locationMapLink() string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + dev
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
	return api + setDeviceID + deviceProfile + command + presence
}

//PresenceInfoLink is return
func applicationInfoLink(setDeviceID string) string {
	return api + setDeviceID + deviceProfile + command + applicationInfo
}

//OSProfileInfo is return
func osProfileInfoLink(setDeviceID string) string {
	return api + setDeviceID + deviceProfile + command + osProfileInfo
}

//instantApplicationInfoLink is return
func instantApplicationInfoLink(setDeviceID string) string {
	return api + setDeviceID + deviceProfile + command + instantApplicationInfo
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

func deviceCode2IDLink(setDeviceCode string) string {
	return api + setDeviceCode + deviceProfile + command + presence
}

func removeApplicationLink(setDeviceID string) string {
	return device + setDeviceID + apps
}

func removeApplicationBody(setPackageName string) string {
	return packageName + setPackageName + packageNameEnd
}

func presenceHistoryLink(deviceID string) string {
	return api + deviceID + deviceProfileHistory + command + presence
}

// https://api.ardich.com/api/v3/device/5a2f475efd6a4fb7ad347131f27e94f3/control/status?command={REFRESH-TYPE}
func refreshGatewayInfoLink(deviceCode string, specificParameter ...string) string {
	u := rest.GetAPITemplate()
	data := url.Values{}
	if specificParameter != nil {
		for _, param := range specificParameter {
			data.Add("command", param)
		}
		u.RawQuery = data.Encode()
	}
	u.Path = u.Path + dev + deviceCode + control + status
	return u.String()
}

func getSensorDataLink(deviceID string, nodeName string, sensorName string) string {
	return device + deviceID + sensorData + "?" + nodeID + nodeName + "&" + sensorID + sensorName
}

func addIOTLabelLink() string {
	return device + iotLabel
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
