package reports

import (
	"fmt"
	"os"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/workgroup"
	"github.com/KMACEL/IITR/writefile"
)

/*
██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗███████╗        ██╗███╗   ██╗        ██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗         ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗
██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝██╔════╝        ██║████╗  ██║        ██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝        ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗
██║  ██║█████╗  ██║   ██║██║██║     █████╗  ███████╗        ██║██╔██╗ ██║        ██║ █╗ ██║██║   ██║██████╔╝█████╔╝         ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝
██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝  ╚════██║        ██║██║╚██╗██║        ██║███╗██║██║   ██║██╔══██╗██╔═██╗         ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝
██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗███████║        ██║██║ ╚████║        ╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗        ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║
╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝╚══════╝        ╚═╝╚═╝  ╚═══╝         ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝         ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝
*/

// DevicesInWorkGroup is
type DevicesInWorkGroup struct {
	WorkgroupName string
}

// Start is
func (d DevicesInWorkGroup) Start() {

	var f *os.File

	timeOperation := time.Now().String()

	var fileName = "WorkGroupDevices_" +
		timeOperation[0:4] + timeOperation[5:7] + timeOperation[8:10] +
		//	"_" +
		//	timeOperation[11:13] + timeOperation[14:16] +
		".csv"
	fmt.Println("Output file name: ", fileName)

	writefile.CreateFile(fileName)
	f = writefile.OpenFile(f, fileName)
	writefile.WriteText(f, "Device ID", "OS Display", "ModeApp Version")

	var wg workgroup.WorkGroup

	wgID := wg.GetGroupID(d.WorkgroupName, rest.Invisible)
	fmt.Println("Working Group Name: ", d.WorkgroupName, "\nWorking Group ID: ", wgID)

	query := wg.GetGroupDevices(wgID, rest.Invisible)

	deviceID := wg.GetGroupDeviceIDs(query)
	fmt.Println("Device IDs in Work Group: ", deviceID)

	osDisplay := wg.GetGroupDeviceOSs(query)
	fmt.Println("Device OSs in Work Group: ", osDisplay)

	modeVersions := wg.GetGroupModeVersions(query)
	fmt.Println("Mode Versions in Work Group: ", modeVersions)

	for i := 0; i < len(deviceID); i++ {
		writefile.WriteText(f, deviceID[i], osDisplay[i], modeVersions[i])
	}
}
