package drom

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
	apiV3                   = "https://api.ardich.com:443/api/v3/"
	dromdeviceconfiguration = "dromdeviceconfiguration/"
	push                    = "push/"
)

func sendDromLink(setDeviceID string) string {
	return apiV3 + dromdeviceconfiguration + push + setDeviceID
}
