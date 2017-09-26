package cases

import (
	"encoding/json"
	"log"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/profile"
	"github.com/KMACEL/IITR/rest/workingset"
	"github.com/KMACEL/IITR/writefile"

	"math/rand"
	"time"
)

//NoTrust is
type NoTrust struct {
	presenceInfoJSONVariable device.PresenceInfoJSON
	workingsetVar            workingset.Workingset
	deviceVar                device.Device
}

const (
	controlSize         = 5
	responseSuccess     = "200"
	timeCoefficient     = 60
	sleepNotBlockedTime = 10
)

//Start is
func (n NoTrust) Start(packageList ...string) {
	writefile.CreateFile("BlockedApp.xlsx")
	writefile.OpenFile("BlockedApp.xlsx")
	writefile.WriteText("Package", "Response Message Statue", "Response Message Description", "Time")

	workingsetKey := n.workingsetVar.CreateWorkingset()
	deviceCode := "cea9bbd434b04a7db1865d210f449f0e"
	deviceID := "867377020747089"
	modeCode := "3BDE218B-3DCC-4B94-8F56-BBB5060B2BC0"
	policyCode := "74A56F6D-D197-488F-B4D6-B586B377D3D8"

retry:
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randNumber := r1.Intn(controlSize)

	n.workingsetVar.AddDeviceWorkingSet(workingsetKey, deviceCode)
	profile.Profile{}.PushMode(string(workingsetKey), modeCode, policyCode)

	for i := 0; i < randNumber; i++ {
		time.Sleep(timeCoefficient * time.Second)
	}
	n.deviceVar.Reboot(deviceCode, rest.Invisible)

	for i := 0; i < controlSize-randNumber; i++ {
		time.Sleep(timeCoefficient * time.Second)
	}
	var counter int

retryApplication:
	query := n.deviceVar.PresenceInfo(deviceID, rest.NOMarshal, rest.Invisible)
	if query != nil {
		if string(query) == rest.ResponseNotFound {

			n.presenceInfoJSONVariable = device.PresenceInfoJSON{}
			json.Unmarshal(query, &n.presenceInfoJSONVariable)
			if n.presenceInfoJSONVariable.Data.State == rest.Online {
				for _, packageVar := range packageList {
					messageResponse := n.deviceVar.AppSS(device.StartApp, deviceCode, packageVar, rest.Invisible)
					responseMessageStatue, responseMessageDescription := n.deviceVar.MessageControl(deviceCode, messageResponse, rest.Invisible)
					log.Println("Statue : ", responseMessageStatue)
					if responseMessageStatue == responseSuccess {
						writefile.WriteText(packageVar, responseMessageStatue, responseMessageDescription, time.Now().String())
						counter++
						log.Println("Not Blocked Error : ", packageVar)
						if counter > controlSize {
							return
						}
						time.Sleep(sleepNotBlockedTime * timeCoefficient * time.Second)
						goto retryApplication
					}
					writefile.WriteText(packageVar, responseMessageStatue, responseMessageDescription, time.Now().String())
				}

				writefile.WriteText()
			} else {
				writefile.WriteText(deviceID, rest.Offline)
			}
			goto retry
		}
	}
}
