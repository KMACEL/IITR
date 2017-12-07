package reports

import (
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest"
	"encoding/json"
	"fmt"
	"os"
	"log"
	"github.com/KMACEL/IITR/writefile"
)

// For use Example:
//     var instantControl util.InstantApplication
//     instantControl.PackageName = "com.estoty.game2048"
//     instantControl.Start("867377020740787")

type InstantApplication struct {
	PackageName string
	file        os.File
}

func (i InstantApplication) Start(devicesID ...string) {
	var (
		devices    device.Device
		instantApp device.InstantApplicationInfoJSON
		file       *os.File
	)

	writefile.CreateFile("instantControl.xlsx")
	file = writefile.OpenFile(file,"instantControl.xlsx")

	for _, deviceID := range devicesID {
		query := devices.InstantApplicationInfo(deviceID, rest.NOMarshal, rest.Visible)

		if query != nil {
			if string(query) != rest.ResponseNotFound {
				json.Unmarshal(query, &instantApp)
				if i.PackageName == instantApp.Data.PackageName {
					fmt.Println(deviceID, " is install : ", i.PackageName)
					writefile.WriteText(file,deviceID,"OK")
				}else{
					writefile.WriteText(file,deviceID,"NO")
				}
			}
			//return []byte(rest.ResponseNotFound)
		}
		//return []byte(rest.ResponseNil)
	}
}

func (i InstantApplication) instantLog(logString ...interface{}) {
	if _, err := os.Stat("./logc"); os.IsNotExist(err) {
		os.MkdirAll("./logc", os.ModePerm)
	}

	f, err := os.OpenFile("logc/instantLog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(logString)
}
