package operations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/KMACEL/IITR/databasecenter"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/workingset"
	"github.com/KMACEL/IITR/writefile"
)

/*
 ██████╗ ████████╗ █████╗         ██████╗  █████╗ ███████╗███████╗
██╔═══██╗╚══██╔══╝██╔══██╗        ██╔══██╗██╔══██╗██╔════╝██╔════╝
██║   ██║   ██║   ███████║        ██████╔╝███████║███████╗█████╗
██║   ██║   ██║   ██╔══██║        ██╔══██╗██╔══██║╚════██║██╔══╝
╚██████╔╝   ██║   ██║  ██║        ██████╔╝██║  ██║███████║███████╗
 ╚═════╝    ╚═╝   ╚═╝  ╚═╝        ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝
 */

/*
███████╗████████╗██████╗ ██╗   ██╗ ██████╗████████╗
██╔════╝╚══██╔══╝██╔══██╗██║   ██║██╔════╝╚══██╔══╝
███████╗   ██║   ██████╔╝██║   ██║██║        ██║
╚════██║   ██║   ██╔══██╗██║   ██║██║        ██║
███████║   ██║   ██║  ██║╚██████╔╝╚██████╗   ██║
╚══════╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝  ╚═════╝   ╚═╝
*/

// OtaOp builds a roof for automobile operations. The data found here is very important for OTA operation.
// With this structure, important data for OTA functions are communicated in a common way.
//     dataBase                  : The common * sql.DB structure to be used in database operations.
//     tableName                 : is the toblo name for OTA.
//     databaseName              : The database created for OTA is unknown.
//     otaUpdaterApplicationCode : this value contains the application's unique id. The application submission process is performed with this ID
//     otaUpdaterPacketName      : OTA hosting variable hosting application.
//     otaUpdaterVersion         : The version of OTA update application.
//     otaUpdaterVersionCode     : The version code of OTA update application.
//     osDisplay                 : The name of the firmware to be migrated.
type OtaOp struct {
	dataBase                  databasecenter.DB
	tableName                 string
	databaseName              string
	otaUpdaterApplicationCode string
	otaUpdaterPacketName      string
	otaUpdaterVersion         string
	otaUpdaterVersionCode     float32
	osDisplay                 string
}

// OtaBaseOp contains the structure types associated with the data base of the device to be built.
// The information here can be entered by the user. At the same time, the devices operate on this information.
// Note: The numbers 0 - 1 - 2 in the following values have the following meanings;
//     0: Not Running,
//     1: Running,
//     2: Finish
// The values taken in this structure are as follows;
//     Imei             : The unique number of the device. With this information, the device performs the rest queries.
//     PushApplication  : Takes values from 0 to 1 - 2. The device performs a data check of the data.
//     PushTime         : Shows the time of the last operation in the PushApplication section.
//     StartApplication : Takes values from 0 to 1 - 2. It checks the status of the application that is running on the device.
//     StartTime        : Indicates the time of the last operation in StartApplication.
//     ControlOta       : takes values from 0 to 1 - 2. The device performs fmrware sine status check.
//     ControlTime       : Shows the time of the last operation in ControlOta.
type OtaBaseOp struct {
	Imei             string
	PushApplication  int
	PushTime         time.Time
	StartApplication int
	StartTime        time.Time
	ControlOta       int
	ControlTime      time.Time
}

// OtaDeviceArray is the type of structure used to retrieve many OtaBaseOp structures.
// The start function waits for this type. Transactions occur with this type.
// You need to enter an OtaBaseOp with at least the Imei number into this loop.
// Example 	Base:
//     otaDeviceArray := operations.OtaDeviceArray{
//          operations.OtaBaseOp{Imei: "867377020740787"},
//          operations.OtaBaseOp{Imei: "867377020747089"}}
type OtaDeviceArray []OtaBaseOp

/*
 ██████╗ ██████╗ ███╗   ██╗███████╗████████╗
██╔════╝██╔═══██╗████╗  ██║██╔════╝╚══██╔══╝
██║     ██║   ██║██╔██╗ ██║███████╗   ██║
██║     ██║   ██║██║╚██╗██║╚════██║   ██║
╚██████╗╚██████╔╝██║ ╚████║███████║   ██║
 ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝   ╚═╝
 */

