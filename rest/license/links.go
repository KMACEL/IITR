package license

import "github.com/KMACEL/IITR/rest"

/*
██╗     ██╗███╗   ██╗██╗  ██╗███████╗
██║     ██║████╗  ██║██║ ██╔╝██╔════╝
██║     ██║██╔██╗ ██║█████╔╝ ███████╗
██║     ██║██║╚██╗██║██╔═██╗ ╚════██║
███████╗██║██║ ╚████║██║  ██╗███████║
╚══════╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
*/

const (
	license  = "license/"
	unassign = "unassign/"
)

func clearLicenseLink(setDeviceID string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + license + unassign + setDeviceID
	return u.String()
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
