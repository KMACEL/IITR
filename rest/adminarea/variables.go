package adminarea

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

// AdminArea is
type AdminArea struct {
}

/*
     ██╗███████╗ ██████╗ ███╗   ██╗███████╗
     ██║██╔════╝██╔═══██╗████╗  ██║██╔════╝
     ██║███████╗██║   ██║██╔██╗ ██║███████╗
██   ██║╚════██║██║   ██║██║╚██╗██║╚════██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║███████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝
*/

// QueryBodyJSON is
type QueryBodyJSON struct {
	Code    string `json:"code"`
	Devices []struct {
		Code string `json:"code"`
	} `json:"devices"`
	Name string `json:"name"`
}

// CodeJSON is
type CodeJSON struct {
	Code string `json:"code"`
}

// QueryRequirements is
type QueryRequirements struct {
	AdminAreaName            string
	AddToAdminAreaDeviceCode []string
}

// GetAllAdminAreaJSON is
type ResponseGetAllAdminAreaJSON []struct {
	Code         string        `json:"code"`
	Name         string        `json:"name"`
	DefaultState bool          `json:"defaultState"`
	Children     []interface{} `json:"children"`
	Devices      []interface{} `json:"devices"`
	Links        []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

// GetAdminAreaJSON is
type ResponseGetAdminAreaJSON struct {
	Code         string        `json:"code"`
	Name         string        `json:"name"`
	DefaultState bool          `json:"defaultState"`
	Children     []interface{} `json:"children"`
	Devices      []interface{} `json:"devices"`
	Links        []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}
