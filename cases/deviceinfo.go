package cases

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
)

//Delete Packet

//DeviceInformation is
type DeviceInformation struct{}

//TODO: Yapılacak

//Start is
func (d DeviceInformation) Start() {
	var counter int
	query := device.Device{}.LocationMap(rest.NOMarshal, rest.Invisible)
	if query != nil {
		if string(query) != rest.ResponseNotFound {

			// This assignment is aimed at resetting the variable
			deviceCode := device.LocationJSON{}
			json.Unmarshal(query, &deviceCode)

			for _, deviceCoding := range deviceCode.Extras {

				queryInformation := device.Device{}.Informations(deviceCoding.DeviceCode, rest.NOMarshal, rest.Invisible)

				if queryInformation != nil {
					if string(queryInformation) != rest.ResponseNotFound {

						deviceInformation := device.InformationJSON{}
						json.Unmarshal(queryInformation, &deviceInformation)
						if deviceInformation.ModeAppVersion == "AR.AMP.r2.1.329" {
							fmt.Println(deviceInformation.DeviceID, ",", deviceInformation.ModeAppVersion, ",", deviceInformation.Presence.State, ",", time.Unix(0, deviceInformation.DetailLastModifiedDate*1000000).String(), ",")
							counter++
						}
					}
				}
			}
		}
	}
	log.Println("Total : ", counter)
}
