package reports

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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

// DetailAllReportReal is provides a detailed report by using REST APIs.
// For use Example:
//	var detail reports.DetailAllReportReal
//	detail.FileName = "DetailReport_" + timop.GetTimeNamesFormat() + ".csv"
//	detail.ControlAppPackage = []string{"package_name_1", "package_name_2"}
//
//	detail.DeviceList = []string{"0000", "1111", "2222"}
//	//or
//	detail.DevicesListFilePath = "REPORT_DEVICES.txt"
//	detail.Start()
type DetailAllReportReal struct {
	writeCsvArray       []string
	devices             device.Device
	actions             action.Action
	detailReportFile    *os.File
	FileName            string
	DeviceList          []string
	DevicesListFilePath string
	Unit                int
	ControlAppPackage   []string
}

//Start is DetailAllReport. These Cases were created to get detailed reports
func (d DetailAllReportReal) Start() {
	// It performs the writing process in one step, not in every step of the way. The goal is to increase
	// the speed and reduce the memory footprint. It is also used to write multiple files at the same time
	writefile.CreateFile(d.FileName)
	d.detailReportFile = writefile.OpenFile(d.detailReportFile, d.FileName)
	writefile.SplitCharacter = ";"

	// His section allows you to write them in the title if the number of applications to be checked is reached.
	// The reason for the substitution assignment is that it will be used later when writing csv
	var packageHeader string
	if d.ControlAppPackage != nil {
		for i, packageHeaderName := range d.ControlAppPackage {
			if i == 0 {
				packageHeader = packageHeaderName + "," + packageHeaderName + "," + packageHeaderName
			} else {
				packageHeader = packageHeaderName + "," + packageHeaderName + "," + packageHeaderName + "," + packageHeader
			}
		}
	} else {
		packageHeader = ""
	}

	strings.Trim(packageHeader, " ")
	// The title of the csv file to be made
	//d.writeCsvArray = append(d.writeCsvArray, ",,Observed Applications", "\n")
	writefile.WriteText(d.detailReportFile, ",,Observed Applications")

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
	//d.writeCsvArray = append(d.writeCsvArray, "Device ID", packageHeader, "Drom Count", "Presence", "Active Profile Name", "Active Policy Name", "Current Profile Name", "Policy Name", "Latitude", "Longitude", "Last Online Time", "Read Time", "Working Group", "\n")
	writefile.WriteText(d.detailReportFile, "Device ID", packageHeader, "Drom Count", "Presence", "Active Profile Name", "Active Policy Name", "Current Profile Name", "Policy Name", "Latitude", "Longitude", "Last Online Time", "Read Time", "Working Group", "Modiverse Version", "OS Version", "Report Time")

	var deviceList []string
	location := make(map[string]map[string]string)
	//TODO AÇIKLAMA EKLE

	if len(d.DevicesListFilePath) != 0 {
		file, err := os.Open(d.DevicesListFilePath)
		if err != nil {
			panic(d.DevicesListFilePath + " is not openned!!!")
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			deviceList = append(deviceList, scanner.Text())
			latitude, longitude := d.devices.LocationDevice(scanner.Text())

			location[scanner.Text()] = map[string]string{
				"Latitude":  latitude,
				"Longitude": longitude}
		}
	} else if d.DeviceList == nil {
		query := d.devices.LocationMap(rest.NOMarshal, rest.Invisible)
		if query != nil {
			if string(query) != rest.ResponseNotFound {
				// This assignment is aimed at resetting the variable
				deviceCode := device.LocationJSON{}
				json.Unmarshal(query, &deviceCode)

				for _, getDeviceID := range deviceCode.Extras {
					deviceList = append(deviceList, getDeviceID.DeviceID)

					location[getDeviceID.DeviceID] = map[string]string{
						"Latitude":  getDeviceID.Latitude,
						"Longitude": getDeviceID.Longitude}
				}
			}
		} //TODO DURUM KONTROL EKLE
	} else {
		//TODO AÇIKLAMA EKLE
		deviceList = append(deviceList, d.DeviceList...)
		for _, getDeviceID := range d.DeviceList {
			latitude, longitude := d.devices.LocationDevice(getDeviceID)

			location[getDeviceID] = map[string]string{
				"Latitude":  latitude,
				"Longitude": longitude}
		}
	}

	if d.Unit == 0 {
		if len(deviceList) > 10 {
			d.Unit = 10
		} else {
			d.Unit = 1
		}
	}

	var unit = d.Unit
	var threadValue = 0

	for i := 1; i < len(deviceList)+1; i++ {
		if i%unit == 0 {
			fmt.Println("Begin : ", i-unit, " End : ", i)
			threadValue++
			go d.controlReport(d.ControlAppPackage, packageHeader, location, threadValue, deviceList[i-unit:i]...)
		}
		if i%unit == 0 && i+unit > len(deviceList) {
			threadValue++
			go d.controlReport(d.ControlAppPackage, packageHeader, location, threadValue, deviceList[i:]...)
			fmt.Println("Begin : ", i, " End : ", len(deviceList))
		}
	}
}

