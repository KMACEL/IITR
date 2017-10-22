package workingset

import "github.com/KMACEL/IITR/rest"

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

//Workingset is
type Workingset struct{}

var (
	workingSetKey string
	queryVariable rest.Query
	err           error
	respBody      []byte
)

var (
	workingsetJSONVariable        DWorkingsetJSON
	workingsetDevicesJSONVariable WorkingsetDevicesJSON
)

var (
	responsePushApplicationJSONVariable responsePushApplicationJSON
)

/*
     ██╗███████╗ ██████╗ ███╗   ██╗
     ██║██╔════╝██╔═══██╗████╗  ██║
     ██║███████╗██║   ██║██╔██╗ ██║
██   ██║╚════██║██║   ██║██║╚██╗██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
*/

// DWorkingsetJSON is
type DWorkingsetJSON struct {
	NoAllowLostDevices interface{}   `json:"noAllowLostDevices"`
	Code               string        `json:"code"`
	DeviceCount        int           `json:"deviceCount"`
	Devices            []interface{} `json:"devices"`
	Links              []interface{} `json:"links"`
	CreatedDate        int64         `json:"createdDate"`
}

// WorkingsetDevicesJSON is
type WorkingsetDevicesJSON struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Content []struct {
		DeviceID    string `json:"deviceId"`
		Status      string `json:"status"`
		Model       string `json:"model"`
		LockStatus  bool   `json:"lockStatus"`
		LostStatus  bool   `json:"lostStatus"`
		CreatedDate int64  `json:"createdDate"`
		Presence struct {
			State    string      `json:"state"`
			ClientIP interface{} `json:"clientIp"`
		} `json:"presence"`
		Network struct {
			Telephony struct {
				NetworkRoaming      bool        `json:"networkRoaming"`
				SimOperator         interface{} `json:"simOperator"`
				NetworkOperatorName interface{} `json:"networkOperatorName"`
				Msisdn              interface{} `json:"msisdn"`
				SimState            interface{} `json:"simState"`
				SimserialNumber     interface{} `json:"simserialNumber"`
			} `json:"telephony"`
			Wifi struct {
				LeaseDuration            interface{} `json:"leaseDuration"`
				Mtu                      interface{} `json:"mtu"`
				DNS1                     interface{} `json:"dns1"`
				DNS2                     interface{} `json:"dns2"`
				NetworkType              interface{} `json:"networkType"`
				CurrentWifiApnSsid       interface{} `json:"currentWifiApnSsid"`
				CurrentWifiApnHiddenSsid bool        `json:"currentWifiApnHiddenSsid"`
				Gateway                  interface{} `json:"gateway"`
				Server                   interface{} `json:"server"`
				Netmask                  interface{} `json:"netmask"`
				IP                       interface{} `json:"ip"`
			} `json:"wifi"`
			Bluetooth interface{} `json:"bluetooth"`
		} `json:"network"`
		OsProfile struct {
			Hardware interface{} `json:"hardware"`
			Host     interface{} `json:"host"`
			Display  interface{} `json:"display"`
			Product  interface{} `json:"product"`
			Board    interface{} `json:"board"`
			Model    interface{} `json:"model"`
			Device   interface{} `json:"device"`
			Serial   interface{} `json:"serial"`
		} `json:"osProfile"`
		CurrentUser struct {
			Code           interface{} `json:"code"`
			Mail           string      `json:"mail"`
			FirstName      string      `json:"firstName"`
			LastName       string      `json:"lastName"`
			ActivationCode interface{} `json:"activationCode"`
			IdentityNo     interface{} `json:"identityNo"`
			Enabled        bool        `json:"enabled"`
			Activated      bool        `json:"activated"`
			ActivationDate interface{} `json:"activationDate"`
			Profile struct {
				Code string `json:"code"`
				Name string `json:"name"`
				Policy struct {
					Code string      `json:"code"`
					Name interface{} `json:"name"`
				} `json:"policy"`
			} `json:"profile"`
			Current bool `json:"current"`
		} `json:"currentUser"`
		Groups []struct {
			Code interface{} `json:"code"`
			Name interface{} `json:"name"`
		} `json:"groups"`
		AdminArea struct {
			Code interface{} `json:"code"`
			Name string      `json:"name"`
		} `json:"adminArea"`
		ActivePolicy struct {
			Code interface{} `json:"code"`
			Name interface{} `json:"name"`
		} `json:"activePolicy"`
		AfexMode string `json:"afexMode"`
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		Code  string `json:"code"`
		Label string `json:"label,omitempty"`
	} `json:"content"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type responsePushApplicationJSON struct {
	Response string        `json:"response"`
	Links    []interface{} `json:"links"`
}
