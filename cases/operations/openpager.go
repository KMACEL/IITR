package operations

import (
	"encoding/json"
	"fmt"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

//OpenPager is
type OpenPager struct {
	deviceVariable device.Device
}

var (
	//deviceVariable device.Device
	controlPackage = "tr.com.innology.taksipager"
)

//Start is
func (o OpenPager) Start() {
	fmt.Println("Entry Open Pager")
	query := o.deviceVariable.LocationMap(rest.NOMarshal, rest.Invisible)

	var locationJSONVariable device.LocationJSON
	json.Unmarshal(query, &locationJSONVariable)

	for _, deviceCoding := range locationJSONVariable.Extras {

		if deviceCoding.DeviceCode != "" && deviceCoding.DeviceID != "" && deviceCoding.Presence != rest.Offline {
			queryDownload := o.deviceVariable.GetDownloadedApplicationsList(deviceCoding.DeviceCode, rest.NOMarshal, rest.Invisible)
			downloadedApplicationListJSONVariable := device.DownloadedApplicationListJSON{}
			json.Unmarshal(queryDownload, &downloadedApplicationListJSONVariable)

			for _, downloadedApp := range downloadedApplicationListJSONVariable {
				if downloadedApp.PackageName == controlPackage {
					if !downloadedApp.Running {
						o.deviceVariable.AppSS(device.StartApp, deviceCoding.DeviceCode, downloadedApp.PackageName, rest.Invisible)
						fmt.Println("Running App Device ID : ", deviceCoding.DeviceID)
					}
				}
			}
		}
	}
}
