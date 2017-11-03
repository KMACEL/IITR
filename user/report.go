package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/KMACEL/IITR/cases/reports"
	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/timop"
)

func report() {
	var reportOperation string

	fmt.Println("\n\n********************************************************************************")
	fmt.Println("********************************************************************************\n ")
	fmt.Println("Report Screen :")
	fmt.Println("adr / alldevicereport : Take the package names of the applications to be looked at from you and control all devices that are tenant\n ")
	fmt.Println("dr / devicereport : I want a way from you. he requests a json format file on the road. it generates a report according to the device and information in that file\n ")
	fmt.Println("br / blockedreport : make devices blocked")
	fmt.Println("\n********************************************************************************")
	fmt.Println("********************************************************************************\n ")

	fmt.Print("Entry Report Operation Parameter : ")
	fmt.Scan(&reportOperation)

	if reportOperation == "alldevicesreport" || reportOperation == "adr" {
		allReport()
	}

	if reportOperation == "devicesreport" || reportOperation == "dr" {
		log.Println("Operation is Read File Selected.")
		devicesReport()
	}
}

func devicesReport() {
	var path string
	log.Println("Welcome to Read File Operation. Please to Change Read File Path...")
	fmt.Print("File Path : ")
	fmt.Scan(&path)

	readingFile, errPath := ioutil.ReadFile(path)
	errc.ErrorCenter("Path Error : ", errPath)

	reportDevices := devicesReportJSON{}
	json.Unmarshal(readingFile, &reportDevices)

	log.Println("Type : ", reportDevices.Type)
	fmt.Println("Cases Name : ", reportDevices.Case.Name)
	fmt.Println("Devices : ", reportDevices.Case.Devices)
	fmt.Println("Applications : ", reportDevices.Case.Packages)

	reports.DetailReportDevices{}.Start("ApplicationPackage_"+timop.GetTimeNamesFormat()+".xlsx", reportDevices.Case.Devices, reportDevices.Case.Packages)
}

func allReport() {
	var applicationPackage string
	fmt.Println("Information : ")
	fmt.Println("Please type in the application package names you want to look like ',' and press the 'enter' key.")

	fmt.Print("Application Package Names : ")
	fmt.Scan(&applicationPackage)

	splitString := strings.Split(applicationPackage, ",")
	fmt.Println("Split : ", splitString)
	reports.DetailAllReport{}.Start("ApplicationPackage_"+timop.GetTimeNamesFormat()+".xlsx", splitString)

}
