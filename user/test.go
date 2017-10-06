package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/drom"
	"github.com/KMACEL/IITR/timop"
	"github.com/KMACEL/IITR/writefile"
)

func test() {

	var testOperation string
	fmt.Println("\n\n********************************************************************************")
	fmt.Println("********************************************************************************\n ")
	log.Println("Test ...")
	fmt.Println("new | add : Create New Test Case")
	fmt.Println("cl  | cases-list : Get Cases Test Case List")
	fmt.Println("sc  | start-cases : Start Test Cases")
	fmt.Println("\n\n********************************************************************************")
	fmt.Println("********************************************************************************\n ")

	fmt.Print("Operation : ")
	fmt.Scan(&testOperation)

	if testOperation == "new" || testOperation == "add" {
		newTest()
	}

	if testOperation == "cases-list" || testOperation == "cl" {
		getCasesList()
	}

	if testOperation == "start-cases" || testOperation == "sc" {
		fmt.Println("===================================================")
		getCasesList()
		fmt.Println("===================================================")

		var caseName string
		fmt.Print("Plase Entry Test Case : ")
		fmt.Scan(&caseName)
		startTestCase(caseName)
	}
}

func newTest() {
	var path string
	log.Println("Welcome to Read File Test Operation. Please to Change Read File Path...")
	fmt.Print("File Path : ")
	fmt.Scan(&path)

	readingFile, errPath := ioutil.ReadFile(path)
	errc.ErrorCenter("Path Error : ", errPath)

	newTestCase := newTestCaseJSON{}
	json.Unmarshal(readingFile, &newTestCase)

	if newTestCase.Type == "test" && newTestCase.Code == "new-test" && newTestCase.Name != "" {
		var saveTestCase savedTestCasesJSON
		saveTestCase.Name = newTestCase.Name

		saveTestCase.Devices = newTestCase.Devices
		saveTestCase.Case = newTestCase.Case
		jsonMarshal, _ := json.Marshal(saveTestCase)

		saveCase(jsonMarshal, saveTestCase.Name)

	}
}

func saveCase(caseValue []byte, fileName string) {
	var saveCaseFile *os.File
	writefile.CreateFile("cases_lib/" + fileName)
	saveCaseFile = writefile.OpenFile("cases_lib/"+fileName, saveCaseFile)
	writefile.WriteByte(caseValue, saveCaseFile)

}

