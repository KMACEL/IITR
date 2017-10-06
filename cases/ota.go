package cases

import (
	"encoding/json"
	"gerrit/DeviceReportApp/rest"
	"gerrit/DeviceReportApp/rest/device"
)

// OtaUpdater is
type OtaUpdater struct {
}

// Start is
func (o OtaUpdater) Start() {
	var (
		devices device.Device
	)
	query := devices.LocationMap(rest.NOMarshal, rest.Invisible)
	if query != nil {
		if string(query) != rest.ResponseNotFound {

			// This assignment is aimed at resetting the variable
			deviceCode := device.LocationJSON{}
			json.Unmarshal(query, &deviceCode)

			for _, deviceCoding := range deviceCode.Extras {
				if deviceCoding.DeviceID != "" {

				}
			}
		}
	}
}

func (o OtaUpdater) getPresences() {

}

func (o OtaUpdater) refleshGateway() {
}

func (o OtaUpdater) getCurrentResource() {
}

func (o OtaUpdater) pushApplication() {
}

func (o OtaUpdater) startApp() {
}

func (o OtaUpdater) writeReport() {
}

func (o OtaUpdater) removeApplication() {
}