// This constant contains the database columns. With this information, strings used in all database
// operations are taken from a common center.
const (
	imei             = "imei"
	pushApplication  = "pushApplication"
	pushTime         = "pushTime"
	startApplication = "startApplication"
	startTime        = "startTime"
	controlOta       = "controlOta"
	controlTime      = "controlTime"
)

// This constant contains database variable types.
const (
	imeiType             = "varchar(15) PRIMARY KEY"
	pushApplicationType  = "INT"
	pushTimeType         = "datetime"
	startApplicationType = "INT"
	startTimeType        = "datetime"
	controlOtaType       = "INT"
	controlTimeType      = "datetime"
)

// this constant contains the numbers 0 - 1 - 2 processed in the database.
// Through these numbers, it is known what the number of the step is in common.
const (
	notRunningStep = 0
	runningStep    = 1
	finishStep     = 2
)

/*
███████╗████████╗ █████╗ ██████╗ ████████╗
██╔════╝╚══██╔══╝██╔══██╗██╔══██╗╚══██╔══╝
███████╗   ██║   ███████║██████╔╝   ██║
╚════██║   ██║   ██╔══██║██╔══██╗   ██║
███████║   ██║   ██║  ██║██║  ██║   ██║
╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝
 */

// It is the point where the start ota application starts.
// It receives data of type OtaDeviceArray which is at least the Imei number from the outside.
func (o OtaOp) Start(otaDevices OtaDeviceArray) {
	o.tableName = "otaOperations"
	o.databaseName = "iTaksi.db"

	// "071503D1-C864-4236-ABF0-DEA26550AF93" Ota Updater Asıl
	o.otaUpdaterApplicationCode = "DD76AFEA-E0A3-4B61-97CA-509B66A884E1"
	o.otaUpdaterPacketName = "com.estoty.game2048"
	o.otaUpdaterVersion = "6.05"
	o.otaUpdaterVersionCode = 46

	o.osDisplay = "rkpx2-eng 4.4.4 KTU84Q eng.turkey.20171019.131121 test-keys"

	db := o.dataBase.Open(o.databaseName)
	defer o.dataBase.Close(db)

	//o.refreshGatewayInfo(otaDevices)
	// todo 15 dk time delay
	if o.deviceControl(db, otaDevices) {
		o.pushApplicationOperation(db, otaDevices)
		o.startApplicationOperation(db, otaDevices)
		o.otaControlOperation(db, otaDevices)
	}

	o.reportDatabase()

}

/*
██████╗ ███████╗███████╗██████╗ ███████╗███████╗██╗  ██╗         ██████╗  █████╗ ████████╗███████╗██╗    ██╗ █████╗ ██╗   ██╗        ██╗███╗   ██╗███████╗ ██████╗
██╔══██╗██╔════╝██╔════╝██╔══██╗██╔════╝██╔════╝██║  ██║        ██╔════╝ ██╔══██╗╚══██╔══╝██╔════╝██║    ██║██╔══██╗╚██╗ ██╔╝        ██║████╗  ██║██╔════╝██╔═══██╗
██████╔╝█████╗  █████╗  ██████╔╝█████╗  ███████╗███████║        ██║  ███╗███████║   ██║   █████╗  ██║ █╗ ██║███████║ ╚████╔╝         ██║██╔██╗ ██║█████╗  ██║   ██║
██╔══██╗██╔══╝  ██╔══╝  ██╔══██╗██╔══╝  ╚════██║██╔══██║        ██║   ██║██╔══██║   ██║   ██╔══╝  ██║███╗██║██╔══██║  ╚██╔╝          ██║██║╚██╗██║██╔══╝  ██║   ██║
██║  ██║███████╗██║     ██║  ██║███████╗███████║██║  ██║        ╚██████╔╝██║  ██║   ██║   ███████╗╚███╔███╔╝██║  ██║   ██║           ██║██║ ╚████║██║     ╚██████╔╝
╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝         ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝ ╚══╝╚══╝ ╚═╝  ╚═╝   ╚═╝           ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
 */

