package device

import (
	"time"

	"github.com/KMACEL/IITR/rest"
)

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

//Device is
type Device struct {
}

const (
	getDownloadApplicationListErrorTag = "Get Download Application List Code :"
	getBuiltInApplicationListErrorTag  = "Get Built In Application List"
	removeApplicationErrorTag          = "Remove Application"
	presenceInfoErrorTag               = "Presence Info"
)

var (
	q                                     rest.Query
	responseMessageCodeJSONVariable       ResponseMessageCodeJSON
	responseMessageJSONVariable           ResponseMessageJSON
	responseDescriptionJSONVariable       ResponseDescriptionJSON
	responseMessageErrorJSONVariable      ResponseMesageErrorJSON
	presenceInfoJSONVariable              PresenceInfoJSON
	locationJSONVariable                  LocationJSON
	downloadedApplicationListJSONVariable DownloadedApplicationListJSON
	builtInApplicationListJSONVariable    BuiltInApplicationListJSON
	activeProfilePolicyJSONVariable       ActiveProfilePolicyJSON
	locationAllJSONVariable               LocationAllJSON
	logListJSONVariable                   LogListJSON
	deviceInformationJSONVariable         InformationJSON
	summaryJSONVariable                   SummaryJSON
	applicationInfoJSONVariable           ApplicationInfoJSON
	osProfileInfoJSONVariable             OSProfileInfoJSON
	instantApplicationInfoJSONVariable    InstantApplicationInfoJSON
	presenceHistroyJSONVariable           PresenceHistroyJSON
	sensorDataJSONVariable                SensorDataJSON
)

const (
	//StartApp Start Parameter
	StartApp = 1
	//StopApp Stop Parameter
	StopApp = 0
)

// Rest -> Device Constant
const (
	summaryTag = "Summary : "
)

// This Const is
const (
	OSProfile             = "OSProfile"
	BatteryInfo           = "BatteryInfo"
	ModiverseInfo         = "ModiverseInfo"
	NetworkInfo           = "NetworkInfo"
	RootedInfo            = "RootedInfo"
	ProcessInfo           = "ProcessInfo"
	StorageInfo           = "StorageInfo"
	UsageInfo             = "UsageInfo"
	ApplicationInfo       = "ApplicationInfo"
	LocationInfo          = "LocationInfo"
	DeviceNodeInventory   = "DeviceNodeInventory"
	DeviceFlowInventory   = "DeviceFlowInventory"
	DeviceConfigInventory = "DeviceConfigInventory"
)

/*
     ██╗███████╗ ██████╗ ███╗   ██╗███████╗
     ██║██╔════╝██╔═══██╗████╗  ██║██╔════╝
     ██║███████╗██║   ██║██╔██╗ ██║███████╗
██   ██║╚════██║██║   ██║██║╚██╗██║╚════██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║███████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝
*/

// ResponseMessageCodeJSON is
type ResponseMessageCodeJSON struct {
	Response string        `json:"response"`
	Links    []interface{} `json:"links"`
}

// ResponseDescriptionJSON is
type ResponseDescriptionJSON struct {
	State string `json:"state"`
	Code  int    `json:"code"`
}

// ResponseMesageErrorJSON is
type ResponseMesageErrorJSON struct {
	ERROR struct {
		Message string `json:"message"`
	} `json:"ERROR"`
	State       string `json:"state"`
	Description string `json:"description"`
	Code        int    `json:"code"`
}

// ResponseMessageJSON is
type ResponseMessageJSON struct {
	Code               string        `json:"code"`
	DeviceID           string        `json:"deviceId"`
	DeviceSerial       string        `json:"deviceSerial"`
	Command            string        `json:"command"`
	DeviceCode         interface{}   `json:"deviceCode"`
	Parameters         string        `json:"parameters"`
	Label              interface{}   `json:"label"`
	ResponseID         string        `json:"responseId"`
	StartDate          interface{}   `json:"startDate"`
	EndDate            interface{}   `json:"endDate"`
	SentStatus         bool          `json:"sentStatus"`
	ErrorMessage       interface{}   `json:"errorMessage"`
	CreatedBy          interface{}   `json:"createdBy"`
	UserMail           interface{}   `json:"userMail"`
	Links              []interface{} `json:"links"`
	SendToDeviceStatus struct {
		Status   string    `json:"status"`
		TryCount string    `json:"tryCount"`
		Body     string    `json:"body"`
		SentDate time.Time `json:"sentDate"`
	} `json:"sendToDeviceStatus"`
	DeliveryFromDeviceStatus struct {
		Desc          string    `json:"desc"`
		DeliveredDate time.Time `json:"deliveredDate"`
	} `json:"deliveryFromDeviceStatus"`
	ExecutionInDeviceStatus struct {
		Result string `json:"result"`
	} `json:"executionInDeviceStatus"`
	SentDate time.Time `json:"sentDate"`
}

