package action

import "time"

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

//Action is
type Action struct{}

var (
	messageJSONVariable ResponseActionMessageJSON
)

// These constants help to identify future types of action messages
const (
	RingStart             = "ringStart"
	SendSensorAgent       = "sendSensorAgent"
	RingStop              = "ringStop"
	OpenURL               = "openUrl"
	LockDevice            = "lockDevice"
	WakeUpDevice          = "wakeUpDevice"
	SwitchContainer       = "switchContainer"
	StartApp              = "startApp"
	RemoveModeApplication = "removeModeApplication"
	RemoveNodeSensor      = "removeNodeSensor"
	UnLockDevice          = "unLockDevice"
	Wipe                  = "wipe"
	Reboot                = "reboot"
	TakeScreenShot        = "takeScreenShot"
	SendMessage           = "sendMessage"
	LogOff                = "logOff"
	PushDROM              = "PUSH_CMD_DROM"
	SendProductProfile    = "sendProductProfile"
	GetCurrentStatus      = "getCurrentStatus"
)

/*
     ██╗███████╗ ██████╗ ███╗   ██╗
     ██║██╔════╝██╔═══██╗████╗  ██║
     ██║███████╗██║   ██║██╔██╗ ██║
██   ██║╚════██║██║   ██║██║╚██╗██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
*/

// ResponseActionMessageJSON is the JSON template created to use the action message.
type ResponseActionMessageJSON struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Content []struct {
		Code               string        `json:"code"`
		DeviceID           string        `json:"deviceId"`
		DeviceSerial       string        `json:"deviceSerial"`
		Command            string        `json:"command"`
		DeviceCode         string        `json:"deviceCode"`
		Parameters         interface{}   `json:"parameters"`
		Label              interface{}   `json:"label"`
		ResponseID         string        `json:"responseId"`
		StartDate          interface{}   `json:"startDate"`
		EndDate            interface{}   `json:"endDate"`
		SentStatus         bool          `json:"sentStatus"`
		ErrorMessage       interface{}   `json:"errorMessage"`
		CreatedBy          string        `json:"createdBy"`
		UserMail           string        `json:"userMail"`
		Links              []interface{} `json:"links"`
		SendToDeviceStatus struct {
			Status   interface{} `json:"status"`
			TryCount interface{} `json:"tryCount"`
			Body     interface{} `json:"body"`
			SentDate time.Time   `json:"sentDate"`
		} `json:"sendToDeviceStatus"`
		DeliveryFromDeviceStatus interface{} `json:"deliveryFromDeviceStatus"`
		ExecutionInDeviceStatus  interface{} `json:"executionInDeviceStatus"`
		SentDate                 time.Time   `json:"sentDate"`
	} `json:"content"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}
