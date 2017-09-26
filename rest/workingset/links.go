package workingset

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
	workingset = "https://api.ardich.com/api/v3/workingset/"
	empty      = "empty"
	deviceAdd  = "/devices/add/"
)

//CreateWorkingsetLink is retrun
func CreateWorkingsetLink() string {
	return workingset + empty
}

//AddDeviceWorkingSetLink is
func AddDeviceWorkingSetLink(setWorkingset string) string {
	return workingset + setWorkingset + deviceAdd
}
