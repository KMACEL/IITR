package reports

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/action"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/writefile"
)

/*
██████╗ ███████╗████████╗ █████╗ ██╗██╗             ██████╗ ███████╗██████╗  ██████╗ ██████╗ ████████╗
██╔══██╗██╔════╝╚══██╔══╝██╔══██╗██║██║             ██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝
██║  ██║█████╗     ██║   ███████║██║██║             ██████╔╝█████╗  ██████╔╝██║   ██║██████╔╝   ██║
██║  ██║██╔══╝     ██║   ██╔══██║██║██║             ██╔══██╗██╔══╝  ██╔═══╝ ██║   ██║██╔══██╗   ██║
██████╔╝███████╗   ██║   ██║  ██║██║███████╗        ██║  ██║███████╗██║     ╚██████╔╝██║  ██║   ██║
╚═════╝ ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝        ╚═╝  ╚═╝╚══════╝╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝
*/

//DetailReportDevices is
type DetailReportDevices struct {
	writeCsvArray []string
	devices       device.Device
	actions       action.Action
}

// todo generic hazır olunca sil

//Start is DetailReportDevices. These Cases were created to get detailed reports
func (d DetailReportDevices) Start(fileName string, devicesID []string, setControlPackage []string) {
	// It performs the writing process in one step, not in every step of the way. The goal is to increase
	// the speed and reduce the memory footprint. It is also used to write multiple files at the same time
	//d.writeCsvArray = make([]string, 0)

	// His section allows you to write them in the title if the number of applications to be checked is reached.
	// The reason for the substitution assignment is that it will be used later when writing csv
	var packageHeader string
	for i, packageHeaderName := range setControlPackage {
		if i == 0 {
			packageHeader = packageHeaderName
		} else {
			packageHeader = packageHeaderName + "," + packageHeader
		}
	}

	strings.Trim(packageHeader, " ")
	// The title of the csv file to be made
	d.writeCsvArray = append(d.writeCsvArray, ",,Observed Applications", "\n")

	// Device ID: It shows the number of the device from which we received the report. This number should be unique.
	// packageHeader: This indicates the name of the application on which we have checked the status on the device.
	//                These functions come in series with these names (...). It converts this into a string in a loop in the 'CSV' format.
	// Drom Count: It shows the number of DROM sent to the device. With this number, we learn how much processing is performed on the device.
	// Presence: This program shows "Online" or "Offline" when looking at the device.
	// Profile Name: The mode name on the device indicates.
	// Policy Name: The name of the policy found in the game.
	// Latitude, Longitude: Gives location information.
	// Last Online Time: Gives us the time that was last online.
	// Read Time: The time of receiving information varies according to computer, inernet, cloud speed and is getting longer. This is why it is important when the device is read.
	// Working Group: Any workgroup should be aware of this.
	d.writeCsvArray = append(d.writeCsvArray, "Device ID", packageHeader, "Drom Count", "Presence", "Profile Name", "Policy Name", "Latitude", "Longitude", "Last Online Time", "Read Time", "Working Group", "\n")

	for i, deviceCoding := range devicesID {
		//TODO : Device Code boş olma durumuna bak
		var (
			applicationsStatus string
			presence           string
			lastOnlineTime     string
			profile            string
			policy             string
			dromSize           int
			workingGroup       string
		)

		// GoRutines Message Chanall
		chApplicationsStatus := make(chan string)
		chPresence := make(chan string)
		chLastOnlineTime := make(chan string)
		chProfile := make(chan string)
		chPolicy := make(chan string)
		chDromSize := make(chan int)
		chWorkingGroup := make(chan string)

		// Start GoRutines
		go d.applicationStatus(device.Device{}.DeviceID2Code(deviceCoding), setControlPackage, chApplicationsStatus)
		go d.presenceStatus(deviceCoding, chPresence, chLastOnlineTime)
		go d.profilePolicy(deviceCoding, chProfile, chPolicy)
		go d.submittedDromSize(device.Device{}.DeviceID2Code(deviceCoding), chDromSize)
		go d.workingGroup(deviceCoding, chWorkingGroup)

		// Recieve Message
		applicationsStatus = <-chApplicationsStatus
		presence = <-chPresence
		lastOnlineTime = <-chLastOnlineTime
		profile = <-chProfile
		policy = <-chPolicy
		dromSize = <-chDromSize
		workingGroup = <-chWorkingGroup

	control:
		//TODO : dromSize bak
		if applicationsStatus != "" || presence != "" || lastOnlineTime != "" || profile != "" || policy != "" || workingGroup != "" {
			_, latitude, longitude := device.Device{}.DeviceID2CodeLocation(deviceCoding)
			d.writeCsvArray = append(d.writeCsvArray,
				deviceCoding,
				applicationsStatus,
				strconv.Itoa(dromSize),
				presence,
				profile,
				policy,
				latitude,
				longitude,
				lastOnlineTime, time.Now().String(),
				workingGroup,
				"\n")

		} else {
			goto control
		}

		if i%100 == 0 {
			log.Println(i, "/", len(devicesID))
		}

	}

	d.writeCSVType(fileName, d.writeCsvArray, packageHeader)
	//log.Println("Finish Read ...")
}