// PresenceInfoJSON is
type PresenceInfoJSON struct {
	DeviceID string `json:"deviceId"`
	Command  string `json:"command"`
	Data     struct {
		State    string `json:"state"`
		ClientIP string `json:"clientIp"`
	} `json:"data"`
	CreateDate int64  `json:"createDate"`
	NodeID     string `json:"nodeId"`
	SensorID   string `json:"sensorId"`
	CloudDate  int    `json:"cloudDate"`
}

//ApplicationInfoJSON is
type ApplicationInfoJSON struct {
	DeviceID string `json:"deviceId"`
	Command  string `json:"command"`
	Data     []struct {
		DataSize       float64 `json:"dataSize"`
		PackageName    string  `json:"packageName"`
		VersionCode    float64 `json:"versionCode"`
		VersionName    string  `json:"versionName"`
		UsbStoreSize   float64 `json:"usbStoreSize"`
		Running        bool    `json:"running"`
		Size           float64 `json:"size"`
		SdcardSize     float64 `json:"sdcardSize"`
		Builtin        float64 `json:"builtin"`
		Name           string  `json:"name"`
		Blocked        float64 `json:"blocked"`
		UUID           float64 `json:"uuid"`
		TotalSize      float64 `json:"totalSize"`
		UpdatedBuiltin float64 `json:"updatedBuiltin,omitempty"`
		CacheSize      float64 `json:"cacheSize"`
	} `json:"data"`
	CreateDate int64  `json:"createDate"`
	NodeID     string `json:"nodeId"`
	SensorID   string `json:"sensorId"`
	CloudDate  int    `json:"cloudDate"`
}

// OSProfileInfoJSON is
type OSProfileInfoJSON struct {
	DeviceID string `json:"deviceId"`
	Command  string `json:"command"`
	Data     struct {
		Model              string  `json:"model"`
		Hardware           string  `json:"hardware"`
		Type               string  `json:"type"`
		ID                 string  `json:"id"`
		Time               int64   `json:"time"`
		ModeApkVersionName string  `json:"modeApkVersionName"`
		CPUAbi2            string  `json:"cpuAbi2"`
		LocalIP            string  `json:"localIp"`
		Tags               string  `json:"tags"`
		DeviceID           string  `json:"DeviceId"`
		ModeApkVersionCode float64 `json:"modeApkVersionCode"`
		Host               string  `json:"host"`
		OsVersion          string  `json:"osVersion"`
		Display            string  `json:"display"`
		Board              string  `json:"board"`
		Product            string  `json:"product"`
		Manufacturer       string  `json:"manufacturer"`
		BootLoader         string  `json:"bootLoader"`
		CPUAbi             string  `json:"cpuAbi"`
		Device             string  `json:"device"`
		Radio              string  `json:"radio"`
		Brand              string  `json:"brand"`
		User               string  `json:"user"`
		Serial             string  `json:"serial"`
		OsName             string  `json:"osName"`
		AfexMode           string  `json:"afexMode"`
	} `json:"data"`
	CreateDate int64  `json:"createDate"`
	NodeID     string `json:"nodeId"`
	SensorID   string `json:"sensorId"`
	CloudDate  int    `json:"cloudDate"`
}

// InstantApplicationInfoJSON is
type InstantApplicationInfoJSON struct {
	DeviceID string `json:"deviceId"`
	Command  string `json:"command"`
	Data     struct {
		AppName        string  `json:"appName"`
		Currentimezone string  `json:"currentimezone"`
		PackageName    string  `json:"packageName"`
		Currenttime    int64   `json:"currenttime"`
		Action         string  `json:"action"`
		VersionCode    float32 `json:"versionCode"`
		VersionName    string  `json:"versionName"`
	} `json:"data"`
	CreateDate int64  `json:"createDate"`
	NodeID     string `json:"nodeId"`
	SensorID   string `json:"sensorId"`
	CloudDate  int    `json:"cloudDate"`
}

