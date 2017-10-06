package cases

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/writefile"
)

//Delete Packet

// BlockedAppList is
type BlockedAppList struct {
	writeCsvArray []string
	devices       device.Device
}

// Start is
func (b BlockedAppList) Start(applicationList ...string) {

	var packageHeader string
	for i, packageHeaderName := range applicationList {
		if i == 0 {
			packageHeader = packageHeaderName + "," + packageHeaderName
		} else {
			packageHeader = packageHeaderName + "," + packageHeaderName + "," + packageHeader
		}
	}

	strings.Trim(packageHeader, " ")
	b.writeCsvArray = append(b.writeCsvArray, "Device ID", packageHeader, "Profile", "Policy", "\n")

	query := b.devices.LocationMap(rest.NOMarshal, rest.Invisible)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			deviceCode := device.LocationJSON{}
			json.Unmarshal(query, &deviceCode)

			for i, deviceCoding := range deviceCode.Extras {

				var (
					applicationsStatus string
					profile            string
					policy             string
				)

				chApplicationsStatus := make(chan string)
				chProfile := make(chan string)
				chPolicy := make(chan string)

				go b.applicationStatus(deviceCoding.DeviceID, applicationList, chApplicationsStatus)
				go b.profilePolicy(deviceCoding.DeviceID, chProfile, chPolicy)

				applicationsStatus = <-chApplicationsStatus
				profile = <-chProfile
				policy = <-chPolicy

			control:
				if applicationsStatus != "" || profile != "" || policy != "" {
					b.writeCsvArray = append(b.writeCsvArray, deviceCoding.DeviceID, applicationsStatus, profile, policy, "\n")
				} else {
					goto control
				}

				if i%100 == 0 {
					log.Println("Processes : ", i, " / ", len(deviceCode.Extras))
				}
			}

			b.writeCSVType("Blocked.xlsx", b.writeCsvArray)
			log.Println("Finish Read ...")
		}
	}
}

/*
 █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗    ███████╗████████╗ █████╗ ████████╗██╗   ██╗███████╗
██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║    ██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝██║   ██║██╔════╝
███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║    ███████╗   ██║   ███████║   ██║   ██║   ██║███████╗
██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║    ╚════██║   ██║   ██╔══██║   ██║   ██║   ██║╚════██║
██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║    ███████║   ██║   ██║  ██║   ██║   ╚██████╔╝███████║
╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝    ╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚══════╝
*/
func (b BlockedAppList) applicationStatus(deviceID string, applicationList []string, chApplicationsStatus chan string) string {
	applicationQuery := b.devices.ApplicationInfo(deviceID, rest.NOMarshal, rest.Invisible)

	applicationsInfo := device.ApplicationInfoJSON{}
	json.Unmarshal(applicationQuery, &applicationsInfo)

	var applicationArray []string

	for _, applicationsLooking := range applicationList {

		var (
			findControl    int
			notFindControl int
			runningControl string
			blockedControl string
		)

		for _, apllicationStatus := range applicationsInfo.Data {
			if applicationsLooking == apllicationStatus.PackageName {
				if apllicationStatus.Running {
					runningControl = rest.Running
				} else {
					runningControl = rest.NotRunning

				}

				if apllicationStatus.Blocked == -1 {
					blockedControl = rest.UnKnow
				} else if apllicationStatus.Blocked == 0 {
					blockedControl = rest.NotBlocked
				} else if apllicationStatus.Blocked == 1 {
					blockedControl = rest.Blocked
				}

				applicationArray = append(applicationArray, runningControl, blockedControl)

				findControl++
				break
			} else {
				notFindControl++
			}
		}
		if notFindControl > 0 && findControl == 0 {
			applicationArray = append(applicationArray, rest.ResponseNil, rest.ResponseNil)
		}
	}

	var applicationStatusString string

	for i, getApplicationName := range applicationArray {
		if i == 0 {
			applicationStatusString = getApplicationName + ","
		} else {
			applicationStatusString = getApplicationName + "," + applicationStatusString
		}
	}

	strings.Trim(applicationStatusString, " ")

	chApplicationsStatus <- applicationStatusString
	close(chApplicationsStatus)
	return applicationStatusString
}

/*
██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗    ██████╗  ██████╗ ██╗     ██╗ ██████╗██╗   ██╗
██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝    ██╔══██╗██╔═══██╗██║     ██║██╔════╝╚██╗ ██╔╝
██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗      ██████╔╝██║   ██║██║     ██║██║      ╚████╔╝
██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝      ██╔═══╝ ██║   ██║██║     ██║██║       ╚██╔╝
██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗    ██║     ╚██████╔╝███████╗██║╚██████╗   ██║
╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝    ╚═╝      ╚═════╝ ╚══════╝╚═╝ ╚═════╝   ╚═╝
*/
func (b BlockedAppList) profilePolicy(deviceID string, chProfile, chPolicy chan string) (string, string) {
	var (
		activeProfile string
		activePolicy  string
	)

	profilePolicyQuery := b.devices.ActiveProfilePolicy(deviceID, rest.NOMarshal, rest.Invisible)
	if profilePolicyQuery != nil {
		if string(profilePolicyQuery) != rest.ResponseNotFound {
			activeProfilePolicy := device.ActiveProfilePolicyJSON{}
			json.Unmarshal(profilePolicyQuery, &activeProfilePolicy)

			activeProfile = activeProfilePolicy.ActiveProfile
			activePolicy = activeProfilePolicy.ActivePolicy

			if len(activeProfile) == 0 {
				activeProfile = rest.ResponseNil
			}
			if len(activePolicy) == 0 {
				activePolicy = rest.ResponseNil
			}

		} else {
			activeProfile = rest.ResponseNotFound
			activePolicy = rest.ResponseNotFound
		}
	} else {
		activeProfile = rest.ResponseNil
		activePolicy = rest.ResponseNil
	}
	chProfile <- activeProfile
	chPolicy <- activePolicy

	close(chProfile)
	close(chPolicy)
	return activeProfile, activePolicy
}

/*
██╗    ██╗██████╗ ██╗████████╗███████╗     ██████╗███████╗██╗   ██╗    ████████╗██╗   ██╗██████╗ ███████╗
██║    ██║██╔══██╗██║╚══██╔══╝██╔════╝    ██╔════╝██╔════╝██║   ██║    ╚══██╔══╝╚██╗ ██╔╝██╔══██╗██╔════╝
██║ █╗ ██║██████╔╝██║   ██║   █████╗      ██║     ███████╗██║   ██║       ██║    ╚████╔╝ ██████╔╝█████╗
██║███╗██║██╔══██╗██║   ██║   ██╔══╝      ██║     ╚════██║╚██╗ ██╔╝       ██║     ╚██╔╝  ██╔═══╝ ██╔══╝
╚███╔███╔╝██║  ██║██║   ██║   ███████╗    ╚██████╗███████║ ╚████╔╝        ██║      ██║   ██║     ███████╗
 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝   ╚═╝   ╚══════╝     ╚═════╝╚══════╝  ╚═══╝         ╚═╝      ╚═╝   ╚═╝     ╚══════╝
*/
func (b BlockedAppList) writeCSVType(fileName string, writeCSVArray []string) {
	var blockedControlFile *os.File

	writefile.CreateFile(fileName)
	blockedControlFile = writefile.OpenFile(fileName, blockedControlFile)
	writefile.WriteArray(writeCSVArray, blockedControlFile)

	log.Println("Finish Write : ")
}