/*
 █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗    ███████╗████████╗ █████╗ ████████╗██╗   ██╗███████╗
██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║    ██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝██║   ██║██╔════╝
███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║    ███████╗   ██║   ███████║   ██║   ██║   ██║███████╗
██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║    ╚════██║   ██║   ██╔══██║   ██║   ██║   ██║╚════██║
██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║    ███████║   ██║   ██║  ██║   ██║   ╚██████╔╝███████║
╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝    ╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚══════╝
*/
func (d DetailReportDevices) applicationStatus(deviceCode string, setControlPackage []string, chApplicationsStatus chan string) string {
	var statusGlobal string
	packageStatus := make([]string, 0)

	applicationQuery := d.devices.GetDownloadedApplicationsList(deviceCode, rest.NOMarshal, rest.Invisible)
	if applicationQuery != nil {
		if string(applicationQuery) != rest.ResponseNotFound {

			deviceApplication := device.DownloadedApplicationListJSON{}
			json.Unmarshal(applicationQuery, &deviceApplication)

			// Looking at the situation will return to the desired application.
			for _, controlPackages := range setControlPackage {
				controlApp := false

				// Related device applications
				for _, downloadedApp := range deviceApplication {

					// First check. Compare these applications with installed applications.
					if downloadedApp.PackageName == controlPackages {
						// Checked
						controlApp = true

						// the working status of the relevant application came up
						running := downloadedApp.Running
						if running {
							statusGlobal = "Running"
							packageStatus = append(packageStatus, statusGlobal)
						} else {
							statusGlobal = "Not Running"
							packageStatus = append(packageStatus, statusGlobal)
						}
						break
					}
				}

				// Absence of application
				if controlApp == false {
					statusGlobal = "No Application"
					packageStatus = append(packageStatus, statusGlobal)
				}
			}

			// If 404 Not Found is
		} else {
			for j := 0; j < len(setControlPackage); j++ {
				statusGlobal = rest.ResponseNotFound
				packageStatus = append(packageStatus, statusGlobal)
			}
		}

		// If Nil is
	} else {
		for j := 0; j < len(setControlPackage); j++ {
			statusGlobal = rest.ResponseNil
			packageStatus = append(packageStatus, statusGlobal)
		}
	}

	var applications string
	for i, applicationLo := range packageStatus {
		if applicationLo != "" {
			if i == 0 {
				applications = applicationLo
			} else {
				applications = applicationLo + "," + applications
			}
		}
	}
	strings.Trim(applications, " ")
	chApplicationsStatus <- applications
	close(chApplicationsStatus)
	return applications
}

/*
██████╗ ██████╗ ███████╗███████╗███████╗███╗   ██╗ ██████╗███████╗    ███████╗████████╗ █████╗ ████████╗██╗   ██╗███████╗
██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝████╗  ██║██╔════╝██╔════╝    ██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝██║   ██║██╔════╝
██████╔╝██████╔╝█████╗  ███████╗█████╗  ██╔██╗ ██║██║     █████╗      ███████╗   ██║   ███████║   ██║   ██║   ██║███████╗
██╔═══╝ ██╔══██╗██╔══╝  ╚════██║██╔══╝  ██║╚██╗██║██║     ██╔══╝      ╚════██║   ██║   ██╔══██║   ██║   ██║   ██║╚════██║
██║     ██║  ██║███████╗███████║███████╗██║ ╚████║╚██████╗███████╗    ███████║   ██║   ██║  ██║   ██║   ╚██████╔╝███████║
╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═══╝ ╚═════╝╚══════╝    ╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚══════╝
*/
func (d DetailReportDevices) presenceStatus(deviceCode string, chPresence, chLastOnlineTime chan string) (string, string) {
	var presence string
	var lastOnlineTime string

	presenceTime := d.devices.PresenceInfo(deviceCode, rest.NOMarshal, rest.Invisible)
	if presenceTime != nil {
		if string(presenceTime) != rest.ResponseNotFound {

			presenceInfo := device.PresenceInfoJSON{}
			json.Unmarshal(presenceTime, &presenceInfo)

			lastOnlineTime = time.Unix(0, presenceInfo.CreateDate*1000000).String()
			presence = presenceInfo.Data.State
		} else {
			lastOnlineTime = rest.ResponseNotFound
			presence = rest.ResponseNotFound
		}
	} else {
		lastOnlineTime = rest.ResponseNil
		presence = rest.ResponseNil
	}
	chPresence <- presence
	chLastOnlineTime <- lastOnlineTime
	close(chPresence)
	close(chLastOnlineTime)
	return presence, lastOnlineTime
}

