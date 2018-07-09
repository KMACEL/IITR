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
)

func sendDromLink(setDeviceID string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + dromdeviceconfiguration + push + setDeviceID

	return u.String()
}
