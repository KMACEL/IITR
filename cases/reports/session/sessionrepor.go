package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/writefile"
)

func main() {
	var f *os.File
	fileName := "TestSessions.xlsx"

	writefile.CreateFile(fileName)
	f = writefile.OpenFile(f, fileName)
	writefile.WriteText(f, "Device ID", "Begin", "End", "Begin Epoch", "End Epoch", "Diff Minute", "Small 5")

	readPingFile, errRead := ioutil.ReadFile("aa.csv")
	errc.ErrorCenter("NOT READ - File Name : "+"aa.csv", errRead)

	sessionList := string(readPingFile)

	sessionListSplit := strings.Split(sessionList, "\n")

	for _, deviceSession := range sessionListSplit {
		sessionDeviceSplit := strings.Split(deviceSession, ",")
		if sessionDeviceSplit[0] == "867377020889089" {
			var (
				beginTime  int
				endTime    int
				difference string
				diffMinute string
			)

			fmt.Println(sessionDeviceSplit)
			beginTime, _ = (strconv.Atoi(sessionDeviceSplit[1]))
			endTime, _ = (strconv.Atoi(sessionDeviceSplit[2]))
			diffMinute = strconv.FormatFloat((float64(endTime-beginTime) / 1000 / 60), 'f', 6, 64)
			if (endTime-beginTime)/1000 < 1 {
				difference = "*"
			}
			writefile.WriteText(f, sessionDeviceSplit[0], time.Unix(0, int64(beginTime)*1000000).String(), time.Unix(0, int64(endTime)*1000000).String(), sessionDeviceSplit[1], sessionDeviceSplit[2], diffMinute, difference)
		}
	}
}
