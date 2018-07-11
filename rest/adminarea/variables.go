package adminarea

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

// AdminArea is used to group devices, provide control, and limit user access.
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

// QueryBodyJSON is contains the parameters required to perform the query.
type QueryBodyJSON struct {
	Code    string `json:"code"`
	Devices []struct {
		Code string `json:"code"`
	} `json:"devices"`
	Name string `json:"name"`
}

// CodeJSON was created to assist in passing arrays to the Devices field in QueryBodyJSON.
type CodeJSON struct {
	Code string `json:"code"`
}

// QueryRequirements specifies the parameters that the functions receive when querying
type QueryRequirements struct {
	AdminAreaName            string
	AddToAdminAreaDeviceCode []string
}

// ResponseGetAllAdminAreaJSON query result is used to parse and use return values.
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

// ResponseGetAdminAreaJSON query result is used to parse and use return values.
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