// The data we receive with the Rest API may not always be up-to-date.
// This data is automatically updated at certain times.
// refreshGatewayInfo performs the request of the data from the device.
// This process loads the current data. This command is recommended to wait 7 minutes after you have worked.
// The collection, processing and updating of the data may take longer depending on the slowness of the internet.
// This should be a little bit of a wait.
func (o OtaOp) refreshGatewayInfo(otaDevices OtaDeviceArray) {
	var (
		devices device.Device
	)
	for _, devicesID := range otaDevices {
		devices.RefreshGatewayInfo(devices.DeviceID2Code(devicesID.Imei))
	}
}

/*
██████╗ ██╗   ██╗███████╗██╗  ██╗         █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗
██╔══██╗██║   ██║██╔════╝██║  ██║        ██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║
██████╔╝██║   ██║███████╗███████║        ███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║
██╔═══╝ ██║   ██║╚════██║██╔══██║        ██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║
██║     ╚██████╔╝███████║██║  ██║        ██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║
╚═╝      ╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
 */

// pushApplicationOperation acts as the submitter to the application.
// It examines the data in the list of supplied devices and the data in the database.
// In this case, if the application has not gone to the device, it calls the send function.
// If there is no status information that is gone but loaded, it checks. If so, it will report it.
func (o OtaOp) pushApplicationOperation(db *sql.DB, ota OtaDeviceArray) {
	var (
		getDevice        map[string]interface{}
		getCurrentNumber int64
	)
	for _, otaDevice := range ota {
		getDevice = o.dataBase.Find(db, o.tableName, imei, otaDevice.Imei)
		if getDevice != nil {
			getCurrentNumber = getDevice[pushApplication].(int64)
			if getCurrentNumber == notRunningStep {
				fmt.Println(otaDevice.Imei, " : 0")
				o.pushApplication(db, otaDevice.Imei)
			} else if getCurrentNumber == runningStep {
				fmt.Println(otaDevice.Imei, " : 1")
				o.controlPushApplication(db, otaDevice.Imei)
			} else if getCurrentNumber == finishStep {
				fmt.Println(otaDevice.Imei, " : 2")
			} else {
				fmt.Println(otaDevice.Imei, " : Err")
			}
		}
	}

}

// The pushApplication function performs the application sending operation
// to the device with the properties reported in the beginning
// if the application has no application and the corresponding field in the database is "0".
func (o OtaOp) pushApplication(db *sql.DB, otaDevice string) {
	var (
		workingsets workingset.Workingset
	)

	if workingsets.PushApplications(o.otaUpdaterApplicationCode, false, otaDevice) {
		o.dataBase.Update(db, o.tableName, pushApplication, "1", imei, otaDevice)
		o.dataBase.Update(db, o.tableName, pushTime, time.Now().String(), imei, otaDevice)
		fmt.Println("Updating...")
	}
}

// The controlPushApplication function checks if the "pushApplication" value is 1, that is,
// if the application is sent, and if it is not loaded. If the application has been installed,
// it will increase this value to "2" in the database.
func (o OtaOp) controlPushApplication(db *sql.DB, otaDevice string) {
	var (
		devices         device.Device
		infoApplication device.InstantApplicationInfoJSON
	)
	getAppInfo := devices.InstantApplicationInfo(otaDevice, rest.NOMarshal, rest.Invisible)
	//todo gelen değer kontrol

	json.Unmarshal(getAppInfo, &infoApplication)

	if infoApplication.Data.PackageName == o.otaUpdaterPacketName {
		if infoApplication.Data.VersionCode == o.otaUpdaterVersionCode {
			fmt.Println("Update : OK")
			o.dataBase.Update(db, o.tableName, pushApplication, "2", imei, otaDevice)
			o.dataBase.Update(db, o.tableName, pushTime, time.Now().String(), imei, otaDevice)
			fmt.Println("Data Base Updateing...")
		}
	}
}

/*
███████╗████████╗ █████╗ ██████╗ ████████╗         █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗
██╔════╝╚══██╔══╝██╔══██╗██╔══██╗╚══██╔══╝        ██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║
███████╗   ██║   ███████║██████╔╝   ██║           ███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║
╚════██║   ██║   ██╔══██║██╔══██╗   ██║           ██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║
███████║   ██║   ██║  ██║██║  ██║   ██║           ██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║
╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝           ╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
 */

