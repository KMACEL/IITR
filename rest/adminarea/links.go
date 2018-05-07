package adminarea

// Posted by Mehmet Aksayan

/*
██╗     ██╗███╗   ██╗██╗  ██╗███████╗
██║     ██║████╗  ██║██║ ██╔╝██╔════╝
██║     ██║██╔██╗ ██║█████╔╝ ███████╗
██║     ██║██║╚██╗██║██╔═██╗ ╚════██║
███████╗██║██║ ╚████║██║  ██╗███████║
╚══════╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
*/

const (
	setadminarea = "https://api.ardich.com/api/v3/device/setadminarea"
)

func moveAdminAreaLink() string {
	return setadminarea
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
