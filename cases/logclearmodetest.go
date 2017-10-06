package cases

import (
	"log"
	"os"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/profile"
	"github.com/KMACEL/IITR/rest/workingset"
	"github.com/KMACEL/IITR/writefile"
)

/*
	lcm := cases.LCMTest{}
	lcm.FileName = "ClearTest.xlsx"
	lcm.ModeID = "3BDE218B-3DCC-4B94-8F56-BBB5060B2BC0"
	lcm.PolicyID = "74A56F6D-D197-488F-B4D6-B586B377D3D8"

	lcm.Start("cea9bbd434b04a7db1865d210f449f0e")
*/

//LCMTest is
type LCMTest struct {
	deviceVariable device.Device
	ModeID         string
	PolicyID       string
	FileName       string
}

//Delete Packet

//Start is
func (l LCMTest) Start(setDeviceCode ...string) {
	var logFile *os.File

	writefile.CreateFile(l.FileName)
	logFile = writefile.OpenFile(l.FileName, logFile)
	writefile.WriteText(logFile, "Device ID", "Operation", "Time")

	workingsetKey := workingset.Workingset{}.CreateWorkingset()

retry:

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 5; k++ {
				l.reboot(setDeviceCode, logFile)
				time.Sleep(12 * 60 * time.Second)
			}
			l.getLog(setDeviceCode, logFile)
			time.Sleep(3 * 60 * time.Second)

			l.clearModiverse(setDeviceCode, logFile)
			time.Sleep(10 * 60 * time.Second)

			l.getLog(setDeviceCode, logFile)
			time.Sleep(3 * 60 * time.Second)
		}
		l.changeMode(setDeviceCode, workingsetKey, logFile)
		time.Sleep(12 * 60 * time.Second)

		l.getLog(setDeviceCode, logFile)
		time.Sleep(3 * 60 * time.Second)
	}
	goto retry

}

func (l LCMTest) reboot(setDeviceCodes []string, logFile *os.File) {
	for _, deviceCode := range setDeviceCodes {
		l.deviceVariable.Reboot(deviceCode, rest.Invisible)
		writefile.WriteText(logFile, deviceCode, "Reboot", time.Now().String())
		log.Println(deviceCode, "Reboot")
	}
}

func (l LCMTest) changeMode(setDeviceCodes []string, workingsetKey string, logFile *os.File) {
	for _, deviceCode := range setDeviceCodes {
		workingset.Workingset{}.AddDeviceWorkingSet(workingsetKey, deviceCode)
		profile.Profile{}.PushMode(string(workingsetKey), l.ModeID, l.PolicyID)
		writefile.WriteText(logFile, deviceCode, "PushMode", time.Now().String())
		log.Println(deviceCode, "PushMode")
	}
}

func (l LCMTest) clearModiverse(setDeviceCodes []string, logFile *os.File) {
	for _, deviceCode := range setDeviceCodes {
		device.Device{}.ClearAppData(deviceCode, "com.ardic.android.modiverse", rest.Invisible)
		writefile.WriteText(logFile, deviceCode, "ClearModiverse", time.Now().String())
		log.Println(deviceCode, "ClearModiverse")
	}
}

func (l LCMTest) getLog(setDeviceCodes []string, logFile *os.File) {
	for _, deviceCode := range setDeviceCodes {
		l.deviceVariable.GetDeviceLog(deviceCode, rest.Invisible)
		writefile.WriteText(logFile, deviceCode, "GetLog", time.Now().String())
		log.Println(deviceCode, "GetLog")
	}
}