/*
██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗    ██████╗  ██████╗ ██╗     ██╗ ██████╗██╗   ██╗
██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝    ██╔══██╗██╔═══██╗██║     ██║██╔════╝╚██╗ ██╔╝
██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗      ██████╔╝██║   ██║██║     ██║██║      ╚████╔╝
██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝      ██╔═══╝ ██║   ██║██║     ██║██║       ╚██╔╝
██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗    ██║     ╚██████╔╝███████╗██║╚██████╗   ██║
╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝    ╚═╝      ╚═════╝ ╚══════╝╚═╝ ╚═════╝   ╚═╝
*/
func (d DetailReportDevices) profilePolicy(deviceID string, chProfile, chPolicy chan string) (string, string) {
	var (
		activeProfile string
		activePolicy  string
	)

	profilePolicyQuery := d.devices.ActiveProfilePolicy(deviceID, rest.NOMarshal, rest.Invisible)
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
███████╗██╗   ██╗██████╗ ███╗   ███╗██╗████████╗████████╗███████╗██████╗     ██████╗ ██████╗  ██████╗ ███╗   ███╗    ███████╗██╗███████╗███████╗
██╔════╝██║   ██║██╔══██╗████╗ ████║██║╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗    ██╔══██╗██╔══██╗██╔═══██╗████╗ ████║    ██╔════╝██║╚══███╔╝██╔════╝
███████╗██║   ██║██████╔╝██╔████╔██║██║   ██║      ██║   █████╗  ██║  ██║    ██║  ██║██████╔╝██║   ██║██╔████╔██║    ███████╗██║  ███╔╝ █████╗
╚════██║██║   ██║██╔══██╗██║╚██╔╝██║██║   ██║      ██║   ██╔══╝  ██║  ██║    ██║  ██║██╔══██╗██║   ██║██║╚██╔╝██║    ╚════██║██║ ███╔╝  ██╔══╝
███████║╚██████╔╝██████╔╝██║ ╚═╝ ██║██║   ██║      ██║   ███████╗██████╔╝    ██████╔╝██║  ██║╚██████╔╝██║ ╚═╝ ██║    ███████║██║███████╗███████╗
╚══════╝ ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝   ╚═╝      ╚═╝   ╚══════╝╚═════╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝    ╚══════╝╚═╝╚══════╝╚══════╝
*/
func (d DetailReportDevices) submittedDromSize(deviceCode string, chDromSize chan int) int {
	var dromCounter int

	actionStatueQuery := d.actions.GetActionStatus(deviceCode, "PUSH_CMD_DROM", 1000, rest.Invisible)
	if actionStatueQuery != nil {
		if string(actionStatueQuery) != rest.ResponseNotFound {

			actionMessage := action.MessageJSON{}
			json.Unmarshal(actionStatueQuery, &actionMessage)

			for _, submittedDrom := range actionMessage.Content {
				if submittedDrom.SentStatus {
					dromCounter++
				}
			}
		} else {
			dromCounter = -1
		}
	} else {
		dromCounter = -5
	}
	chDromSize <- dromCounter
	close(chDromSize)
	return dromCounter
}

/*
██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗██╗███╗   ██╗ ██████╗      ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗
██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝██║████╗  ██║██╔════╝     ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗
██║ █╗ ██║██║   ██║██████╔╝█████╔╝ ██║██╔██╗ ██║██║  ███╗    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝
██║███╗██║██║   ██║██╔══██╗██╔═██╗ ██║██║╚██╗██║██║   ██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝
╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗██║██║ ╚████║╚██████╔╝    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║
 ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝
*/
func (d DetailReportDevices) workingGroup(deviceID string, chWorkingGroup chan string) string {
	workingGroup := d.devices.WorkingGroupControl(deviceID, rest.Invisible)
	chWorkingGroup <- workingGroup
	close(chWorkingGroup)
	return workingGroup
}

/*
██╗    ██╗██████╗ ██╗████████╗███████╗     ██████╗███████╗██╗   ██╗    ████████╗██╗   ██╗██████╗ ███████╗
██║    ██║██╔══██╗██║╚══██╔══╝██╔════╝    ██╔════╝██╔════╝██║   ██║    ╚══██╔══╝╚██╗ ██╔╝██╔══██╗██╔════╝
██║ █╗ ██║██████╔╝██║   ██║   █████╗      ██║     ███████╗██║   ██║       ██║    ╚████╔╝ ██████╔╝█████╗
██║███╗██║██╔══██╗██║   ██║   ██╔══╝      ██║     ╚════██║╚██╗ ██╔╝       ██║     ╚██╔╝  ██╔═══╝ ██╔══╝
╚███╔███╔╝██║  ██║██║   ██║   ███████╗    ╚██████╗███████║ ╚████╔╝        ██║      ██║   ██║     ███████╗
 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝   ╚═╝   ╚══════╝     ╚═════╝╚══════╝  ╚═══╝         ╚═╝      ╚═╝   ╚═╝     ╚══════╝
*/
func (d DetailReportDevices) writeCSVType(fileName string, writeCSVArray []string, setControlPackage string) {
	var detailReportFile *os.File
	writefile.CreateFile(fileName)
	detailReportFile = writefile.OpenFile(detailReportFile, fileName)
	writefile.WriteArray(detailReportFile, writeCSVArray)

	log.Println("Finish Write : ", setControlPackage)
}
