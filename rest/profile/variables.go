package profile

import "github.com/KMACEL/IITR/rest"

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

//Profile is
type Profile struct{}

var (
	queryVariable rest.Query
	err           error
	respBody      []byte
)

var (
	profileInformationJSONVariable ProfileInformationJSON
	modeResponseJSONVariable       ModeResponseJSON
	profileAllListJSONVariable     ProfileAllListJSON
)

/*
     ██╗███████╗ ██████╗ ███╗   ██╗
     ██║██╔════╝██╔═══██╗████╗  ██║
     ██║███████╗██║   ██║██╔██╗ ██║
██   ██║╚════██║██║   ██║██║╚██╗██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
*/

type SetPolicyJSON struct {
	DefaultPolicy struct {
		Code string `json:"code"`
	} `json:"defaultPolicy"`
}

type ModeResponseJSON []struct {
	DeviceID []string `json:"deviceId"`
	Result   string   `json:"result"`
	Status   string   `json:"status"`
	Ok       bool     `json:"ok"`
}

type ProfileInformationJSON []struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	DefaultProfile bool   `json:"defaultProfile"`
	DefaultPolicy  struct {
		Code string      `json:"code"`
		Name interface{} `json:"name"`
	} `json:"defaultPolicy"`
	Configurations struct {
		EmergencySettings interface{} `json:"emergencySettings"`
		LocationSettings  interface{} `json:"locationSettings"`
		BatterySettings   interface{} `json:"batterySettings"`
		DeactivatedMode   interface{} `json:"deactivatedMode"`
		OlaSettings       interface{} `json:"olaSettings"`
	} `json:"configurations"`
	Scheduled                interface{} `json:"scheduled"`
	IsSwitchable             bool        `json:"isSwitchable"`
	IsInContainer            bool        `json:"isInContainer"`
	SwitchPassword           string      `json:"switchPassword"`
	IsBypassGoogleActivation bool        `json:"isBypassGoogleActivation"`
	IsEmpoweredMode          bool        `json:"isEmpoweredMode"`
	ActivationCode           string      `json:"activationCode"`
	Links                    []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

type ProfileAllListJSON struct {
	Code        string      `json:"code"`
	SubCode     interface{} `json:"subCode"`
	Status      string      `json:"status"`
	Result      interface{} `json:"result"`
	Message     interface{} `json:"message"`
	Exception   interface{} `json:"exception"`
	Description interface{} `json:"description"`
	Success     string      `json:"success"`
	Error       interface{} `json:"error"`
	Extras      struct {
		Profiles []struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			DefaultProfile bool   `json:"defaultProfile"`
			DefaultPolicy  string `json:"defaultPolicy"`
			Policies       []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"policies"`
		} `json:"profiles"`
	} `json:"extras"`
	Ok   bool `json:"ok"`
	Sent bool `json:"sent"`
}