// LocationJSON is
type LocationJSON struct {
	Code        interface{} `json:"code"`
	SubCode     interface{} `json:"subCode"`
	Status      string      `json:"status"`
	Result      interface{} `json:"result"`
	Message     interface{} `json:"message"`
	Exception   interface{} `json:"exception"`
	Description interface{} `json:"description"`
	Success     interface{} `json:"success"`
	Error       interface{} `json:"error"`
	Extras      []struct {
		DeviceID   string `json:"deviceId"`
		DeviceCode string `json:"deviceCode"`
		Longitude  string `json:"longitude"`
		Latitude   string `json:"latitude"`
		Provider   string `json:"provider"`
		InsertDate int64  `json:"insertDate"`
		Presence   string `json:"presence"`
	} `json:"extras"`
	Ok   bool `json:"ok"`
	Sent bool `json:"sent"`
}

// DownloadedApplicationListJSON is
type DownloadedApplicationListJSON []struct {
	PackageName  string        `json:"packageName"`
	Notify       bool          `json:"notify"`
	Name         string        `json:"name"`
	VersionCode  string        `json:"versionCode"`
	VersionName  string        `json:"versionName"`
	Type         string        `json:"type"`
	Size         int           `json:"size"`
	DeviceID     string        `json:"deviceId"`
	TotalSize    int           `json:"totalSize"`
	Running      bool          `json:"running"`
	Blocked      float64       `json:"blocked"`
	SdcardSize   int           `json:"sdcardSize"`
	UsbStoreSize int           `json:"usbStoreSize"`
	UUID         int           `json:"uuid"`
	DataSize     int           `json:"dataSize"`
	CacheSize    int           `json:"cacheSize"`
	AddedDate    int64         `json:"addedDate"`
	Links        []interface{} `json:"links"`
}

// BuiltInApplicationListJSON is
type BuiltInApplicationListJSON []struct {
	PackageName  string        `json:"packageName"`
	Notify       bool          `json:"notify"`
	Name         string        `json:"name"`
	VersionCode  string        `json:"versionCode"`
	VersionName  string        `json:"versionName"`
	Type         string        `json:"type"`
	Size         int           `json:"size"`
	DeviceID     string        `json:"deviceId"`
	TotalSize    int           `json:"totalSize"`
	Running      bool          `json:"running"`
	Blocked      int           `json:"blocked"`
	SdcardSize   int           `json:"sdcardSize"`
	UsbStoreSize int           `json:"usbStoreSize"`
	UUID         int           `json:"uuid"`
	DataSize     int           `json:"dataSize"`
	CacheSize    int           `json:"cacheSize"`
	AddedDate    int64         `json:"addedDate"`
	Links        []interface{} `json:"links"`
}

// ActiveProfilePolicyJSON is
type ActiveProfilePolicyJSON struct {
	ActiveProfile  string `json:"activeProfile"`
	ActivePolicy   string `json:"activePolicy"`
	CurrentProfile string `json:"currentProfile"`
	CurrentPolicy  string `json:"currentPolicy"`
	Device         string `json:"device"`
}

// LocationAllJSON is
type LocationAllJSON []struct {
	DeviceID        string        `json:"deviceId"`
	DeviceCode      string        `json:"deviceCode"`
	User            string        `json:"user"`
	Longitude       string        `json:"longitude"`
	Latitude        string        `json:"latitude"`
	Provider        string        `json:"provider"`
	UserCreatedDate int64         `json:"userCreatedDate"`
	Links           []interface{} `json:"links"`
}

// LogListJSON is
type LogListJSON []struct {
	DeviceID string        `json:"deviceId"`
	Name     string        `json:"name"`
	URL      string        `json:"url"`
	Token    string        `json:"token"`
	Links    []interface{} `json:"links"`
}

