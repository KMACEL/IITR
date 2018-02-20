package adminarea

// Posted by Mehmet Akasayan

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
