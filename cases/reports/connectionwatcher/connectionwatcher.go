package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	"github.com/KMACEL/IITR/timop"
	"github.com/KMACEL/IITR/writefile"
)

var writeArray []string

func main() {
	var f *os.File

	var deviceList []string

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		panic("Please entry file name...")
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	devList := strings.Split(string(data), "\n")
	for _, deviceID := range devList {
		if deviceID != "" {
			deviceList = append(deviceList, deviceID)
		}
	}

	writeFileName := "ConnectionWatcher-" + timop.GetTimeNamesFormat() + ".csv"
	writefile.CreateFile(writeFileName)
	f = writefile.OpenFile(f, writeFileName)
	writefile.WriteText(f, " ", "IP", "Server Time", "Device ID", "Send Time", "Counter", "Connection Type", "Signal DB", "New Time", "J", "W", "GPS Time", "PN", "Power State")

	writeArray = append(writeArray, "")

	for i, rowData := range deviceList {

		u, err := url.Parse(rowData)
		if err != nil {
			panic(err)
		}
		m, _ := url.ParseQuery(u.RawQuery)

		pingFileColumn := strings.Split(rowData, " ")
		ip := pingFileColumn[0]
		serverTime := pingFileColumn[3] + pingFileColumn[4]
		writeArray = append(writeArray, ip, serverTime)

		ControlAdd(m["id"])
		ControlAdd(m["time"])
		ControlAdd(m["counter"])
		ControlAdd(m["connType"])
		ControlAdd(m["signalDb"])
		ControlAdd(m["NewTime"])
		ControlAdd(m["J"])
		ControlAdd(m["W"])
		ControlAdd(m["T"])
		ControlAdd(m["PN"])
		ControlAdd(m["ACC"])

		writeArray = append(writeArray, "\n")

		if i%100 == 0 {
			fmt.Println(i, " - ", len(deviceList))
		}
	}
	writefile.WriteArray(f, writeArray)
}

func ControlAdd(setValue []string) {
	if setValue != nil {

		counterSplit := strings.Split(setValue[0], " ")
		if len(counterSplit) != 0 {
			writeArray = append(writeArray, counterSplit[0])
		}

	} else {
		writeArray = append(writeArray, " ")
	}
}
