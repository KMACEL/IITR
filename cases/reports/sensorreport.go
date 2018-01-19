package reports

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/writefile"
)

//SensorData is
type SensorData struct {
	NodeID    string
	ThingID   string
	DelayTime time.Duration
}

// Start is
func (s SensorData) Start(devicesID ...string) {
	var d device.Device
	var f *os.File
	var fileName = "SensorReport.xlsx"

	writefile.CreateFile(fileName)
	f = writefile.OpenFile(f, fileName)
	writefile.WriteText(f, "Device ID", "Data", "Create Time")

	if s.NodeID != "" && s.ThingID != "" && devicesID != nil {
		for i, deviceID := range devicesID {
			getSensorInformation := d.GetSensorData(deviceID, s.NodeID, s.ThingID, rest.NOMarshal, rest.Invisible)

			var sensorDataJSON device.SensorDataJSON
			json.Unmarshal(getSensorInformation, &sensorDataJSON)
			writefile.WriteText(f, deviceID, sensorDataJSON.Data.Data, time.Unix(0, sensorDataJSON.Data.CreateDate*1000000).String())

			if s.DelayTime != 0 {
				time.Sleep(s.DelayTime * time.Second)
			}
			if len(devicesID) > 100 {
				if i%100 == 0 {
					log.Println(i, "/", len(devicesID))
				}
			} else if len(devicesID) > 10 {
				if i%10 == 0 {
					log.Println(i, "/", len(devicesID))
				}
			}
		}

	} else {
		fmt.Println("Nil")
	}
}
