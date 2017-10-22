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

// OtaOp is
type OtaOp struct {
	dataBase              databasecenter.DB
	tableName             string
	databaseName          string
	otaUpdaterPacketName  string
	otaUpdaterVersion     string
	otaUpdaterVersionCode float32
	osDisplay             string
}

// OtaBaseOp is
type OtaBaseOp struct {
	Imei             string
	PushApplication  int
	PushTime         time.Time
	StartApplication int
	StartTime        time.Time
	ControlOta       int
	ControlTime      time.Time
}

// OtaDeviceArray is
type OtaDeviceArray []OtaBaseOp

const (
	imei             = "imei"
	pushApplication  = "pushApplication"
	pushTime         = "pushTime"
	startApplication = "startApplication"
	startTime        = "startTime"
	controlOta       = "controlOta"
	controlTime      = "controlTime"
)

const (
	imeiType             = "varchar(15) PRIMARY KEY"
	pushApplicationType  = "INT"
	pushTimeType         = "datetime"
	startApplicationType = "INT"
	startTimeType        = "datetime"
	controlOtaType       = "INT"
	controlTimeType      = "datetime"
)

const (
	notRunningStep = 0
	runningStep    = 1
	finishStep     = 2
)

// Start is
func (o OtaOp) Start(otaDevices OtaDeviceArray) {
	o.tableName = "otaOperations"
	o.databaseName = "iTaksi.db"

	o.otaUpdaterPacketName = "com.estoty.game2048"
	o.otaUpdaterVersion = "6.05"
	o.otaUpdaterVersionCode = 46

	o.osDisplay = "rkpx2-eng 4.4.4 KTU84Q eng.turkey.20171019.131121 test-keys"

	db := o.dataBase.Open(o.databaseName)
	defer o.dataBase.Close(db)

	//o.refreshGatewayInfo(otaDevices)
	// todo 15 dk time delay
	if o.control(db, otaDevices) {
		o.pushApplicationOperation(db, otaDevices)
		o.startApplicationOperation(db, otaDevices)
		o.otaControlOperation(db, otaDevices)
	}

	o.reportDatabase()

}

func (o OtaOp) control(db *sql.DB, ota OtaDeviceArray) bool {
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

func (o OtaOp) refreshGatewayInfo(otaDevices OtaDeviceArray) {
	var (
		devices device.Device
	)
	for _, devicesID := range otaDevices {
		devices.RefreshGatewayInfo(devices.DeviceID2Code(devicesID.Imei))
	}
}

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
func (o OtaOp) pushApplication(db *sql.DB, otaDevice string) {
	var (
		workingsets workingset.Workingset
		//otaUpdaterCode = "071503D1-C864-4236-ABF0-DEA26550AF93" Ota Updater Asıl
		otaUpdaterCode = "DD76AFEA-E0A3-4B61-97CA-509B66A884E1"
	)

	if workingsets.PushApplications(otaUpdaterCode, false, otaDevice) {
		o.dataBase.Update(db, o.tableName, pushApplication, "1", imei, otaDevice)
		o.dataBase.Update(db, o.tableName, pushTime, time.Now().String(), imei, otaDevice)
		fmt.Println("Updateing...")
	}

}

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

	var file *os.File
	writefile.CreateFile("OtaReport.xlsx")
	file = writefile.OpenFile("OtaReport.xlsx", file)
	writefile.WriteText(file, imei, pushApplication, pushTime, startApplication, startTime, controlOta, controlTime)

	var dataBaseReport databasecenter.DB
	db := dataBaseReport.Open(o.databaseName)

	getDatabase := dataBaseReport.Select(db, o.tableName)

	for i := range getDatabase {
		writefile.WriteText(file,
			string(getDatabase[i][imei].([]byte)),

			strconv.FormatInt(getDatabase[i][pushApplication].(int64), 10),
			getDatabase[i][pushTime].(time.Time).String(),

			strconv.FormatInt(getDatabase[i][startApplication].(int64), 10),
			getDatabase[i][startTime].(time.Time).String(),

			strconv.FormatInt(getDatabase[i]["controlOta"].(int64), 10),
			getDatabase[i][controlTime].(time.Time).String())

	}
}
