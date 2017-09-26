package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/KMACEL/IITR/cases"
	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/timop"
)

func report() {
	var reportOperation string

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

	cases.DetailReportDevices{}.Start("ApplicationPackage_"+timop.GetTimeNamesFormat()+".xlsx", reportDevices.Case.Devices, reportDevices.Case.Packages)
}

func allReport() {
	var applicationPackage string
	fmt.Println("Information : ")
	fmt.Println("Please type in the application package names you want to look like ',' and press the 'enter' key.")

	fmt.Print("Application Package Names : ")
	fmt.Scan(&applicationPackage)

	splitString := strings.Split(applicationPackage, ",")
	fmt.Println("Split : ", splitString)
	cases.DetailReport{}.Start("ApplicationPackage_"+timop.GetTimeNamesFormat()+".xlsx", splitString)

}