func (d DetailAllReportReal) controlReport(packageList []string, packageHeader string, location map[string]map[string]string, threadValue int, deviceList ...string) {
	fmt.Println(deviceList)
	for i, deviceID := range deviceList {
		if deviceID != "" {

			var (
				applicationsStatus string
				presence           string
				lastOnlineTime     string
				profile            string
				policy             string
				dromSize           int
				workingGroup       string
				modiverseVersion   string
				osDisplay          string
				reportTime         string
				profileCurrent     string
				policyCurrent      string
				adminArea          string
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
			chProfileCurrent := make(chan string)
			chPolicyCurrent := make(chan string)
			chDromSize := make(chan int)
			chWorkingGroup := make(chan string)
			chModiverseVersion := make(chan string)
			chOSDisplay := make(chan string)
			chReportTime := make(chan string)
			chAdminArea := make(chan string)

			// Start GoRutines.
			// applicationStatus: Returns the application status and block status of applications that are initially given a package name.
			// presenceStatus: The status of the device is online - offline. If offline, it will tell you when it was last online.
			// profilePolicy: Provides the mode and policy information of the device.
			// submittedDromSize: How many DROMs are sent to the device.
			// workingGroup: Tells us if there is a group on the device.
			go d.applicationStatus(deviceID, packageList, chApplicationsStatus)
			go d.presenceStatus(deviceID, chPresence, chLastOnlineTime)
			go d.profilePolicy(deviceID, chProfile, chPolicy, chProfileCurrent, chPolicyCurrent)
			go d.submittedDromSize(deviceID, chDromSize)
			go d.workingGroup(deviceID, chWorkingGroup)
			go d.deviceInformation(deviceID, chModiverseVersion, chOSDisplay, chReportTime, chAdminArea)

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

			for getItemProfileCurrent, status := <-chProfileCurrent; status; getItemProfileCurrent, status = <-chProfileCurrent {
				profileCurrent = getItemProfileCurrent
				if status {
					break
				}
			}

			for getItemPolicyCurrent, status := <-chPolicyCurrent; status; getItemPolicyCurrent, status = <-chPolicyCurrent {
				policyCurrent = getItemPolicyCurrent
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

			for getItemModiverseVersion, status := <-chModiverseVersion; status; getItemModiverseVersion, status = <-chModiverseVersion {
				modiverseVersion = getItemModiverseVersion
				if status {
					break
				}
			}

			for getItemOSDisplay, status := <-chOSDisplay; status; getItemOSDisplay, status = <-chOSDisplay {
				osDisplay = getItemOSDisplay
				if status {
					break
				}
			}

			for getItemReportTime, status := <-chReportTime; status; getItemReportTime, status = <-chReportTime {
				reportTime = getItemReportTime
				if status {
					break
				}
			}

			for getItemAdminArea, status := <-chAdminArea; status; getItemAdminArea, status = <-chAdminArea {
				adminArea = getItemAdminArea
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
				logc.ReportPrint("Application Status Get Nil Passing Device : ", deviceID)
				continue
			case presence == "":
				logc.ReportPrint("Presence Get Nil Passing Device : ", deviceID)
				continue
			case lastOnlineTime == "":
				logc.ReportPrint("Last Online Time Get Nil Passing Device : ", deviceID)
				continue
			case profile == "":
				logc.ReportPrint("Profile Get Nil Passing Device : ", deviceID)
				continue
			case policy == "":
				logc.ReportPrint("Policy Get Nil Passing Device : ", deviceID)
				continue
			case profileCurrent == "":
				logc.ReportPrint("Profile Get Nil Passing Device : ", deviceID)
				continue
			case policyCurrent == "":
				logc.ReportPrint("Policy Get Nil Passing Device : ", deviceID)
				continue
			case workingGroup == "":
				logc.ReportPrint("WorkingGroup Get Nil Passing Device : ", deviceID)
				continue
			case modiverseVersion == "":
				logc.ReportPrint("ModiverseVersion Get Nil Passing Device : ", deviceID)
				continue
			case osDisplay == "":
				logc.ReportPrint("OsDisplay Get Nil Passing Device : ", deviceID)
				continue
			case reportTime == "":
				logc.ReportPrint("ReportTime Get Nil Passing Device : ", deviceID)
				continue
			case adminArea == "":
				logc.ReportPrint("ReportTime Get Nil Passing Device : ", deviceID)
				continue
			default:
				d.writeCsvArray = append(d.writeCsvArray,
					deviceID,
					applicationsStatus,
					strconv.Itoa(dromSize),
					presence,
					profile,
					policy,
					profileCurrent,
					policyCurrent,
					location[deviceID]["Latitude"],
					location[deviceID]["Longitude"],
					lastOnlineTime,
					time.Now().String(),
					workingGroup,
					modiverseVersion,
					osDisplay,
					reportTime+
						"\n")
			}

			// The display shows the sequence and the duration of the operation. Every 100 devices will give us information.
			deviceListLen := len(deviceList)
			if deviceListLen > 100 {
				if i%100 == 0 {
					fmt.Println(threadValue, " : ", i, "/", deviceListLen)
				}
			} else if deviceListLen > 10 {
				if i%10 == 0 {
					fmt.Println(threadValue, " : ", i, "/", deviceListLen)
				}
			} else if deviceListLen <= 10 {
				fmt.Println(threadValue, " : ", i+1, "/", deviceListLen)
			}
		}
	}

	//The data sent to the array sends the necessary information to the function that performs the write operation.
	//There is something that needs attention. All the data are collected and then sent for registration.
	//This is set in this way to gain speed. If you do not want to risk it, check out the other functions in the 'writefile' library.
	d.writeCSVType(d.writeCsvArray, packageHeader)
	log.Println("Finish Read ...")
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
func (d DetailAllReportReal) applicationStatus(deviceID string, setControlPackage []string, chApplicationsStatus chan string) string {
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
				versionCode    string
				versionCodeStr string
				findControl    int
				notFindControl int
			)

			responseNilFunction := func() string {
				var applicationsNil string
				for j := 0; j < len(setControlPackage); j++ {
					if j == 0 {
						applicationsNil = rest.ResponseNil + "," + rest.ResponseNil + "," + rest.ResponseNil
					} else {
						applicationsNil = rest.ResponseNil + "," + rest.ResponseNil + "," + applicationsNil + "," + applicationsNil
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

								versionCode = downloadedApp.VersionName

								if len(versionCode) <= 0 {
									versionCodeStr = rest.ResponseNil
								} else {
									versionCodeStr = versionCode
								}

								packageStatus = append(packageStatus, statusGlobal, blockedControl, versionCodeStr)
								findControl++
								break
							} else {
								notFindControl++
							}
						}
						// Absence of application
						if notFindControl > 0 && findControl == 0 {
							packageStatus = append(packageStatus, rest.NoApplication, rest.NoApplication, rest.NoApplication)
						}
					}

					// If 404 Not Found is
				} else {
					for j := 0; j < len(setControlPackage); j++ {
						packageStatus = append(packageStatus, rest.ResponseNotFound, rest.ResponseNotFound, rest.ResponseNotFound)
					}
				}

				// If Nil is
			} else {
				for j := 0; j < len(setControlPackage); j++ {
					packageStatus = append(packageStatus, rest.ResponseNil, rest.ResponseNil, rest.ResponseNil)
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

func (d DetailAllReportReal) presenceStatus(deviceCode string, chPresence, chLastOnlineTime chan string) (string, string) {
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
func (d DetailAllReportReal) profilePolicy(deviceID string, chProfile, chPolicy, chProfileCurrent, chPolicyCurrent chan string) (string, string) {
	// This section specifies the space - fill state of the values that come into the function.
	// If this information is empty, the util performed in the function will fail.
	if deviceID != "" {
		var (
			activeProfile string
			activePolicy  string

			currentProfile string
			currentPolicy  string
		)

		profilePolicyQuery := d.devices.ActiveProfilePolicy(deviceID, rest.NOMarshal, rest.Invisible)
		if profilePolicyQuery != nil {
			if string(profilePolicyQuery) != rest.ResponseNotFound {
				activeProfilePolicy := device.ActiveProfilePolicyJSON{}
				json.Unmarshal(profilePolicyQuery, &activeProfilePolicy)

				activeProfile = activeProfilePolicy.ActiveProfile
				activePolicy = activeProfilePolicy.ActivePolicy

				currentProfile = activeProfilePolicy.CurrentProfile
				currentPolicy = activeProfilePolicy.CurrentPolicy

				if len(activeProfile) == 0 {
					activeProfile = rest.ResponseNil
				}
				if len(activePolicy) == 0 {
					activePolicy = rest.ResponseNil
				}

				if len(currentProfile) == 0 {
					currentProfile = rest.ResponseNil
				}
				if len(currentPolicy) == 0 {
					currentPolicy = rest.ResponseNil
				}

			} else {
				activeProfile = rest.ResponseNotFound
				activePolicy = rest.ResponseNotFound

				currentProfile = rest.ResponseNotFound
				currentPolicy = rest.ResponseNotFound
			}
		} else {
			activeProfile = rest.ResponseNil
			activePolicy = rest.ResponseNil

			currentProfile = rest.ResponseNil
			currentPolicy = rest.ResponseNil
		}

		if activeProfile == "" {
			activeProfile = rest.ResponseNil
		}

		if activePolicy == "" {
			activePolicy = rest.ResponseNil
		}

		if currentProfile == "" {
			currentProfile = rest.ResponseNil
		}

		if currentPolicy == "" {
			currentPolicy = rest.ResponseNil
		}

		chProfile <- activeProfile
		chPolicy <- activePolicy

		chProfileCurrent <- currentProfile
		chPolicyCurrent <- currentPolicy

		close(chProfile)
		close(chPolicy)

		close(chProfileCurrent)
		close(chPolicyCurrent)

		return activeProfile, activePolicy
	}

	chProfile <- rest.ResponseNil
	chPolicy <- rest.ResponseNil

	chProfileCurrent <- rest.ResponseNil
	chPolicyCurrent <- rest.ResponseNil

	close(chProfile)
	close(chPolicy)

	close(chProfileCurrent)
	close(chPolicyCurrent)

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

func (d DetailAllReportReal) submittedDromSize(deviceCode string, chDromSize chan int) int {
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
██████╗ ███████╗██╗   ██╗██╗ ██████╗███████╗        ██╗███╗   ██╗███████╗ ██████╗ ██████╗ ███╗   ███╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗
██╔══██╗██╔════╝██║   ██║██║██╔════╝██╔════╝        ██║████╗  ██║██╔════╝██╔═══██╗██╔══██╗████╗ ████║██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║
██║  ██║█████╗  ██║   ██║██║██║     █████╗          ██║██╔██╗ ██║█████╗  ██║   ██║██████╔╝██╔████╔██║███████║   ██║   ██║██║   ██║██╔██╗ ██║
██║  ██║██╔══╝  ╚██╗ ██╔╝██║██║     ██╔══╝          ██║██║╚██╗██║██╔══╝  ██║   ██║██╔══██╗██║╚██╔╝██║██╔══██║   ██║   ██║██║   ██║██║╚██╗██║
██████╔╝███████╗ ╚████╔╝ ██║╚██████╗███████╗        ██║██║ ╚████║██║     ╚██████╔╝██║  ██║██║ ╚═╝ ██║██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║
╚═════╝ ╚══════╝  ╚═══╝  ╚═╝ ╚═════╝╚══════╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
*/
func (d DetailAllReportReal) deviceInformation(deviceID string, chModiverseVersion chan string, chOSDisplay chan string, chReportTime chan string, chAdminArea chan string) (string, string, string, string) {
	var (
		devices          device.Device
		osDisplay        string
		modiverseVersion string
		reportTime       string
		adminArea        string
	)

	queryInformation := devices.DeviceInformation(devices.DeviceID2Code(deviceID), rest.NOMarshal, rest.Invisible)

	if queryInformation != nil {
		if string(queryInformation) != rest.ResponseNotFound {
			deviceInformation := device.InformationJSON{}
			json.Unmarshal(queryInformation, &deviceInformation)

			osDisplay = deviceInformation.OsProfile.Display
			modiverseVersion = deviceInformation.ModeAppVersion
			reportTime = deviceInformation.DeviceCurrentTime
			adminArea = deviceInformation.AdminArea.Name

			if osDisplay == "" {
				osDisplay = rest.ResponseNil
			}

			if modiverseVersion == "" {
				modiverseVersion = rest.ResponseNil
			}

			if reportTime == "" {
				reportTime = rest.ResponseNil
			}

			if adminArea == "" {
				adminArea = rest.ResponseNil
			}

			chModiverseVersion <- modiverseVersion
			close(chModiverseVersion)

			chOSDisplay <- osDisplay
			close(chOSDisplay)

			chReportTime <- reportTime
			close(chReportTime)

			chAdminArea <- adminArea
			close(chAdminArea)

			return osDisplay, modiverseVersion, reportTime, adminArea
		}

		chModiverseVersion <- rest.ResponseNotFound
		close(chModiverseVersion)

		chOSDisplay <- rest.ResponseNotFound
		close(chOSDisplay)

		chReportTime <- rest.ResponseNotFound
		close(chReportTime)
		return rest.ResponseNotFound, rest.ResponseNotFound, rest.ResponseNotFound, rest.ResponseNotFound
	}
	chModiverseVersion <- rest.ResponseNil
	close(chModiverseVersion)

	chOSDisplay <- rest.ResponseNil
	close(chOSDisplay)

	chReportTime <- rest.ResponseNil
	close(chReportTime)
	return rest.ResponseNil, rest.ResponseNil, rest.ResponseNil, rest.ResponseNotFound
}

/*
██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗██╗███╗   ██╗ ██████╗      ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗
██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝██║████╗  ██║██╔════╝     ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗
██║ █╗ ██║██║   ██║██████╔╝█████╔╝ ██║██╔██╗ ██║██║  ███╗    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝
██║███╗██║██║   ██║██╔══██╗██╔═██╗ ██║██║╚██╗██║██║   ██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝
╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗██║██║ ╚████║╚██████╔╝    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║
 ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝
*/

func (d DetailAllReportReal) workingGroup(deviceID string, chWorkingGroup chan string) string {
	// This section specifies the space - fill state of the values that come into the function.
	// If this information is empty, the util performed in the function will fail.
	if deviceID != "" {

		workingGroup := d.devices.WorkingGroupControl(deviceID, rest.Invisible)

		if workingGroup == "" {
			workingGroup = rest.ResponseNil
		}
		workingGroup = strings.Replace(workingGroup, ",", ":", -1)
		workingGroup = strings.Replace(workingGroup, ";", ":", -1)

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

func (d DetailAllReportReal) writeCSVType(writeCSVArray []string, setControlPackage string) {

	d.detailReportFile = writefile.OpenFile(d.detailReportFile, d.FileName)
	writefile.WriteArray(d.detailReportFile, writeCSVArray)
	logc.ReportPrint("Finish Write : ", setControlPackage)
}
