package reports

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"strconv"

	"github.com/KMACEL/IITR/writefile"
	"github.com/KMACEL/IITR/timop"
	"github.com/KMACEL/IITR/errc"
)

type PingControl struct {
	ControlLogFileName string
	OutputFileName     string
}

func (p PingControl) Start() {
	outputFileName := p.OutputFileName + "_" + timop.GetTimeNamesFormat() + ".xlsx"

	var f *os.File
	writefile.CreateFile(outputFileName)
	f = writefile.OpenFile(f, outputFileName)

	var (
		ip            string
		getCloudDate  string
		id            string
		pingCount     string
		getDeviceDate string
		deviceList    []string
	)

	oldPingCount := make(map[string]int)
	oldPingInformation := make(map[string][]string)

	readPingFile, errRead := ioutil.ReadFile(p.ControlLogFileName)
	errc.ErrorCenter("NOT READ - File Name : "+p.ControlLogFileName, errRead)

	str := string(readPingFile)

	pingFileRow := strings.Split(str, "\n")

	for i := 0; i < len(pingFileRow)-1; i++ {

		pingFileRow[i] = strings.Replace(pingFileRow[i], "%20", " ", -1)
		pingFileColumn := strings.Split(pingFileRow[i], " ")

		ip = pingFileColumn[0]
		getCloudDate = pingFileColumn[3] + pingFileColumn[4]
		getCloudDate = getCloudDate[1:len(getCloudDate)-1]

		idPingDateArray := strings.Split(pingFileColumn[6]+" "+pingFileColumn[7], "id=")
		idArray := strings.Split(idPingDateArray[1], "&time=")
		pingDateArray := strings.Split(idArray[1], "&counter=")

		id = idArray[0]
		pingCount = pingDateArray[1]
		getDeviceDate = pingDateArray[0]

		if !controlArray(id, deviceList) {
			deviceList = append(deviceList, id)
		} else {
			pingCountINT, errConvert := strconv.Atoi(pingCount)
			errc.ErrorCenter("NOT INTEGER - ID : "+id, errConvert)

			if oldPingCount[id] != 0 && pingCountINT != oldPingCount[id] && pingCountINT > oldPingCount[id] {
				fmt.Println(id, pingCount, getCloudDate, getDeviceDate, ip)
				writefile.WriteText(f, id, pingCount, getCloudDate, getDeviceDate, ip, " ", oldPingInformation[id][1], oldPingInformation[id][2], oldPingInformation[id][3])
			}
			oldPingCount[id], _ = strconv.Atoi(pingCount)
			oldPingCount[id] = oldPingCount[id] + 2

			oldPingInformation[id] = []string{id, pingCount, getCloudDate, getDeviceDate, ip}
		}
	}
}

func controlArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
