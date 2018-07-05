package reports

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KMACEL/IITR/logc"
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

//DetailAllReport is
type DetailAllReport struct {
	writeCsvArray []string
	devices       device.Device
	actions       action.Action
}

// todo generic hazır olunca sil

//Start is DetailAllReport. These Cases were created to get detailed reports
func (d DetailAllReport) Start(fileName string, setControlPackage []string) {
	// It performs the writing process in one step, not in every step of the way. The goal is to increase
	// the speed and reduce the memory footprint. It is also used to write multiple files at the same time
	//d.writeCsvArray = make([]string, 0)

	// His section allows you to write them in the title if the number of applications to be checked is reached.
	// The reason for the substitution assignment is that it will be used later when writing csv
	var packageHeader string
	for i, packageHeaderName := range setControlPackage {
		if i == 0 {
			packageHeader = packageHeaderName + "," + packageHeaderName
		} else {
			packageHeader = packageHeaderName + "," + packageHeaderName + "," + packageHeader
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

	// For device information query
	query := d.devices.LocationMap(rest.NOMarshal, rest.Invisible)
	if query != nil {
		if string(query) != rest.ResponseNotFound {

			// This assignment is aimed at resetting the variable
			deviceCode := device.LocationJSON{}
			json.Unmarshal(query, &deviceCode)

			for i, deviceCoding := range deviceCode.Extras {
				if deviceCoding.DeviceID != "" {

					var (
						applicationsStatus string
						presence           string
						lastOnlineTime     string
						profile            string
						policy             string
						dromSize           int
						workingGroup       string
					)

					// GoRoutine Message Channel
					// chApplicationsStatus: The channel in the string type that gives application information.
					// chPresence: Online - The channel in the string type that provides offline information.
					// chLastOnlineTime: The channel in the string type that gives the time to be last online.
					// chProfile: The channel in string type that gives active mode information.
					// chPolicy: The channel in string type that provides active policy information.
					// chDromSize is the integer type channel that reports the number of droms
					// chWorkingGroup: The channel in the string type that reports the group.

					chApplicationsStatus := make(chan string)
					chPresence := make(chan string)
					chLastOnlineTime := make(chan string)
					chProfile := make(chan string)
					chPolicy := make(chan string)
					chDromSize := make(chan int)
					chWorkingGroup := make(chan string)

					// Start GoRutines.
					// applicationStatus: Returns the application status and block status of applications that are initially given a package name.
					// presenceStatus: The status of the device is online - offline. If offline, it will tell you when it was last online.
					// profilePolicy: Provides the mode and policy information of the device.
					// submittedDromSize: How many DROMs are sent to the device.
					// workingGroup: Tells us if there is a group on the device.
					go d.applicationStatus(deviceCoding.DeviceID, setControlPackage, chApplicationsStatus)
					go d.presenceStatus(deviceCoding.DeviceID, chPresence, chLastOnlineTime)
					go d.profilePolicy(deviceCoding.DeviceID, chProfile, chPolicy)
					go d.submittedDromSize(deviceCoding.DeviceCode, chDromSize)
					go d.workingGroup(deviceCoding.DeviceID, chWorkingGroup)

					// This section writes the messages from the channels in the go routines to the variables.
					for getItemApplicationsStatus, status := <-chApplicationsStatus; status; getItemApplicationsStatus, status = <-chApplicationsStatus {
						applicationsStatus = getItemApplicationsStatus
						if status {
							break
						}
					}

					for getItemPresence, status := <-chPresence; status; getItemPresence, status = <-chPresence {
						presence = getItemPresence
						if status {
							break
						}
					}

					for getItemLastOnlineTime, status := <-chLastOnlineTime; status; getItemLastOnlineTime, status = <-chLastOnlineTime {
						lastOnlineTime = getItemLastOnlineTime
						if status {
							break
						}
					}

					for getItemProfile, status := <-chProfile; status; getItemProfile, status = <-chProfile {
						profile = getItemProfile
						if status {
							break
						}
					}

					for getItemPolicy, status := <-chPolicy; status; getItemPolicy, status = <-chPolicy {
						policy = getItemPolicy
						if status {
							break
						}
					}

					for getItemDromSize, status := <-chDromSize; status; getItemDromSize, status = <-chDromSize {
						dromSize = getItemDromSize
						if status {
							break
						}
					}

					for getItemWorkingGroup, status := <-chWorkingGroup; status; getItemWorkingGroup, status = <-chWorkingGroup {
						workingGroup = getItemWorkingGroup
						if status {
							break
						}
					}

					//This place is bigger. Control of messages from the channels is done here. If the channel has not yet received data, it will look at it again.
					//control:
					switch {
					// Here, array is written. After all the data is complete, the array is saved. Then '\ n' is passed to the next batch.
					//				This process is processed for all devices. The tabloda order to be created is as follows.
					//deviceCoding.DevicesID: Get the ID of this device.
					// applicationsStatus: This will retrieve the application information as previously set.
					// strConv.Itoa (dromSize): Changes the number of drom numbers to string type and returns the information inside.
					// Presence: brings the device 's online - offline information.
					// profile: gets active mode information.
					// policy: retrieves active policy information.
					// deviceCoding.Latitude: Get position information in latitude type.
					// deviceCoding.Longitude: Get position information in longitude type.
					// lastOnlineTime, time.Now (). String (): Writes the time the data is read. The longer the process, the longer it is given.
					// 				In some cases momentary time is very important. This gives the time of the recorded data.
					// workingGroup: The devices give us information in any group.
					case applicationsStatus == "":
						logc.ReportPrint("Application Status Get Nil Passing Device : ", deviceCoding.DeviceID)
						continue
					case presence == "":
						logc.ReportPrint("Presence Get Nil Passing Device : ", deviceCoding.DeviceID)
						continue
					case lastOnlineTime == "":
						logc.ReportPrint("Last Online Time Get Nil Passing Device : ", deviceCoding.DeviceID)
						continue
					case profile == "":
						logc.ReportPrint("Profile Get Nil Passing Device : ", deviceCoding.DeviceID)
						continue
					case policy == "":
						logc.ReportPrint("Policy Get Nil Passing Device : ", deviceCoding.DeviceID)
						continue
					case workingGroup == "":
						logc.ReportPrint("WorkingGroup Get Nil Passing Device : ", deviceCoding.DeviceID)
						continue
					default:
						d.writeCsvArray = append(d.writeCsvArray,
							deviceCoding.DeviceID,
							applicationsStatus,
							strconv.Itoa(dromSize),
							presence,
							profile,
							policy,
							deviceCoding.Latitude,
							deviceCoding.Longitude,
							lastOnlineTime, time.Now().String(),
							workingGroup,
							"\n")
					}

					// The display shows the sequence and the duration of the operation. Every 100 devices will give us information.
					if len(deviceCode.Extras) > 100 {
						if i%100 == 0 {
							logc.ReportPrint(i, "/", len(deviceCode.Extras))
						}
					} else if len(deviceCode.Extras) > 10 {
						if i%10 == 0 {
							logc.ReportPrint(i, "/", len(deviceCode.Extras))
						}
					}

				}
			}
		}
	} else {
		logc.ReportPrint("Querry Error ...")
	}

	//The data sent to the array sends the necessary information to the function that performs the write operation.
	//There is something that needs attention. All the data are collected and then sent for registration.
	//This is set in this way to gain speed. If you do not want to risk it, check out the other functions in the 'writefile' library.
	d.writeCSVType(fileName, d.writeCsvArray, packageHeader)
	logc.ReportPrint("Finish Read ...")
}

/*
 █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗    ███████╗████████╗ █████╗ ████████╗██╗   ██╗███████╗
██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║    ██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝██║   ██║██╔════╝
███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║    ███████╗   ██║   ███████║   ██║   ██║   ██║███████╗
██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║    ╚════██║   ██║   ██╔══██║   ██║   ██║   ██║╚════██║
██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║    ███████║   ██║   ██║  ██║   ██║   ╚██████╔╝███████║
╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝    ╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚══════╝
*/

// This function takes the external deviceID, setControlPackage, chApplicationsStatus data. These.
// deviceID: The number of the device to be controlled. String type. one is sent. So you can only look at one device at a time. The entire query is carried out using this ID number.
// setControlPackage: Contains packages to be checked. The String is of type Array. It checks all the packets if they are entered.
// chApplicationsStatus: Go creates a message channel for the routine
func (d DetailAllReport) applicationStatus(deviceID string, setControlPackage []string, chApplicationsStatus chan string) string {
	// This section specifies the space - fill state of the values that come into the function.
	// If this information is empty, the util performed in the function will fail.
	if deviceID != "" {
		if setControlPackage != nil {

			// statusGlobal: variable that contains the Runnig or NotRunning state
			// blockedControl: The variable that hosts the Blocked or NotBlocked state.
			// findControl: If the application in the loop matches, the variable is incremented. This will be told in the place where it is used.
			// notFindControl: If the application package name is not found in the loop, the variable is incremented by one. This will be told in the place where it is used.
			var (
				statusGlobal   string
				blockedControl string
				findControl    int
				notFindControl int
			)

			responseNilFunction := func() string {
				var applicationsNil string
				for j := 0; j < len(setControlPackage); j++ {
					if j == 0 {
						applicationsNil = rest.ResponseNil + "," + rest.ResponseNil
					} else {
						applicationsNil = rest.ResponseNil + "," + rest.ResponseNil + "," + applicationsNil
					}
				}
				strings.Trim(applicationsNil, " ")
				return applicationsNil
			}

			//Host the Running state and Blocked state of the referenced application Array
			packageStatus := make([]string, 0)

			//It makes the ApplicationInfo query with the given deviceID. This query returns all application data of the device backwards.
			appliacationQuery := d.devices.ApplicationInfo(deviceID, rest.NOMarshal, rest.Invisible)

			if appliacationQuery != nil {
				if string(appliacationQuery) != rest.ResponseNotFound {

					deviceApplication := device.ApplicationInfoJSON{}
					json.Unmarshal(appliacationQuery, &deviceApplication)

					// Looking at the situation will return to the desired application.
					for _, controlPackages := range setControlPackage {
						notFindControl = 0
						findControl = 0
						// Related device applications
						for _, downloadedApp := range deviceApplication.Data {
							// First check. Compare these applications with installed applications.
							if downloadedApp.PackageName == controlPackages {
								// Checked

								// the working status of the relevant application came up

								if downloadedApp.Running {
									statusGlobal = rest.Running
								} else if !downloadedApp.Running {
									statusGlobal = rest.NotRunning
								} else {
									statusGlobal = rest.ResponseNil
								}

								if downloadedApp.Blocked == -1 {
									blockedControl = rest.UnKnow
								} else if downloadedApp.Blocked == 0 {
									blockedControl = rest.NotBlocked
								} else if downloadedApp.Blocked == 1 {
									blockedControl = rest.Blocked
								} else {
									blockedControl = rest.ResponseNil
								}

								packageStatus = append(packageStatus, statusGlobal, blockedControl)
								findControl++
								break
							} else {
								notFindControl++
							}
						}
						// Absence of application
						if notFindControl > 0 && findControl == 0 {
							packageStatus = append(packageStatus, rest.NoApplication, rest.NoApplication)
						}
					}

					// If 404 Not Found is
				} else {
					for j := 0; j < len(setControlPackage); j++ {
						packageStatus = append(packageStatus, rest.ResponseNotFound, rest.ResponseNotFound)
					}
				}

				// If Nil is
			} else {
				for j := 0; j < len(setControlPackage); j++ {
					packageStatus = append(packageStatus, rest.ResponseNil, rest.ResponseNil)
				}
			}

			var applications string
			if packageStatus != nil {
				for i, applicationLo := range packageStatus {
					if applicationLo != "" {
						if i == 0 {
							applications = applicationLo
						} else {
							applications = applicationLo + "," + applications
						}
					} else {
						applications = responseNilFunction()
					}
				}
			} else {
				applications = responseNilFunction()
			}

			if applications == "" {
				applications = responseNilFunction()
			}

			chApplicationsStatus <- applications
			close(chApplicationsStatus)
			return applications
		}
		chApplicationsStatus <- rest.ResponseNil
		close(chApplicationsStatus)
		return rest.ResponseNil
	}

	chApplicationsStatus <- rest.ResponseNil
	close(chApplicationsStatus)

	return rest.ResponseNil
}

/*
██████╗ ██████╗ ███████╗███████╗███████╗███╗   ██╗ ██████╗███████╗    ███████╗████████╗ █████╗ ████████╗██╗   ██╗███████╗
██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝████╗  ██║██╔════╝██╔════╝    ██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝██║   ██║██╔════╝
██████╔╝██████╔╝█████╗  ███████╗█████╗  ██╔██╗ ██║██║     █████╗      ███████╗   ██║   ███████║   ██║   ██║   ██║███████╗
██╔═══╝ ██╔══██╗██╔══╝  ╚════██║██╔══╝  ██║╚██╗██║██║     ██╔══╝      ╚════██║   ██║   ██╔══██║   ██║   ██║   ██║╚════██║
██║     ██║  ██║███████╗███████║███████╗██║ ╚████║╚██████╗███████╗    ███████║   ██║   ██║  ██║   ██║   ╚██████╔╝███████║
╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═══╝ ╚═════╝╚══════╝    ╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚══════╝
*/

func (d DetailAllReport) presenceStatus(deviceCode string, chPresence, chLastOnlineTime chan string) (string, string) {
	// This section specifies the space - fill state of the values that come into the function.
	// If this information is empty, the util performed in the function will fail.
	if deviceCode != "" {
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

		if presence == "" {
			presence = rest.ResponseNil
		}

		if lastOnlineTime == "" {
			lastOnlineTime = rest.ResponseNil
		}

		chPresence <- presence
		chLastOnlineTime <- lastOnlineTime
		close(chPresence)
		close(chLastOnlineTime)
		return presence, lastOnlineTime
	}

	chPresence <- rest.ResponseNil
	chLastOnlineTime <- rest.ResponseNil

	close(chPresence)
	close(chLastOnlineTime)

	return rest.ResponseNil, rest.ResponseNil
}

/*
██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗    ██████╗  ██████╗ ██╗     ██╗ ██████╗██╗   ██╗
██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝    ██╔══██╗██╔═══██╗██║     ██║██╔════╝╚██╗ ██╔╝
██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗      ██████╔╝██║   ██║██║     ██║██║      ╚████╔╝
██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝      ██╔═══╝ ██║   ██║██║     ██║██║       ╚██╔╝
██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗    ██║     ╚██████╔╝███████╗██║╚██████╗   ██║
╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝    ╚═╝      ╚═════╝ ╚══════╝╚═╝ ╚═════╝   ╚═╝
*/

func (d DetailAllReport) profilePolicy(deviceID string, chProfile, chPolicy chan string) (string, string) {
	// This section specifies the space - fill state of the values that come into the function.
	// If this information is empty, the util performed in the function will fail.
	if deviceID != "" {
		var (
			activeProfile string
			activePolicy  string
		)

		profilePolicyQuery := d.devices.ActiveProfilePolicy(deviceID, rest.NOMarshal, rest.Invisible)
		if profilePolicyQuery != nil {
			if string(profilePolicyQuery) != rest.ResponseNotFound {
				activeProfilePolicy := device.ActiveProfilePolicyJSON{}
				json.Unmarshal(profilePolicyQuery, &activeProfilePolicy)

				activeProfile = activeProfilePolicy.CurrentProfile
				activePolicy = activeProfilePolicy.CurrentPolicy

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

		if activeProfile == "" {
			activeProfile = rest.ResponseNil
		}

		if activePolicy == "" {
			activePolicy = rest.ResponseNil
		}

		chProfile <- activeProfile
		chPolicy <- activePolicy

		close(chProfile)
		close(chPolicy)
		return activeProfile, activePolicy
	}

	chProfile <- rest.ResponseNil
	chPolicy <- rest.ResponseNil

	close(chProfile)
	close(chPolicy)

	return rest.ResponseNil, rest.ResponseNil
}

/*
███████╗██╗   ██╗██████╗ ███╗   ███╗██╗████████╗████████╗███████╗██████╗     ██████╗ ██████╗  ██████╗ ███╗   ███╗    ███████╗██╗███████╗███████╗
██╔════╝██║   ██║██╔══██╗████╗ ████║██║╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗    ██╔══██╗██╔══██╗██╔═══██╗████╗ ████║    ██╔════╝██║╚══███╔╝██╔════╝
███████╗██║   ██║██████╔╝██╔████╔██║██║   ██║      ██║   █████╗  ██║  ██║    ██║  ██║██████╔╝██║   ██║██╔████╔██║    ███████╗██║  ███╔╝ █████╗
╚════██║██║   ██║██╔══██╗██║╚██╔╝██║██║   ██║      ██║   ██╔══╝  ██║  ██║    ██║  ██║██╔══██╗██║   ██║██║╚██╔╝██║    ╚════██║██║ ███╔╝  ██╔══╝
███████║╚██████╔╝██████╔╝██║ ╚═╝ ██║██║   ██║      ██║   ███████╗██████╔╝    ██████╔╝██║  ██║╚██████╔╝██║ ╚═╝ ██║    ███████║██║███████╗███████╗
╚══════╝ ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝   ╚═╝      ╚═╝   ╚══════╝╚═════╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝    ╚══════╝╚═╝╚══════╝╚══════╝
*/

func (d DetailAllReport) submittedDromSize(deviceCode string, chDromSize chan int) int {
	// This section specifies the space - fill state of the values that come into the function.
	// If this information is empty, the util performed in the function will fail.
	if deviceCode != "" {
		var dromCounter int

		actionStatueQuery := d.actions.GetActionStatus(deviceCode, "PUSH_CMD_DROM", 1000, rest.Invisible)
		if actionStatueQuery != nil {
			if string(actionStatueQuery) != rest.ResponseNotFound {

				actionMessage := action.ResponseActionMessageJSON{}
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
	chDromSize <- -5
	close(chDromSize)
	return -5
}

/*
██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗██╗███╗   ██╗ ██████╗      ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗
██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝██║████╗  ██║██╔════╝     ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗
██║ █╗ ██║██║   ██║██████╔╝█████╔╝ ██║██╔██╗ ██║██║  ███╗    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝
██║███╗██║██║   ██║██╔══██╗██╔═██╗ ██║██║╚██╗██║██║   ██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝
╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗██║██║ ╚████║╚██████╔╝    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║
 ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝
*/

func (d DetailAllReport) workingGroup(deviceID string, chWorkingGroup chan string) string {
	// This section specifies the space - fill state of the values that come into the function.
	// If this information is empty, the util performed in the function will fail.
	if deviceID != "" {

		workingGroup := d.devices.WorkingGroupControl(deviceID, rest.Invisible)

		if workingGroup == "" {
			workingGroup = rest.ResponseNil
		}

		chWorkingGroup <- workingGroup
		close(chWorkingGroup)
		return workingGroup
	}
	chWorkingGroup <- rest.ResponseNil
	close(chWorkingGroup)
	return rest.ResponseNil
}

/*
██╗    ██╗██████╗ ██╗████████╗███████╗     ██████╗███████╗██╗   ██╗    ████████╗██╗   ██╗██████╗ ███████╗
██║    ██║██╔══██╗██║╚══██╔══╝██╔════╝    ██╔════╝██╔════╝██║   ██║    ╚══██╔══╝╚██╗ ██╔╝██╔══██╗██╔════╝
██║ █╗ ██║██████╔╝██║   ██║   █████╗      ██║     ███████╗██║   ██║       ██║    ╚████╔╝ ██████╔╝█████╗
██║███╗██║██╔══██╗██║   ██║   ██╔══╝      ██║     ╚════██║╚██╗ ██╔╝       ██║     ╚██╔╝  ██╔═══╝ ██╔══╝
╚███╔███╔╝██║  ██║██║   ██║   ███████╗    ╚██████╗███████║ ╚████╔╝        ██║      ██║   ██║     ███████╗
 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝   ╚═╝   ╚══════╝     ╚═════╝╚══════╝  ╚═══╝         ╚═╝      ╚═╝   ╚═╝     ╚══════╝
*/

func (d DetailAllReport) writeCSVType(fileName string, writeCSVArray []string, setControlPackage string) {
	var detailReportFile *os.File

	writefile.CreateFile(fileName)
	detailReportFile = writefile.OpenFile(detailReportFile, fileName)
	writefile.WriteArray(detailReportFile, writeCSVArray)
	logc.ReportPrint("Finish Write : ", setControlPackage)
}
