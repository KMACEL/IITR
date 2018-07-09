package profile

import (
	"encoding/json"
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
	//profile = "https://api.ardich.com/api/v3/profile/"
	profile = "profile/"
	push    = "/push/"
	list    = "list"
	name    = "?name="
)

//PushProfileLink is return
func pushProfileLink(setMode string, workingset string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + profile + setMode + push + workingset
	return u.String()
	//return profile + setMode + push + workingset
}

//GetProfileListLink is
func getProfileListLink() string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + profile + list
	return u.String()
	//return profile + list
}

//GetProfileLink is
func getProfileLink(setProfileName string) string {
	data := url.Values{}
	data.Add("name", setProfileName)

	u := rest.GetAPITemplate()
	u.Path = u.Path + profile
	u.RawQuery = data.Encode()
	return u.String()
}

func pushProfileBody(setPolicy string) string {
	var pushProfile PushProfileJSON
	pushProfile.DefaultPolicy.Code = setPolicy
	jsonConvert, _ := json.Marshal(pushProfile)

	return string(jsonConvert)
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
