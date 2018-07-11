package profile

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

// The Profile package is used to send predefined operations to the device such as applications, configurations, constraints.
const (
	errPushModeQueryTAG     = "profile->profile.go::PushMode->PostQuery"
	errPushModeUnmarshalTAG = "profile->profile.go::PushMode->Unmarshal"

	errPushModeAutoQueryTAG     = "profile->profile.go::PushModeAuto->PostQuery"
	errPushModeAutoUnmarshalTAG = "profile->profile.go::PushModeAuto->Unmarshal"

	errGetProfileListQueryTAG     = "profile->profileinformation.go::GetProfileList->GetQuery"
	errGetProfileListUnmarshalTAG = "profile->profileinformation.go::GetProfileList->Unmarshal"

	errGetProfileQueryTAG     = "profile->profileinformation.go::GetProfile->GetQuery"
	errGetProfileUnmarshalTAG = "profile->profileinformation.go::GetProfile->Unmarshal"
)

/*
     ██╗███████╗ ██████╗ ███╗   ██╗
     ██║██╔════╝██╔═══██╗████╗  ██║
     ██║███████╗██║   ██║██╔██╗ ██║
██   ██║╚════██║██║   ██║██║╚██╗██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
*/

// ModeResponseJSON profile is used for query result returning data
type ModeResponseJSON []struct {
	DeviceID []string `json:"deviceId"`
	Result   string   `json:"result"`
	Status   string   `json:"status"`
	Ok       bool     `json:"ok"`
}

// PushProfileJSON is used for the body part in the profile submission query
type PushProfileJSON struct {
	DefaultPolicy struct {
		Code string `json:"code"`
	} `json:"defaultPolicy"`
}

// ResponseProfileJSONArray includes the return values when performing the profile query
type ResponseProfileJSONArray []struct {
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

// ResponseProfileJSON includes the return a values when performing the profile query
type ResponseProfileJSON struct {
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

// ResponseProfileListJSON gives a list of all existing profiles
type ResponseProfileListJSON struct {
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
