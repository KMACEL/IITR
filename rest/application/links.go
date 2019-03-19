package application

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
	application = "application"
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

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
