package profile

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
//This page is the part that shows the links that the queries will use.
//It is designed in such a way that the administration is easy.

const (
	profile  = "profile/"
	push     = "/push/"
	list     = "list"
	policies = "/policies"
)

// https://api.ardich.com/api/v3/profile/{setMode}/push/{workingset}
func pushProfileLink(setMode string, workingset string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + profile + setMode + push + workingset

	return u.String()
}

// https://api.ardich.com/api/v3/profile/list
func getProfileListLink() string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + profile + list
	fmt.Println(u.String())

	return u.String()
}

// https://api.ardich.com/api/v3/profile/?name={YOUR_PROFILE_NAME}}
func getProfileLink(setProfileName string) string {
	data := url.Values{}
	data.Add("name", setProfileName)

	u := rest.GetAPITemplate()
	u.Path = u.Path + profile
	u.RawQuery = data.Encode()

	return u.String()
}

// https://api.ardich.com/api/v3/profile/{PROFILE_CODE}/policies
func getProfileInPolicyListLink(profileCode string) string {
	u := rest.GetAPITemplate()
	u.Path = u.Path + profile + profileCode + policies
	fmt.Println(u.String())

	return u.String()
}

// {"defaultPolicy":{"code":"{setPolicy}"}}
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
