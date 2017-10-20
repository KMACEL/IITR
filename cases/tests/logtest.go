package tests

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

//Delete Packet

//LogTest is
type LogTest struct {
	deviceVariable device.Device
}

//Start is
func (l LogTest) Start(setDeviceCode string, fileName string, vasualFlag bool) {
	var logTestFile *os.File

	writefile.CreateFile(fileName)
	logTestFile = writefile.OpenFile(fileName, logTestFile)
	writefile.WriteText(logTestFile, "Device ID", "Operation", "Time")

	workingsetKey := workingset.Workingset{}.CreateWorkingset()

retry:
	for k := 0; k < 10; k++ {
		for j := 0; j < 3; j++ {
			for i := 0; i < 5; i++ {
				l.deviceVariable.Reboot(setDeviceCode, rest.Invisible)
				writefile.WriteText(logTestFile, setDeviceCode, "Reboot", time.Now().String())
				log.Println(setDeviceCode, "Reboot")
				time.Sleep(12 * 60 * time.Second)
			}
			l.getLog(setDeviceCode, vasualFlag, logTestFile)
		}

		l.pushMode(workingsetKey, setDeviceCode, logTestFile)
		writefile.WriteText(logTestFile, "")
	}

	for i := 0; i < 80; i++ {
		l.deviceVariable.Reboot(setDeviceCode, rest.Invisible)
		writefile.WriteText(logTestFile, setDeviceCode, "Reboot", time.Now().String())
		log.Println(setDeviceCode, "Reboot")
		time.Sleep(12 * 60 * time.Second)
	}
	l.getLog(setDeviceCode, vasualFlag, logTestFile)

	goto retry

}

func (l LogTest) pushMode(workingsetKey string, setDeviceCode string, logTestFile *os.File) {
	modeCode := "3BDE218B-3DCC-4B94-8F56-BBB5060B2BC0"
	policyCode := "74A56F6D-D197-488F-B4D6-B586B377D3D8"

	workingset.Workingset{}.AddDeviceWorkingSet(workingsetKey, setDeviceCode)
	profile.Profile{}.PushMode(string(workingsetKey), modeCode, policyCode)

	writefile.WriteText(logTestFile, "PushMode", time.Now().String())
	log.Println(setDeviceCode, "PushMode")

	time.Sleep(1 * 60 * time.Second)
}

func (l LogTest) getLog(setDeviceCode string, vasualFlag bool, logTestFile *os.File) {
	l.deviceVariable.GetDeviceLog(setDeviceCode, vasualFlag)

	writefile.WriteText(logTestFile, setDeviceCode, "GetLog", time.Now().String())
	log.Println(setDeviceCode, "GetLog")

	time.Sleep(7 * 60 * time.Second)
}
