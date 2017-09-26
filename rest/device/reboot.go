package device

/*
██████╗ ███████╗██████╗  ██████╗  ██████╗ ████████╗
██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝
██████╔╝█████╗  ██████╔╝██║   ██║██║   ██║   ██║
██╔══██╗██╔══╝  ██╔══██╗██║   ██║██║   ██║   ██║
██║  ██║███████╗██████╔╝╚██████╔╝╚██████╔╝   ██║
╚═╝  ╚═╝╚══════╝╚═════╝  ╚═════╝  ╚═════╝    ╚═╝
*/

//Reboot is
func (d Device) Reboot(deviceCode string, vasualFlag bool) {
	setQueryAdress := rebootLink(deviceCode)
	queryVariable.PostQuery(setQueryAdress, "", contentTypeJSON(), vasualFlag)
}
