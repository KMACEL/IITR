package drom

import (
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

var (
	dromdeviceconfiguration = "dromdeviceconfiguration/"
	push                    = "push/"
	delete                  = "delete/"
)

// https://api.ardich.com/api/v3/dromdeviceconfiguration/push/{YOUR_DEVICE_ID}
func sendDromLink(setDeviceID string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + dromdeviceconfiguration + push + setDeviceID

	return u.String()
}

// https://api.ardich.com/api/v3/dromdeviceconfiguration/delete/{YOUR_DEVICE_ID}
func deleteDromLink(setDeviceID string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + dromdeviceconfiguration + delete + setDeviceID

	return u.String()
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
