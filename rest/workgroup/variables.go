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
	Name string `json:"name"`
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

/*
{
  "links": [
    {
      "rel": "self",
      "href": "http://api.ardich.com/api/v3/devicegroup{?page,size,sort}"
    }
  ],
  "content": [
    {
      "code": "820105324fc944f7815eb66012b59a47",
      "name": "NFC_20180123_1218",
      "desc": "NFC_20180123_1218",
      "deviceCount": 5,
      "links": [
        {
          "rel": "self",
          "href": "/devicegroup/820105324fc944f7815eb66012b59a47"
        }
      ],
      "createdDate": "2018-01-23T12:18:45.986+03:00"
    }
  ],
  "page": {
    "size": 10,
    "totalElements": 1,
    "totalPages": 1,
    "number": 0
  }
}
*/

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

/*
[
  {
    "deviceId": "867377020892307",
    "imei": "867377020892307",
    "status": "VALID",
    "osVersion": "4.4.4",
    "model": "tp2",
    "modeAppVersion": "AR.AMP.r2.1.340",
    "lockStatus": false,
    "mandatoryLockStatus": false,
    "lostStatus": false,
    "createdDate": 1498723486894,
    "lastModifiedDate": 1516697580320,
    "detailLastModifiedDate": 1516693837690,
    "lastPresenceDate": 1516697579986,
    "presence": {
      "state": "OFFLINE",
      "clientIp": "176.54.117.207"
    },
    "location": {
      "longitude": "28.929343333333335",
      "latitude": "41.01670166666666",
      "provider": "gps",
      "userCreatedDate": 1515674731376,
      "links": []
    },
    "battery": {
      "scale": "100",
      "level": "85",
      "source": "1",
      "voltage": "0mV"
    },
    "network": {
      "telephony": {
        "networkRoaming": false,
        "simOperator": "28602",
        "networkOperatorName": "vodafone TR",
        "msisdn": "",
        "simState": "SIM_STATE_READY",
        "simserialNumber": null
      },
      "wifi": {
        "leaseDuration": "0.0.0.0",
        "mtu": "-1",
        "dns1": "0.0.0.0",
        "dns2": "0.0.0.0",
        "networkType": "mobile",
        "currentWifiApnSsid": null,
        "currentWifiApnHiddenSsid": false,
        "gateway": "0.0.0.0",
        "server": "0.0.0.0",
        "netmask": "0.0.0.0",
        "ip": "192.168.50.233"
      },
      "bluetooth": {
        "bluetoothState": "STATE_UNKNOWN",
        "bluetoothMacId": null,
        "bluetoothSupported": false,
        "bluetoothpaireddevices": null
      }
    },
    "storage": {
      "availIntMemSize": "4,62 GB",
      "totalExtMemSize": "5,5 GB",
      "totalIntMemSize": "5,5 GB",
      "isExtMemAvail": true,
      "availExtMemSize": "4,62 GB"
    },
    "osProfile": {
      "hardware": "rk30board",
      "host": "streamax-PowerEdge-R620",
      "display": "rkpx2-eng 4.4.4 KTU84Q eng.turkey.20171027.143708 test-keys",
      "product": "rkpx2",
      "board": "rk30sdk",
      "model": "tp2",
      "device": "4.4.4",
      "serial": "I0MX8I5MA0"
    },
    "currentUser": {
      "code": "e187d7ead08b4c4a8e63819efd8da60c",
      "mail": "itaksi_user",
      "firstName": "iTaksi",
      "lastName": "User",
      "activationCode": "512013",
      "identityNo": "",
      "enabled": true,
      "activated": true,
      "activationDate": 1516659341529,
      "profile": {
        "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
        "name": null,
        "policy": {
          "code": "CC5750FC-8860-4D50-9290-6A0011448700",
          "name": "iTaksiTrust"
        }
      },
      "current": false
    },
    "users": [
      {
        "code": "e187d7ead08b4c4a8e63819efd8da60c",
        "mail": "itaksi_user",
        "firstName": "iTaksi",
        "lastName": "User",
        "activationCode": "512013",
        "identityNo": "",
        "enabled": true,
        "activated": true,
        "activationDate": 1516659341529,
        "profile": {
          "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
          "name": null,
          "policy": {
            "code": "CC5750FC-8860-4D50-9290-6A0011448700",
            "name": "iTaksiTrust"
          }
        },
        "current": true
      }
    ],
    "groups": [
      {
        "code": "820105324fc944f7815eb66012b59a47",
        "name": "NFC_20180123_1218"
      },
      {
        "code": "cd51dfe72e114ca9bd351a8cabc5f25a",
        "name": "prod release"
      }
    ],
    "adminArea": {
      "code": "f72098ec835044c19db2061302c4a103",
      "name": "DEFAULT_ADMIN_AREA"
    },
    "afexMode": "afex",
    "deviceTimezone": "Doğu Avrupa Standart Saati",
    "deviceCurrentTime": "23 01 2018 09:50",
    "links": [
      {
        "rel": "self",
        "href": "/device/fd0337aab1f64d13a9a4ea6f7cfb7a3f"
      },
      {
        "rel": "Users",
        "href": "/device/fd0337aab1f64d13a9a4ea6f7cfb7a3f/enduser"
      }
    ],
    "code": "fd0337aab1f64d13a9a4ea6f7cfb7a3f"
  },
  {
    "deviceId": "867377020903914",
    "imei": "867377020903914",
    "status": "VALID",
    "osVersion": "4.4.4",
    "model": "tp2",
    "modeAppVersion": "AR.AMP.r2.1.332",
    "lockStatus": false,
    "mandatoryLockStatus": false,
    "lostStatus": false,
    "createdDate": 1498723597021,
    "lastModifiedDate": 1516544221503,
    "detailLastModifiedDate": 1516509058034,
    "lastPresenceDate": 1516544221249,
    "presence": {
      "state": "OFFLINE",
      "clientIp": "176.55.29.42"
    },
    "location": {
      "longitude": "29.15122666666667",
      "latitude": "41.015854999999995",
      "provider": "gps",
      "userCreatedDate": 1516543215714,
      "links": []
    },
    "battery": {
      "scale": "100",
      "level": "85",
      "source": "1",
      "voltage": "0mV"
    },
    "network": {
      "telephony": {
        "networkRoaming": false,
        "simOperator": "28602",
        "networkOperatorName": "vodafone TR",
        "msisdn": "",
        "simState": "SIM_STATE_READY",
        "simserialNumber": null
      },
      "wifi": {
        "leaseDuration": "0.0.0.0",
        "mtu": "-1",
        "dns1": "0.0.0.0",
        "dns2": "0.0.0.0",
        "networkType": "mobile",
        "currentWifiApnSsid": null,
        "currentWifiApnHiddenSsid": false,
        "gateway": "0.0.0.0",
        "server": "0.0.0.0",
        "netmask": "0.0.0.0",
        "ip": "192.168.50.233"
      },
      "bluetooth": {
        "bluetoothState": "STATE_UNKNOWN",
        "bluetoothMacId": null,
        "bluetoothSupported": false,
        "bluetoothpaireddevices": null
      }
    },
    "storage": {
      "availIntMemSize": "5,02 GB",
      "totalExtMemSize": "5,5 GB",
      "totalIntMemSize": "5,5 GB",
      "isExtMemAvail": true,
      "availExtMemSize": "5,02 GB"
    },
    "osProfile": {
      "hardware": "rk30board",
      "host": "streamax-PowerEdge-R620",
      "display": "rkpx2-eng 4.4.4 KTU84Q eng.turkey.20171027.143708 test-keys",
      "product": "rkpx2",
      "board": "rk30sdk",
      "model": "tp2",
      "device": "4.4.4",
      "serial": "KM8VMGVS8Q"
    },
    "currentUser": {
      "code": "e187d7ead08b4c4a8e63819efd8da60c",
      "mail": "itaksi_user",
      "firstName": "iTaksi",
      "lastName": "User",
      "activationCode": "512013",
      "identityNo": "",
      "enabled": true,
      "activated": true,
      "activationDate": 1516659341529,
      "profile": {
        "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
        "name": null,
        "policy": {
          "code": "CC5750FC-8860-4D50-9290-6A0011448700",
          "name": "iTaksiTrust"
        }
      },
      "current": false
    },
    "users": [
      {
        "code": "e187d7ead08b4c4a8e63819efd8da60c",
        "mail": "itaksi_user",
        "firstName": "iTaksi",
        "lastName": "User",
        "activationCode": "512013",
        "identityNo": "",
        "enabled": true,
        "activated": true,
        "activationDate": 1516659341529,
        "profile": {
          "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
          "name": null,
          "policy": {
            "code": "CC5750FC-8860-4D50-9290-6A0011448700",
            "name": "iTaksiTrust"
          }
        },
        "current": true
      }
    ],
    "groups": [
      {
        "code": "820105324fc944f7815eb66012b59a47",
        "name": "NFC_20180123_1218"
      },
      {
        "code": "cd51dfe72e114ca9bd351a8cabc5f25a",
        "name": "prod release"
      },
      {
        "code": "00d6ab19ce1e490e8e57f25c7ccb9230",
        "name": "p1p2"
      }
    ],
    "adminArea": {
      "code": "f72098ec835044c19db2061302c4a103",
      "name": "DEFAULT_ADMIN_AREA"
    },
    "afexMode": "afex",
    "deviceTimezone": "Doğu Avrupa Standart Saati",
    "deviceCurrentTime": "5 01 2018 17:26",
    "links": [
      {
        "rel": "self",
        "href": "/device/f9ec6efd27b646648d12707dc7fd0a67"
      },
      {
        "rel": "Users",
        "href": "/device/f9ec6efd27b646648d12707dc7fd0a67/enduser"
      }
    ],
    "code": "f9ec6efd27b646648d12707dc7fd0a67"
  },
  {
    "deviceId": "867377021066810",
    "imei": "867377021066810",
    "status": "VALID",
    "osVersion": "4.4.4",
    "model": "tp2",
    "modeAppVersion": "AR.AMP.r2.1.340",
    "lockStatus": false,
    "mandatoryLockStatus": false,
    "lostStatus": false,
    "createdDate": 1498728438703,
    "lastModifiedDate": 1516651676557,
    "detailLastModifiedDate": 1516649766384,
    "lastPresenceDate": 1516651676402,
    "presence": {
      "state": "OFFLINE",
      "clientIp": "176.55.153.31"
    },
    "location": {
      "longitude": "28.97238333333333",
      "latitude": "41.00180666666667",
      "provider": "gps",
      "userCreatedDate": 1516651224696,
      "links": []
    },
    "battery": {
      "scale": "100",
      "level": "85",
      "source": "1",
      "voltage": "0mV"
    },
    "network": {
      "telephony": {
        "networkRoaming": false,
        "simOperator": "28602",
        "networkOperatorName": "vodafone TR",
        "msisdn": "",
        "simState": "SIM_STATE_READY",
        "simserialNumber": null
      },
      "wifi": {
        "leaseDuration": "0.0.0.0",
        "mtu": "-1",
        "dns1": "0.0.0.0",
        "dns2": "0.0.0.0",
        "networkType": "mobile",
        "currentWifiApnSsid": null,
        "currentWifiApnHiddenSsid": false,
        "gateway": "0.0.0.0",
        "server": "0.0.0.0",
        "netmask": "0.0.0.0",
        "ip": "192.168.50.233"
      },
      "bluetooth": {
        "bluetoothState": "STATE_UNKNOWN",
        "bluetoothMacId": null,
        "bluetoothSupported": false,
        "bluetoothpaireddevices": null
      }
    },
    "storage": {
      "availIntMemSize": "4,66 GB",
      "totalExtMemSize": "5,5 GB",
      "totalIntMemSize": "5,5 GB",
      "isExtMemAvail": true,
      "availExtMemSize": "4,66 GB"
    },
    "osProfile": {
      "hardware": "rk30board",
      "host": "streamax-PowerEdge-R620",
      "display": "rkpx2-eng 4.4.4 KTU84Q eng.turkey.20171027.143708 test-keys",
      "product": "rkpx2",
      "board": "rk30sdk",
      "model": "tp2",
      "device": "4.4.4",
      "serial": "MFFRHT6LZM"
    },
    "currentUser": {
      "code": "e187d7ead08b4c4a8e63819efd8da60c",
      "mail": "itaksi_user",
      "firstName": "iTaksi",
      "lastName": "User",
      "activationCode": "512013",
      "identityNo": "",
      "enabled": true,
      "activated": true,
      "activationDate": 1516659341529,
      "profile": {
        "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
        "name": null,
        "policy": {
          "code": "CC5750FC-8860-4D50-9290-6A0011448700",
          "name": "iTaksiTrust"
        }
      },
      "current": false
    },
    "users": [
      {
        "code": "e187d7ead08b4c4a8e63819efd8da60c",
        "mail": "itaksi_user",
        "firstName": "iTaksi",
        "lastName": "User",
        "activationCode": "512013",
        "identityNo": "",
        "enabled": true,
        "activated": true,
        "activationDate": 1516659341529,
        "profile": {
          "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
          "name": null,
          "policy": {
            "code": "CC5750FC-8860-4D50-9290-6A0011448700",
            "name": "iTaksiTrust"
          }
        },
        "current": true
      }
    ],
    "groups": [
      {
        "code": "820105324fc944f7815eb66012b59a47",
        "name": "NFC_20180123_1218"
      }
    ],
    "adminArea": {
      "code": "f72098ec835044c19db2061302c4a103",
      "name": "DEFAULT_ADMIN_AREA"
    },
    "afexMode": "afex",
    "deviceTimezone": "Doğu Avrupa Standart Saati",
    "deviceCurrentTime": "22 01 2018 21:36",
    "links": [
      {
        "rel": "self",
        "href": "/device/e3f272bbbaba4e978fbc105b9b424747"
      },
      {
        "rel": "Users",
        "href": "/device/e3f272bbbaba4e978fbc105b9b424747/enduser"
      }
    ],
    "code": "e3f272bbbaba4e978fbc105b9b424747"
  },
  {
    "deviceId": "867377021067040",
    "imei": "867377021067040",
    "status": "VALID",
    "osVersion": "4.4.4",
    "model": "tp2",
    "modeAppVersion": "AR.AMP.r2.1.340",
    "lockStatus": false,
    "mandatoryLockStatus": false,
    "lostStatus": false,
    "createdDate": 1498727639005,
    "lastModifiedDate": 1516688407181,
    "detailLastModifiedDate": 1516681394879,
    "lastPresenceDate": 1516688407002,
    "presence": {
      "state": "OFFLINE",
      "clientIp": "176.55.94.23"
    },
    "location": {
      "longitude": "29.023915000000002",
      "latitude": "41.093878333333336",
      "provider": "gps",
      "userCreatedDate": 1516688104767,
      "links": []
    },
    "battery": {
      "scale": "100",
      "level": "85",
      "source": "1",
      "voltage": "0mV"
    },
    "network": {
      "telephony": {
        "networkRoaming": false,
        "simOperator": "28602",
        "networkOperatorName": "vodafone TR",
        "msisdn": "",
        "simState": "SIM_STATE_READY",
        "simserialNumber": null
      },
      "wifi": {
        "leaseDuration": "0.0.0.0",
        "mtu": "-1",
        "dns1": "0.0.0.0",
        "dns2": "0.0.0.0",
        "networkType": "mobile",
        "currentWifiApnSsid": null,
        "currentWifiApnHiddenSsid": false,
        "gateway": "0.0.0.0",
        "server": "0.0.0.0",
        "netmask": "0.0.0.0",
        "ip": "192.168.50.233"
      },
      "bluetooth": {
        "bluetoothState": "STATE_UNKNOWN",
        "bluetoothMacId": null,
        "bluetoothSupported": false,
        "bluetoothpaireddevices": null
      }
    },
    "storage": {
      "availIntMemSize": "5,02 GB",
      "totalExtMemSize": "5,5 GB",
      "totalIntMemSize": "5,5 GB",
      "isExtMemAvail": true,
      "availExtMemSize": "5,02 GB"
    },
    "osProfile": {
      "hardware": "rk30board",
      "host": "streamax-PowerEdge-R620",
      "display": "rkpx2-eng 4.4.4 KTU84Q eng.turkey.20171027.143708 test-keys",
      "product": "rkpx2",
      "board": "rk30sdk",
      "model": "tp2",
      "device": "4.4.4",
      "serial": "KAO7QWRGOI"
    },
    "currentUser": {
      "code": "e187d7ead08b4c4a8e63819efd8da60c",
      "mail": "itaksi_user",
      "firstName": "iTaksi",
      "lastName": "User",
      "activationCode": "512013",
      "identityNo": "",
      "enabled": true,
      "activated": true,
      "activationDate": 1516659341529,
      "profile": {
        "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
        "name": null,
        "policy": {
          "code": "CC5750FC-8860-4D50-9290-6A0011448700",
          "name": "iTaksiTrust"
        }
      },
      "current": false
    },
    "users": [
      {
        "code": "e187d7ead08b4c4a8e63819efd8da60c",
        "mail": "itaksi_user",
        "firstName": "iTaksi",
        "lastName": "User",
        "activationCode": "512013",
        "identityNo": "",
        "enabled": true,
        "activated": true,
        "activationDate": 1516659341529,
        "profile": {
          "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
          "name": null,
          "policy": {
            "code": "CC5750FC-8860-4D50-9290-6A0011448700",
            "name": "iTaksiTrust"
          }
        },
        "current": true
      }
    ],
    "groups": [
      {
        "code": "820105324fc944f7815eb66012b59a47",
        "name": "NFC_20180123_1218"
      }
    ],
    "adminArea": {
      "code": "f72098ec835044c19db2061302c4a103",
      "name": "DEFAULT_ADMIN_AREA"
    },
    "afexMode": "afex",
    "deviceTimezone": "Doğu Avrupa Standart Saati",
    "deviceCurrentTime": "23 01 2018 06:23",
    "links": [
      {
        "rel": "self",
        "href": "/device/31ff8595e09c49bf94cb7d9632fac5c8"
      },
      {
        "rel": "Users",
        "href": "/device/31ff8595e09c49bf94cb7d9632fac5c8/enduser"
      }
    ],
    "code": "31ff8595e09c49bf94cb7d9632fac5c8"
  },
  {
    "deviceId": "867377021406834",
    "imei": "867377021406834",
    "status": "VALID",
    "osVersion": "4.4.4",
    "model": "tp2",
    "label": "34DF2024 (DEMO)",
    "modeAppVersion": "AR.AMP.r2.1.340",
    "lockStatus": false,
    "mandatoryLockStatus": false,
    "lostStatus": false,
    "createdDate": 1511166984457,
    "lastModifiedDate": 1516630520043,
    "detailLastModifiedDate": 1516081650588,
    "lastPresenceDate": 1516630519874,
    "presence": {
      "state": "OFFLINE",
      "clientIp": "176.55.33.162"
    },
    "location": {
      "longitude": "29.016444999999997",
      "latitude": "41.004838333333325",
      "provider": "gps",
      "userCreatedDate": 1516629991693,
      "links": []
    },
    "battery": {
      "scale": "100",
      "level": "85",
      "source": "1",
      "voltage": "0mV"
    },
    "network": {
      "telephony": {
        "networkRoaming": false,
        "simOperator": "28602",
        "networkOperatorName": "vodafone TR",
        "msisdn": "",
        "simState": "SIM_STATE_READY",
        "simserialNumber": null
      },
      "wifi": {
        "leaseDuration": "0.0.0.0",
        "mtu": "-1",
        "dns1": "0.0.0.0",
        "dns2": "0.0.0.0",
        "networkType": "mobile",
        "currentWifiApnSsid": null,
        "currentWifiApnHiddenSsid": false,
        "gateway": "0.0.0.0",
        "server": "0.0.0.0",
        "netmask": "0.0.0.0",
        "ip": "192.168.50.233"
      },
      "bluetooth": {
        "bluetoothState": "STATE_UNKNOWN",
        "bluetoothMacId": null,
        "bluetoothSupported": false,
        "bluetoothpaireddevices": null
      }
    },
    "storage": {
      "availIntMemSize": "5.01 GB",
      "totalExtMemSize": "5.5 GB",
      "totalIntMemSize": "5.5 GB",
      "isExtMemAvail": true,
      "availExtMemSize": "5.01 GB"
    },
    "osProfile": {
      "hardware": "rk30board",
      "host": "streamax-PowerEdge-R620",
      "display": "rkpx2-eng 4.4.4 KTU84Q eng.turkey.20171027.143708 test-keys",
      "product": "rkpx2",
      "board": "rk30sdk",
      "model": "tp2",
      "device": "4.4.4",
      "serial": "Q8I1WIDQMG"
    },
    "currentUser": {
      "code": "e187d7ead08b4c4a8e63819efd8da60c",
      "mail": "itaksi_user",
      "firstName": "iTaksi",
      "lastName": "User",
      "activationCode": "512013",
      "identityNo": "",
      "enabled": true,
      "activated": true,
      "activationDate": 1516659341529,
      "profile": {
        "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
        "name": null,
        "policy": {
          "code": "CC5750FC-8860-4D50-9290-6A0011448700",
          "name": "iTaksiTrust"
        }
      },
      "current": false
    },
    "users": [
      {
        "code": "e187d7ead08b4c4a8e63819efd8da60c",
        "mail": "itaksi_user",
        "firstName": "iTaksi",
        "lastName": "User",
        "activationCode": "512013",
        "identityNo": "",
        "enabled": true,
        "activated": true,
        "activationDate": 1516659341529,
        "profile": {
          "code": "D42C2D8F-0BAB-4602-8B0E-94B8DA9394B3",
          "name": null,
          "policy": {
            "code": "CC5750FC-8860-4D50-9290-6A0011448700",
            "name": "iTaksiTrust"
          }
        },
        "current": true
      }
    ],
    "groups": [
      {
        "code": "820105324fc944f7815eb66012b59a47",
        "name": "NFC_20180123_1218"
      }
    ],
    "adminArea": {
      "code": "f72098ec835044c19db2061302c4a103",
      "name": "DEFAULT_ADMIN_AREA"
    },
    "afexMode": "afex",
    "deviceTimezone": "Eastern European Standard Time",
    "deviceCurrentTime": "1/16/2018 07:47",
    "links": [
      {
        "rel": "self",
        "href": "/device/f975739a11a54ff99b6bcb11973874ba"
      },
      {
        "rel": "Users",
        "href": "/device/f975739a11a54ff99b6bcb11973874ba/enduser"
      }
    ],
    "code": "f975739a11a54ff99b6bcb11973874ba"
  }
]
*/
