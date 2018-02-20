package adminarea

// Posted by Mehmet Akasayan

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/
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

type AdminAreaBodyJSON struct {
	Code    string `json:"code"`
	Devices []struct {
		Code string `json:"code"`
	} `json:"devices"`
	Name string `json:"name"`
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}

type CodeJSON struct {
	Code string `json:"code"`
}

type AdminAreaRequirements struct {
	AdminAreaName            string
	AddToAdminAreaDeviceCode []string
}

/*
{
  "code": "",
  "devices": [
    {
      "code": "cea9bbd434b04a7db1865d210f449f0e"
    },
    {
      "code": "84d0aae6300e4a6a81a3e554785b2e54"
    }
  ],
  "name": "test2"
}
*/
