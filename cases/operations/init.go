package operations

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/KMACEL/IITR/rest"
	"golang.org/x/crypto/ssh/terminal"
)

/*
██╗███╗   ██╗██╗████████╗
██║████╗  ██║██║╚══██╔══╝
██║██╔██╗ ██║██║   ██║
██║██║╚██╗██║██║   ██║
██║██║ ╚████║██║   ██║
╚═╝╚═╝  ╚═══╝╚═╝   ╚═╝
*/

// Use the in "main" as
//		_ = operations.Operations{}

// Starter is operations interface
type Starter interface {
	Start(deviceID ...string)
}

// Operations is
type Operations struct {
}

var (
	password       = ""
	userName       = ""
	initPath       = ""
	operationsCase = ""
)

const (
	caseLabel        = "label"
	caseDeleteDrom   = "deletedrom"
	caseClearLicense = "clearlicense"
	caseAddDrom      = "caseadddrom"
)

func init() {

	_userName := flag.String("username", "", "https://enterprise.iot-ignite.com Login Username. \n")
	_path := flag.String("path", "", "For Label Device ID, Label Format. \n For Delete Drom Device ID")
	_case := flag.String("case", "", "Operations Case Type : \n\t"+
		caseLabel+"        : this case updates the label of the device. DeviceID to be selected for use DeviceID, Label <newline>\n\t"+
		caseClearLicense+" : this case deletes the license defined on the device. File content to be selected for use: DeviceID <newline>\n\t"+
		caseDeleteDrom+"   : this case deletes the drom defined on the device. File content to be selected for use: DeviceID <newline>\n\t"+
		caseAddDrom+"  	   : this case add the drom defined on the device. File content to be selected for use: ConfigurationnName <newline> DeviceID <newline>\n\t")

	flag.Parse()

	userName = *_userName
	initPath = *_path
	operationsCase = *_case

	if len(operationsCase) != 0 {
		setPassword()
	}

	if operationsCase == caseLabel {
		start(AddLabel{})
	} else if operationsCase == caseClearLicense {
		start(ClearLicense{})
	} else if operationsCase == caseDeleteDrom {
		start(DeleteDrom{})
	} else if operationsCase == caseAddDrom {
		start(AddDrom{})
	}
}

func start(starter Starter) {
	starter.Start()
}

func setPassword() {
	if len(userName) > 0 {
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err == nil {
			password := string(bytePassword)
			connection, err := rest.Connect(userName, password)
			if connection == true {
				if len(os.Args) < 2 {
					fmt.Println("Missing parameter, provide file name!")
					return
				}
			} else {
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	} else {
		fmt.Println("Use -username={YOUR_USERNAME} flag")
		panic("-username not found")
	}
}
