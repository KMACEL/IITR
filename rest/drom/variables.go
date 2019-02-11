package drom

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

/*
██████╗ ██████╗  ██████╗ ███╗   ███╗
██╔══██╗██╔══██╗██╔═══██╗████╗ ████║
██║  ██║██████╔╝██║   ██║██╔████╔██║
██║  ██║██╔══██╗██║   ██║██║╚██╔╝██║
██████╔╝██║  ██║╚██████╔╝██║ ╚═╝ ██║
╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝
*/

// Drom is a method for remotely licensing a device.
// The important thing to note here is that the device
// must have a drom-recording. Otherwise, no licensing will be done.
type Drom struct{}

/*
     ██╗███████╗ ██████╗ ███╗   ██╗
     ██║██╔════╝██╔═══██╗████╗  ██║
     ██║███████╗██║   ██║██╔██╗ ██║
██   ██║╚════██║██║   ██║██║╚██╗██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
*/

// ConfigurationListJSON is
type ConfigurationListJSON []struct {
	Name            string        `json:"name"`
	Configuration   string        `json:"configuration"`
	TenantDomain    string        `json:"tenantDomain"`
	Created         string        `json:"created"`
	Modified        string        `json:"modified"`
	ByDefault       bool          `json:"byDefault"`
	ConfigurationID string        `json:"configurationId"`
	Links           []interface{} `json:"links"`
}

type addDeviceBodyJSON struct {
	ConfigurationID string `json:"configurationId"`
	DeviceID        string `json:"deviceId"`
}
