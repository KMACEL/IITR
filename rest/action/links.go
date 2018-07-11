package action

import (
	"net/url"
	"strconv"

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
	action = "action"
)

var (
	getActionStatusTag          = "Action Status"
	getActionStatusUnMarshalTag = "Action Status Unmarshal"
)

// https://api.ardich.com/api/v3/action?command={setControlType}&deviceCode={setDeviceCode}&size={setSize}&sort=sentDate&sort=desc
func getActionStatusLink(setDeviceCode string, setControlType string, setSize int) string {
	data := url.Values{}
	data.Add("sort", "sentDate")
	data.Add("sort", "desc")
	data.Add("size", strconv.Itoa(setSize))
	data.Add("deviceCode", setDeviceCode)
	data.Add("command", setControlType)

	u := rest.GetAPITemplate()
	u.Path = u.Path + action
	u.RawQuery = data.Encode()

	return u.String()
}
