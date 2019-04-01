package workgroup

import (
	"time"
)

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

/*
     ██╗███████╗ ██████╗ ███╗   ██╗
     ██║██╔════╝██╔═══██╗████╗  ██║
     ██║███████╗██║   ██║██╔██╗ ██║
██   ██║╚════██║██║   ██║██║╚██╗██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
*/

//WorkGroupBodyJSON is
type WorkGroupBodyJSON struct {
	Devices []struct {
		Code string `json:"code"`
	} `json:"devices"`
	Name       string   `json:"name"`
	GroupCodes []string `json:"groupCodes"`
}

func contentTypeJSON() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}

//CodeJSON is
type CodeJSON struct {
	Code string `json:"code"`
}

// WorkGroupRequirements is
type WorkGroupRequirements struct {
	WorkGroupName            string
	AddToWorkGroupDeviceCode []string
	GroupCodes               []string
}

/*
Post : https://api.ardich.com/api/v3/devicegroup/extend
Body :
{
  "devices": [
    {
      "code": "cea9bbd434b04a7db1865d210f449f0e"
    },
    {
      "code": "84d0aae6300e4a6a81a3e554785b2e54"
    }
  ],
  "name": "test2",
  "groupCodes": {}
}
*/

var (
	getWorkGroupIDJSONVariable  GetWorkGroupIDJSON
	getGroupDevicesJSONVariable GetGroupDevicesJSON
)

type WorkGroup struct {
}

type GetWorkGroupIDJSON struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Content []struct {
		Code        string `json:"code"`
		Name        string `json:"name"`
		Desc        string `json:"desc"`
		DeviceCount int    `json:"deviceCount"`
		Links       []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		CreatedDate time.Time `json:"createdDate"`
	} `json:"content"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type GetGroupDevicesJSON []struct {
	DeviceID               string `json:"deviceId"`
	Imei                   string `json:"imei"`
	Status                 string `json:"status"`
	OsVersion              string `json:"osVersion"`
	Model                  string `json:"model"`
	ModeAppVersion         string `json:"modeAppVersion"`
	LockStatus             bool   `json:"lockStatus"`
	MandatoryLockStatus    bool   `json:"mandatoryLockStatus"`
	LostStatus             bool   `json:"lostStatus"`
	CreatedDate            int64  `json:"createdDate"`
	LastModifiedDate       int64  `json:"lastModifiedDate"`
	DetailLastModifiedDate int64  `json:"detailLastModifiedDate"`
	LastPresenceDate       int64  `json:"lastPresenceDate"`
	Presence               struct {
		State    string `json:"state"`
		ClientIP string `json:"clientIp"`
	} `json:"presence"`
	Location struct {
		Longitude       string        `json:"longitude"`
		Latitude        string        `json:"latitude"`
		Provider        string        `json:"provider"`
		UserCreatedDate int64         `json:"userCreatedDate"`
		Links           []interface{} `json:"links"`
	} `json:"location"`
	Battery struct {
		Scale   string `json:"scale"`
		Level   string `json:"level"`
		Source  string `json:"source"`
		Voltage string `json:"voltage"`
	} `json:"battery"`
	Network struct {
		Telephony struct {
			NetworkRoaming      bool        `json:"networkRoaming"`
			SimOperator         string      `json:"simOperator"`
			NetworkOperatorName string      `json:"networkOperatorName"`
			Msisdn              string      `json:"msisdn"`
			SimState            string      `json:"simState"`
			SimserialNumber     interface{} `json:"simserialNumber"`
		} `json:"telephony"`
		Wifi struct {
			LeaseDuration            string      `json:"leaseDuration"`
			Mtu                      string      `json:"mtu"`
			DNS1                     string      `json:"dns1"`
			DNS2                     string      `json:"dns2"`
			NetworkType              string      `json:"networkType"`
			CurrentWifiApnSsid       interface{} `json:"currentWifiApnSsid"`
			CurrentWifiApnHiddenSsid bool        `json:"currentWifiApnHiddenSsid"`
			Gateway                  string      `json:"gateway"`
			Server                   string      `json:"server"`
			Netmask                  string      `json:"netmask"`
			IP                       string      `json:"ip"`
		} `json:"wifi"`
		Bluetooth struct {
			BluetoothState         string      `json:"bluetoothState"`
			BluetoothMacID         interface{} `json:"bluetoothMacId"`
			BluetoothSupported     bool        `json:"bluetoothSupported"`
			Bluetoothpaireddevices interface{} `json:"bluetoothpaireddevices"`
		} `json:"bluetooth"`
	} `json:"network"`
	Storage struct {
		AvailIntMemSize string `json:"availIntMemSize"`
		TotalExtMemSize string `json:"totalExtMemSize"`
		TotalIntMemSize string `json:"totalIntMemSize"`
		IsExtMemAvail   bool   `json:"isExtMemAvail"`
		AvailExtMemSize string `json:"availExtMemSize"`
	} `json:"storage"`
	OsProfile struct {
		Hardware string `json:"hardware"`
		Host     string `json:"host"`
		Display  string `json:"display"`
		Product  string `json:"product"`
		Board    string `json:"board"`
		Model    string `json:"model"`
		Device   string `json:"device"`
		Serial   string `json:"serial"`
	} `json:"osProfile"`
	CurrentUser struct {
		Code           string `json:"code"`
		Mail           string `json:"mail"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		ActivationCode string `json:"activationCode"`
		IdentityNo     string `json:"identityNo"`
		Enabled        bool   `json:"enabled"`
		Activated      bool   `json:"activated"`
		ActivationDate int64  `json:"activationDate"`
		Profile        struct {
			Code   string      `json:"code"`
			Name   interface{} `json:"name"`
			Policy struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"policy"`
		} `json:"profile"`
		Current bool `json:"current"`
	} `json:"currentUser"`
	Users []struct {
		Code           string `json:"code"`
		Mail           string `json:"mail"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		ActivationCode string `json:"activationCode"`
		IdentityNo     string `json:"identityNo"`
		Enabled        bool   `json:"enabled"`
		Activated      bool   `json:"activated"`
		ActivationDate int64  `json:"activationDate"`
		Profile        struct {
			Code   string      `json:"code"`
			Name   interface{} `json:"name"`
			Policy struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"policy"`
		} `json:"profile"`
		Current bool `json:"current"`
	} `json:"users"`
	Groups []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"groups"`
	AdminArea struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"adminArea"`
	AfexMode          string `json:"afexMode"`
	DeviceTimezone    string `json:"deviceTimezone"`
	DeviceCurrentTime string `json:"deviceCurrentTime"`
	Links             []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Code  string `json:"code"`
	Label string `json:"label,omitempty"`
}