func getCasesList() {
	files, err := ioutil.ReadDir("cases_lib/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func startTestCase(testCases string) {
	readingFile, errPath := ioutil.ReadFile("cases_lib/" + testCases)
	errc.ErrorCenter("Path Error : ", errPath)

	changeTestCase := savedTestCasesJSON{}
	json.Unmarshal(readingFile, &changeTestCase)

	log.Println("Start Test : ", changeTestCase.Name)

	for _, testObject := range changeTestCase.Case {
		var loopValue int
		if testObject.Loop == -1 {
			loopValue = 150
		} else {
			loopValue = testObject.Loop
		}
		for loop := 0; loop < loopValue; loop++ {
			for _, cases := range testObject.Steps {
				fmt.Println(cases)
				casesSplit := strings.Split(cases, "::")

				/*
				   ██████╗ ██████╗  ██████╗ ███╗   ███╗
				   ██╔══██╗██╔══██╗██╔═══██╗████╗ ████║
				   ██║  ██║██████╔╝██║   ██║██╔████╔██║
				   ██║  ██║██╔══██╗██║   ██║██║╚██╔╝██║
				   ██████╔╝██║  ██║╚██████╔╝██║ ╚═╝ ██║
				   ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝
				*/
				if casesSplit[0] == "drom" {
					for _, dromDevices := range changeTestCase.Devices {
						drom.Drom{}.SendDrom(rest.Invisible, dromDevices)
						log.Println("Send Drom : ", dromDevices)
					}
				}

				/*
				   ██████╗ ███████╗██████╗  ██████╗  ██████╗ ████████╗
				   ██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝
				   ██████╔╝█████╗  ██████╔╝██║   ██║██║   ██║   ██║
				   ██╔══██╗██╔══╝  ██╔══██╗██║   ██║██║   ██║   ██║
				   ██║  ██║███████╗██████╔╝╚██████╔╝╚██████╔╝   ██║
				   ╚═╝  ╚═╝╚══════╝╚═════╝  ╚═════╝  ╚═════╝    ╚═╝
				*/
				if casesSplit[0] == "reboot" {
					for _, dromDevices := range changeTestCase.Devices {
						device.Device{}.Reboot(device.Device{}.DeviceID2Code(dromDevices), rest.Invisible)
						log.Println("Send Reboot : ", dromDevices)
					}
				}

				/*
				    ██████╗██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███████╗    ███╗   ███╗ ██████╗ ██████╗ ███████╗
				   ██╔════╝██║  ██║██╔══██╗████╗  ██║██╔════╝ ██╔════╝    ████╗ ████║██╔═══██╗██╔══██╗██╔════╝
				   ██║     ███████║███████║██╔██╗ ██║██║  ███╗█████╗      ██╔████╔██║██║   ██║██║  ██║█████╗
				   ██║     ██╔══██║██╔══██║██║╚██╗██║██║   ██║██╔══╝      ██║╚██╔╝██║██║   ██║██║  ██║██╔══╝
				   ╚██████╗██║  ██║██║  ██║██║ ╚████║╚██████╔╝███████╗    ██║ ╚═╝ ██║╚██████╔╝██████╔╝███████╗
				    ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝    ╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝
				*/
				if casesSplit[0] == "changemode" {
				}

				/*
				    ██████╗ ███████╗████████╗    ██╗      ██████╗  ██████╗
				   ██╔════╝ ██╔════╝╚══██╔══╝    ██║     ██╔═══██╗██╔════╝
				   ██║  ███╗█████╗     ██║       ██║     ██║   ██║██║  ███╗
				   ██║   ██║██╔══╝     ██║       ██║     ██║   ██║██║   ██║
				   ╚██████╔╝███████╗   ██║       ███████╗╚██████╔╝╚██████╔╝
				    ╚═════╝ ╚══════╝   ╚═╝       ╚══════╝ ╚═════╝  ╚═════╝
				*/
				if casesSplit[0] == "getlog" {
					for _, dromDevices := range changeTestCase.Devices {
						device.Device{}.GetDeviceLog(device.Device{}.DeviceID2Code(dromDevices), rest.Invisible)
						log.Println("Send Get Log : ", dromDevices)
					}
				}

				/*
				   ███████╗ ██████╗██████╗ ███████╗███████╗███╗   ██╗███████╗██╗  ██╗ ██████╗ ████████╗
				   ██╔════╝██╔════╝██╔══██╗██╔════╝██╔════╝████╗  ██║██╔════╝██║  ██║██╔═══██╗╚══██╔══╝
				   ███████╗██║     ██████╔╝█████╗  █████╗  ██╔██╗ ██║███████╗███████║██║   ██║   ██║
				   ╚════██║██║     ██╔══██╗██╔══╝  ██╔══╝  ██║╚██╗██║╚════██║██╔══██║██║   ██║   ██║
				   ███████║╚██████╗██║  ██║███████╗███████╗██║ ╚████║███████║██║  ██║╚██████╔╝   ██║
				   ╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝    ╚═╝
				*/
				if casesSplit[0] == "screenshot" {
				}

				/*
				   ███████╗████████╗ █████╗ ██████╗ ████████╗     █████╗ ██████╗ ██████╗
				   ██╔════╝╚══██╔══╝██╔══██╗██╔══██╗╚══██╔══╝    ██╔══██╗██╔══██╗██╔══██╗
				   ███████╗   ██║   ███████║██████╔╝   ██║       ███████║██████╔╝██████╔╝
				   ╚════██║   ██║   ██╔══██║██╔══██╗   ██║       ██╔══██║██╔═══╝ ██╔═══╝
				   ███████║   ██║   ██║  ██║██║  ██║   ██║       ██║  ██║██║     ██║
				   ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝       ╚═╝  ╚═╝╚═╝     ╚═╝
				*/
				if casesSplit[0] == "startapp" {
				}

				/*
				   ███████╗████████╗ ██████╗ ██████╗      █████╗ ██████╗ ██████╗
				   ██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗    ██╔══██╗██╔══██╗██╔══██╗
				   ███████╗   ██║   ██║   ██║██████╔╝    ███████║██████╔╝██████╔╝
				   ╚════██║   ██║   ██║   ██║██╔═══╝     ██╔══██║██╔═══╝ ██╔═══╝
				   ███████║   ██║   ╚██████╔╝██║         ██║  ██║██║     ██║
				   ╚══════╝   ╚═╝    ╚═════╝ ╚═╝         ╚═╝  ╚═╝╚═╝     ╚═╝
				*/
				if casesSplit[0] == "stopapp" {
				}

				/*
				    ██████╗██╗     ███████╗ █████╗ ██████╗     ██████╗  █████╗ ████████╗ █████╗
				   ██╔════╝██║     ██╔════╝██╔══██╗██╔══██╗    ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗
				   ██║     ██║     █████╗  ███████║██████╔╝    ██║  ██║███████║   ██║   ███████║
				   ██║     ██║     ██╔══╝  ██╔══██║██╔══██╗    ██║  ██║██╔══██║   ██║   ██╔══██║
				   ╚██████╗███████╗███████╗██║  ██║██║  ██║    ██████╔╝██║  ██║   ██║   ██║  ██║
				    ╚═════╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝    ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝
				*/
				if casesSplit[0] == "cleardata" {
				}

				/*
				   ██╗    ██╗██╗██████╗ ███████╗
				   ██║    ██║██║██╔══██╗██╔════╝
				   ██║ █╗ ██║██║██████╔╝█████╗
				   ██║███╗██║██║██╔═══╝ ██╔══╝
				   ╚███╔███╔╝██║██║     ███████╗
				    ╚══╝╚══╝ ╚═╝╚═╝     ╚══════╝
				*/
				if casesSplit[0] == "wipe" {

				}

				if strings.Contains(casesSplit[1], "d&") {
					index := strings.Index(casesSplit[1], "d&")
					delayTime := (casesSplit[1])[index+2:]

					if !strings.Contains(delayTime, "?") {
						delayInt, err := strconv.ParseInt(delayTime, 10, 64)
						errc.ErrorCenter("Convert Delay Time Error : ", err)

						fmt.Println("Delay Time : ", delayInt)
						timop.Delay{}.M(int(delayInt))
					} else {
						randomRange := delayTime[strings.Index(delayTime, "?")+1:]
						randomArray := strings.Split(randomRange, "-")
						begin, _ := strconv.ParseInt(randomArray[0], 10, 64)
						end, _ := strconv.ParseInt(randomArray[1], 10, 64)
						randomValue := timop.Random(int(begin), int(end))

						fmt.Println("Delay Time : ", randomValue)
						timop.Delay{}.M(randomValue)

					}
				}

			}
		}
	}
}
