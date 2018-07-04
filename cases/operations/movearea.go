package operations

// Posted by Mehmet Akasayan

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KMACEL/IITR/rest/adminarea"
	"github.com/KMACEL/IITR/rest/device"
)

//MoveArea to be edited
type MoveArea struct{}

//Start is
func (m MoveArea) Start(areaName string) {

	var deviceList []string
	inp := bufio.NewScanner(os.Stdin)

	for inp.Scan() {
		deviceList = append(deviceList, inp.Text())
	}
	if err := inp.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Count Array : ", len(deviceList))

	fmt.Println("Device ID, Device Code")

	var d device.Device
	var deviceCodeList []string
	for _, deviceID := range deviceList {
		deviceCode := d.DeviceID2Code(deviceID)
		deviceCodeList = append(deviceCodeList, deviceCode)
		fmt.Println(deviceID, ",", deviceCode)
	}

	var adminAreaVar adminarea.AdminArea

	var adr adminarea.QueryRequirements
	adr.AdminAreaName = areaName
	adr.AddToAdminAreaDeviceCode = deviceCodeList

	adminAreaVar.MoveAdminArea(adr)
}