// InformationJSON is
type InformationJSON struct {
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
	FirstPresenceDate      int64  `json:"firstPresenceDate"`
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
			Code   string `json:"code"`
			Name   string `json:"name"`
			Policy struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"policy"`
		} `json:"profile"`
		Current bool `json:"current"`
	} `json:"users"`
	AdminArea struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"adminArea"`
	ActivePolicy struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"activePolicy"`
	AfexMode          string `json:"afexMode"`
	DeviceTimezone    string `json:"deviceTimezone"`
	DeviceCurrentTime string `json:"deviceCurrentTime"`
	Links             []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Code string `json:"code"`
}

// SummaryJSON is
type SummaryJSON struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Content []struct {
		DeviceID         string `json:"deviceId"`
		Imei             string `json:"imei"`
		Status           string `json:"status"`
		Model            string `json:"model"`
		ModeAppVersion   string `json:"modeAppVersion"`
		LockStatus       bool   `json:"lockStatus"`
		LostStatus       bool   `json:"lostStatus"`
		CreatedDate      int64  `json:"createdDate"`
		LastPresenceDate int64  `json:"lastPresenceDate"`
		Presence         struct {
			State    string `json:"state"`
			ClientIP string `json:"clientIp"`
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
				IP                       string      `json:"ip"`
			} `json:"wifi"`
			Bluetooth interface{} `json:"bluetooth"`
		} `json:"network"`
		OsProfile struct {
			Hardware interface{} `json:"hardware"`
			Host     interface{} `json:"host"`
			Display  interface{} `json:"display"`
			Product  interface{} `json:"product"`
			Board    interface{} `json:"board"`
			Model    string      `json:"model"`
			Device   interface{} `json:"device"`
			Serial   string      `json:"serial"`
		} `json:"osProfile"`
		CurrentUser struct {
			Code           interface{} `json:"code"`
			Mail           string      `json:"mail"`
			FirstName      string      `json:"firstName"`
			LastName       string      `json:"lastName"`
			ActivationCode interface{} `json:"activationCode"`
			IdentityNo     string      `json:"identityNo"`
			Enabled        bool        `json:"enabled"`
			Activated      bool        `json:"activated"`
			ActivationDate interface{} `json:"activationDate"`
			Profile        struct {
				Code   string `json:"code"`
				Name   string `json:"name"`
				Policy struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"policy"`
			} `json:"profile"`
			Current bool `json:"current"`
		} `json:"currentUser"`
		Groups []struct {
			Code interface{} `json:"code"`
			Name string      `json:"name"`
		} `json:"groups"`
		AdminArea struct {
			Code interface{} `json:"code"`
			Name string      `json:"name"`
		} `json:"adminArea"`
		ActivePolicy struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"activePolicy"`
		AfexMode          string `json:"afexMode"`
		CurrentPolicy     string `json:"currentPolicy"`
		DeviceTimezone    string `json:"deviceTimezone"`
		DeviceCurrentTime string `json:"deviceCurrentTime"`
		Links             []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		Code string `json:"code"`
	} `json:"content"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

//PresenceHistroyJSON is
type PresenceHistroyJSON struct {
	List []struct {
		DeviceID string `json:"deviceId"`
		Command  string `json:"command"`
		Data     struct {
			State    string `json:"state"`
			ClientIP string `json:"clientIp"`
		} `json:"data"`
		CreateDate int64  `json:"createDate"`
		NodeID     string `json:"nodeId"`
		SensorID   string `json:"sensorId"`
		CloudDate  int    `json:"cloudDate"`
	} `json:"list"`
	Count  int    `json:"count"`
	LastID string `json:"lastId"`
}

//SensorDataJSON is
type SensorDataJSON struct {
	Code        interface{} `json:"code"`
	SubCode     interface{} `json:"subCode"`
	Status      string      `json:"status"`
	Description interface{} `json:"description"`
	Data        struct {
		DeviceID   string `json:"deviceId"`
		Command    string `json:"command"`
		Data       string `json:"data"`
		CreateDate int64  `json:"createDate"`
		NodeID     string `json:"nodeId"`
		SensorID   string `json:"sensorId"`
		CloudDate  int64  `json:"cloudDate"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

//AddIOTLabelJSON is
type AddIOTLabelJSON struct {
	DeviceNodeSensors []struct {
		DeviceID string `json:"deviceId"`
		NodeID   string `json:"nodeId"`
		SensorID string `json:"sensorId"`
	} `json:"deviceNodeSensors"`
	Label     string `json:"label"`
	LabelType string `json:"labelType"`
}

//AddIOTLabelDeviceIDJSON is
type AddIOTLabelDeviceIDJSON struct {
	DeviceID string `json:"deviceId"`
	NodeID   string `json:"nodeId"`
	SensorID string `json:"sensorId"`
}

// SetLabelBodyJSON is
type SetLabelBodyJSON struct {
	Label string `json:"label"`
}

// ModiverseInfoJSON is
type ModiverseInfoJSON struct {
	DeviceID string `json:"deviceId"`
	Command  string `json:"command"`
	Data     struct {
		DeviceAdmin        bool   `json:"deviceAdmin"`
		ModeApkVersionCode string `json:"modeApkVersionCode"`
		ModeApkVersionName string `json:"modeApkVersionName"`
		ProfileOwner       bool   `json:"profileOwner"`
		UsageStatsEnabled  bool   `json:"usageStatsEnabled"`
		DeviceOwner        bool   `json:"deviceOwner"`
		AfexMode           string `json:"afexMode"`
	} `json:"data"`
	CreateDate int64  `json:"createDate"`
	NodeID     string `json:"nodeId"`
	SensorID   string `json:"sensorId"`
	CloudDate  int    `json:"cloudDate"`
}
