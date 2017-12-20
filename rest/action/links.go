package action

import "strconv"

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
	action     = "https://api.ardich.com/api/v3/action"
	sort       = "?sort=sentDate,desc&size="
	deviceCode = "&deviceCode="
	command    = "&command="
)

var (
	getActionStatusTag          = "Action Status"
	getActionStatusUnMarshalTag = "Action Status Unmarshal"
)

//GetActionStatusLink is
func GetActionStatusLink(setDeviceCode string, setControlType string, setSize int) string {
	return action + sort + strconv.Itoa(setSize) + deviceCode + setDeviceCode + command + setControlType
}
