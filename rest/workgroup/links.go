package workgroup

//Posted by Mehmet Akasayan

/*
██╗     ██╗███╗   ██╗██╗  ██╗███████╗
██║     ██║████╗  ██║██║ ██╔╝██╔════╝
██║     ██║██╔██╗ ██║█████╔╝ ███████╗
██║     ██║██║╚██╗██║██╔═██╗ ╚════██║
███████╗██║██║ ╚████║██║  ██╗███████║
╚══════╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
*/

const (
	devicegroup       = "https://api.ardich.com:443/api/v3/devicegroup/"
	api               = "https://api.ardich.com/api/v3/"
	devicegroupSingle = "devicegroup"
	deviceParam       = "/device"
	extendParam       = "extend"
	name              = "?name="
)

func getGroupDevicesLink(workGroupID string) string {
	return devicegroup + workGroupID + deviceParam
}

func moveWorkGroupLink() string {
	return devicegroup + extendParam
}

func getGroupIDLink(workGroupName string) string {
	return api + devicegroupSingle + name + workGroupName
}
