package cases

import (
	"log"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/profile"
	"github.com/KMACEL/IITR/rest/workingset"
	"github.com/KMACEL/IITR/writefile"
)

//LogTest is
type LogTest struct {
	deviceVariable device.Device
}

//Start is
func (l LogTest) Start(setDeviceCode string, fileName string, vasualFlag bool) {
	writefile.CreateFile(fileName)
	writefile.OpenFile(fileName)
	writefile.WriteText("Device ID", "Operation", "Time")

	workingsetKey := workingset.Workingset{}.CreateWorkingset()

retry:
	for k := 0; k < 10; k++ {
		for j := 0; j < 3; j++ {
			for i := 0; i < 5; i++ {
				l.deviceVariable.Reboot(setDeviceCode, rest.Invisible)
				writefile.WriteText(setDeviceCode, "Reboot", time.Now().String())
				log.Println(setDeviceCode, "Reboot")
				time.Sleep(12 * 60 * time.Second)
			}
			l.getLog(setDeviceCode, vasualFlag)
		}

		l.pushMode(workingsetKey, setDeviceCode)
		writefile.WriteText("")
	}

	for i := 0; i < 80; i++ {
		l.deviceVariable.Reboot(setDeviceCode, rest.Invisible)
		writefile.WriteText(setDeviceCode, "Reboot", time.Now().String())
		log.Println(setDeviceCode, "Reboot")
		time.Sleep(12 * 60 * time.Second)
	}
	l.getLog(setDeviceCode, vasualFlag)

	goto retry

}

func (l LogTest) pushMode(workingsetKey string, setDeviceCode string) {
	modeCode := "3BDE218B-3DCC-4B94-8F56-BBB5060B2BC0"
	policyCode := "74A56F6D-D197-488F-B4D6-B586B377D3D8"

	workingset.Workingset{}.AddDeviceWorkingSet(workingsetKey, setDeviceCode)
	profile.Profile{}.PushMode(string(workingsetKey), modeCode, policyCode)

	writefile.WriteText("PushMode", time.Now().String())
	log.Println(setDeviceCode, "PushMode")

	time.Sleep(1 * 60 * time.Second)
}

func (l LogTest) getLog(setDeviceCode string, vasualFlag bool) {
	l.deviceVariable.GetDeviceLog(setDeviceCode, vasualFlag)

	writefile.WriteText(setDeviceCode, "GetLog", time.Now().String())
	log.Println(setDeviceCode, "GetLog")

	time.Sleep(7 * 60 * time.Second)
}
