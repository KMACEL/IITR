package application

import (
	"encoding/json"
	"fmt"
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

const (
	application = "application"
	uninstall   = "/uninstall"
)

// https://api.ardich.com/api/v3/application
func getAllApplicationsLink() string {
	data := url.Values{}

	u := rest.GetAPITemplate()
	u.Path = u.Path + application
	u.RawQuery = data.Encode()

	return u.String()
}

// https://api.ardich.com/api/v3/application?packageName={PACKAGE_NAME}
func getApplicationsLink(packageName string) string {
	data := url.Values{}
	data.Add("packageName", packageName)

	u := rest.GetAPITemplate()
	u.Path = u.Path + application
	u.RawQuery = data.Encode()

	return u.String()
}

// https://api.ardich.com:443/api/v3/application/uninstall
func uninstallLink() string {
	data := url.Values{}

	u := rest.GetAPITemplate()
	u.Path = u.Path + application + uninstall
	u.RawQuery = data.Encode()

	return u.String()
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}

/*
██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██████╔╝██║   ██║██║  ██║ ╚████╔╝
██╔══██╗██║   ██║██║  ██║  ╚██╔╝
██████╔╝╚██████╔╝██████╔╝   ██║
╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

//{"apps":[{"packageName": "YOUR_APPLICATION_PACKAGE_NAME"},{"packageName": "YOUR_APPLICATION_PACKAGE_NAME"}],"devices": [{"code": "YOUR_DEVICE_ID"}],"notifyUser": false}
func uninstallBody(notifyUser bool, packageNames []string, devicesID []string) string {
	var uninstallApplicationJSONVar UninstallApplicationJSON

	for _, packageName := range packageNames {
		uninstallApplicationJSONVar.Apps = append(uninstallApplicationJSONVar.Apps, packageNameJSON{PackageName: packageName})
	}

	for _, device := range devicesID {
		uninstallApplicationJSONVar.Devices = append(uninstallApplicationJSONVar.Devices, deviceJSON{Code: device})
	}

	uninstallApplicationJSONVar.NotifyUser = notifyUser
	jsonConvert, _ := json.Marshal(uninstallApplicationJSONVar)
	fmt.Println(string(jsonConvert))
	return string(jsonConvert)
}