// startApplicationOperation is used to run the application. If the application is installed, run it and check the status
func (o OtaOp) startApplicationOperation(db *sql.DB, ota OtaDeviceArray) {
	var (
		getDevicePushApplication map[string]interface{}
		getPushAppCurrentNumber  int64

		getDeviceStartApp        map[string]interface{}
		getStartAppCurrentNumber int64
	)
	fmt.Println("Start app")

	for _, otaDevice := range ota {
		getDevicePushApplication = o.dataBase.Find(db, o.tableName, imei, otaDevice.Imei)
		if getDevicePushApplication != nil {
			getPushAppCurrentNumber = getDevicePushApplication[pushApplication].(int64)
			if getPushAppCurrentNumber == finishStep { //todo ???? kontrol
				getDeviceStartApp = o.dataBase.Find(db, o.tableName, imei, otaDevice.Imei)
				if getDevicePushApplication != nil {
					getStartAppCurrentNumber = getDeviceStartApp["startApplication"].(int64)
					if getStartAppCurrentNumber == notRunningStep {
						o.startApplication(db, otaDevice.Imei)
					} else if getStartAppCurrentNumber == runningStep {
						o.controlStartApplication(db, otaDevice.Imei)
					} else if getStartAppCurrentNumber == finishStep {
						fmt.Println(otaDevice.Imei, " : is Starting App Now")
					}
				}

			} else {
				fmt.Println(otaDevice.Imei, " : Err Last OPerationNot Fİnish")
			}
		}
	}
}

func (o OtaOp) startApplication(db *sql.DB, otaDevice string) {
	var (
		devices device.Device
	)
	fmt.Println("Update Start App")
	devices.AppSS(device.StartApp, devices.DeviceID2Code(otaDevice), o.otaUpdaterPacketName, rest.Visible)
	o.dataBase.Update(db, o.tableName, startApplication, "1", imei, otaDevice)
	o.dataBase.Update(db, o.tableName, startTime, time.Now().String(), imei, otaDevice)
	//todo kontrol
}

func (o OtaOp) controlStartApplication(db *sql.DB, otaDevice string) {
	//todo kontrol
	o.dataBase.Update(db, o.tableName, startApplication, "2", imei, otaDevice)
	o.dataBase.Update(db, o.tableName, startTime, time.Now().String(), imei, otaDevice)
}

/*
 ██████╗ ████████╗ █████╗          ██████╗ ██████╗ ███╗   ██╗████████╗██████╗  ██████╗ ██╗
██╔═══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██╔═══██╗██║
██║   ██║   ██║   ███████║        ██║     ██║   ██║██╔██╗ ██║   ██║   ██████╔╝██║   ██║██║
██║   ██║   ██║   ██╔══██║        ██║     ██║   ██║██║╚██╗██║   ██║   ██╔══██╗██║   ██║██║
╚██████╔╝   ██║   ██║  ██║        ╚██████╗╚██████╔╝██║ ╚████║   ██║   ██║  ██║╚██████╔╝███████╗
 ╚═════╝    ╚═╝   ╚═╝  ╚═╝         ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝
 */

func (o OtaOp) otaControlOperation(db *sql.DB, ota OtaDeviceArray) {
	var (
		getDevicePushApplication map[string]interface{}
		getPushAppCurrentNumber  int64

		getDeviceStartApp        map[string]interface{}
		getStartAppCurrentNumber int64

		getDeviceOtaControl map[string]interface{}
		getOtaCurrentNumber int64
	)
	fmt.Println("OtaKOntrol")

	for _, otaDevice := range ota {
		getDevicePushApplication = o.dataBase.Find(db, o.tableName, imei, otaDevice.Imei)
		if getDevicePushApplication != nil {
			getPushAppCurrentNumber = getDevicePushApplication[pushApplication].(int64)
			if getPushAppCurrentNumber == finishStep { //todo ???? kontrol
				getDeviceStartApp = o.dataBase.Find(db, o.tableName, imei, otaDevice.Imei)
				if getDevicePushApplication != nil {
					getStartAppCurrentNumber = getDeviceStartApp[startApplication].(int64)
					if getStartAppCurrentNumber == finishStep {

						getDeviceOtaControl = o.dataBase.Find(db, o.tableName, imei, otaDevice.Imei)
						if getDeviceOtaControl != nil {
							getOtaCurrentNumber = getDeviceOtaControl[controlOta].(int64)
							if getOtaCurrentNumber == notRunningStep {
								o.controlOtaUpdate(db, otaDevice.Imei)
							}
						}
					}
				}

			} else {
				fmt.Println(otaDevice.Imei, " : Err Last OPerationNot Fİnish")
			}
		}
	}
}

