package operations

// Posted by Mehmet Akasayan

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/workgroup"
)

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗         ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝        ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██║ █╗ ██║██║   ██║██████╔╝█████╔╝         ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██║███╗██║██║   ██║██╔══██╗██╔═██╗         ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗        ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝         ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝
*/

// CreateWorkGroup creates a new user group and identifies the "deviceIDs" in the given file.
type CreateWorkGroup struct{}

// Start is the function of ClearApp. This function creates a group with the name of the parameter given when the program is run.
func (c CreateWorkGroup) Start() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	txtName := strings.Split(os.Args[1], ".")[0]

	fmt.Println("Work Group Name: ", txtName)
	var deviceList []string

	devList := strings.Split(string(data), "\n")
	for _, deviceID := range devList {
		if deviceID != "" {
			deviceList = append(deviceList, deviceID)
		}
	}

	fmt.Println("Number of Devices: ", len(deviceList))
	if len(deviceList) == 0 {
		panic("No data to process ...")
	}

	fmt.Println("Device ID, Device Code")

	var d device.Device
	var deviceCodeList []string
	for _, deviceID := range deviceList {
		deviceCode := d.DeviceID2Code(deviceID)
		if deviceCode != rest.ResponseNil {
			deviceCodeList = append(deviceCodeList, deviceCode)
			fmt.Println(deviceID, ",", deviceCode)

		} else {
			fmt.Println(deviceID, ",", "Can not Add to Work Group ...")
		}
	}

	var workgroupVar workgroup.WorkGroup

	var wgr workgroup.WorkGroupRequirements
	wgr.WorkGroupName = txtName
	wgr.AddToWorkGroupDeviceCode = deviceCodeList

	workgroupVar.MoveWorkGroup(wgr)
}
