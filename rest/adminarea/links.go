package adminarea

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

const (
	adminarea    = "adminarea"
	setadminarea = "device/setadminarea"
)

func moveAdminAreaLink() string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + setadminarea
	return u.String()
}

func getAllAdminAreaLink() string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + adminarea
	return u.String()
}

func getAdminAreaLink(getAdminAreaName string) string {
	data := url.Values{}
	data.Add("name", getAdminAreaName)

	u := rest.GetAPITemplate()
	u.Path = u.Path + adminarea
	u.RawQuery = data.Encode()

	return u.String()
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
