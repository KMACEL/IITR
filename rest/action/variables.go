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
	messageJSONVariable MessageJSON
)

/*
     ██╗███████╗ ██████╗ ███╗   ██╗
     ██║██╔════╝██╔═══██╗████╗  ██║
     ██║███████╗██║   ██║██╔██╗ ██║
██   ██║╚════██║██║   ██║██║╚██╗██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
*/

//MessageJSON is
type MessageJSON struct {
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