func (o OtaOp) controlOtaUpdate(db *sql.DB, otaDevice string) {
	var (
		devices           device.Device
		deviceInformation device.InformationJSON
	)
	fmt.Println("Update Ota KOntrol")
	deviceInfo := devices.Informations(devices.DeviceID2Code(otaDevice), rest.NOMarshal, rest.Invisible)

	json.Unmarshal(deviceInfo, &deviceInformation)

	if deviceInformation.OsProfile.Display == o.osDisplay {
		fmt.Println("huhuhu upgraded")
		o.dataBase.Update(db, o.tableName, controlOta, "2", imei, otaDevice)
		o.dataBase.Update(db, o.tableName, controlTime, time.Now().String(), imei, otaDevice)
	} else {
		fmt.Println("I m sory upgraded")
		o.dataBase.Update(db, o.tableName, controlOta, "1", imei, otaDevice)
		o.dataBase.Update(db, o.tableName, controlTime, time.Now().String(), imei, otaDevice)
	}
}

/*
██████╗  █████╗ ████████╗ █████╗ ██████╗  █████╗ ███████╗███████╗
██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝
██║  ██║███████║   ██║   ███████║██████╔╝███████║███████╗█████╗
██║  ██║██╔══██║   ██║   ██╔══██║██╔══██╗██╔══██║╚════██║██╔══╝
██████╔╝██║  ██║   ██║   ██║  ██║██████╔╝██║  ██║███████║███████╗
╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝
 */

func (o OtaOp) createDatabse(db *sql.DB) {
	var dataBase databasecenter.DB
	dataBase.CreateTable(db, "otaOperations",
		"imei varchar(15) PRIMARY KEY ,"+
			"pushApplication INT,"+
			"pushTime datetime,"+
			"startApplication INT,"+
			"startTime datetime,"+
			"controlOta int,"+
			"controlTime datetime")
}

func (o OtaOp) deviceControl(db *sql.DB, ota OtaDeviceArray) bool {
	var getMessage map[string]interface{}
	for _, devices := range ota {
		fmt.Print(devices.Imei, " : ")
		getMessage = o.dataBase.Find(db, o.tableName, "imei", devices.Imei)
		if getMessage == nil {
			fmt.Println("Not Find Device. Inserting...")
			o.insertDataBase(devices)
		} else {
			fmt.Println("Find Device")
		}
	}
	return true
}

func (o OtaOp) insertDataBase(otaDevice OtaBaseOp) {
	var dataBase databasecenter.DB
	db := dataBase.Open(o.databaseName)

	dataBase.InsertInto(db,
		o.tableName,
		imei+","+pushApplication+","+pushTime+","+startApplication+","+startTime+","+controlOta+","+controlTime,
		otaDevice.Imei,
		otaDevice.PushApplication,
		otaDevice.PushTime, otaDevice.StartApplication,
		otaDevice.StartTime, otaDevice.ControlOta,
		otaDevice.ControlTime)
	/*"867377020740791", 2, time.Now(), 2, time.Now(), 1, time.Now()*/
}

func (o OtaOp) reportDatabase() {
	var (
		file *os.File
		dataBaseReport databasecenter.DB
	)

	const(
		otaFileName ="OtaReport.xlsx"
	)

	writefile.CreateFile(otaFileName)
	file = writefile.OpenFile(otaFileName, file)

	// Column Header
	writefile.WriteText(file, imei, pushApplication, pushTime, startApplication, startTime, controlOta, controlTime)

	db := dataBaseReport.Open(o.databaseName)
	getDatabase := dataBaseReport.Select(db, o.tableName)

	for i := range getDatabase {
		writefile.WriteText(file,
			string(getDatabase[i][imei].([]byte)),

			strconv.FormatInt(getDatabase[i][pushApplication].(int64), 10),
			getDatabase[i][pushTime].(time.Time).String(),

			strconv.FormatInt(getDatabase[i][startApplication].(int64), 10),
			getDatabase[i][startTime].(time.Time).String(),

			strconv.FormatInt(getDatabase[i][controlOta].(int64), 10),
			getDatabase[i][controlTime].(time.Time).String())
	}
}
