package profile

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
	profile = "https://api.ardich.com/api/v3/profile/"
	push    = "/push/"
	list    = "list"
	name    = "?name="
)

//PushProfileLink is retrun
func pushProfileLink(setMode string, workingset string) string {
	return profile + setMode + push + workingset
}

//GetProfileListLink is
func getProfileListLink() string {
	return profile + list
}

//GetProfileLink is
func getProfileLink(setProfileName string) string {
	return profile + name + setProfileName
}

func pushProfileBody(setPolicy string) string {
	return "{\"defaultPolicy\":{\"code\": \"" + setPolicy + "\"}}"
}
func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}
