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

func main() {
	var f *os.File
	var writeArray []string
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

	writeFileName := "ConnectionWatcher-" + timop.GetTimeNamesFormat() + ".xlsx"
	writefile.CreateFile(writeFileName)
	f = writefile.OpenFile(f, writeFileName)
	writefile.WriteText(f, " ", "IP", "Server Time", "Device ID", "Send Time", "Counter", "Connection Type", "Signal DB", "New Time", "J", "W", "GPS Time", "PN")

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

		if m["id"] != nil {
			writeArray = append(writeArray, m["id"][0], m["time"][0], m["counter"][0], m["connType"][0], m["signalDb"][0])
		} else {
			writeArray = append(writeArray, " ", " ", " ", " ", " ")
		}

		if m["NewTime"] != nil {
			writeArray = append(writeArray, m["NewTime"][0])
		} else {
			writeArray = append(writeArray, " ")
		}

		if m["J"] != nil {
			writeArray = append(writeArray, m["J"][0], m["W"][0], m["T"][0], m["PN"][0])
		} else {
			writeArray = append(writeArray, " ", " ", " ", " ")
		}

		writeArray = append(writeArray, "\n")

		if i%100 == 0 {
			fmt.Println(i, " - ", len(deviceList))
		}
	}
	writefile.WriteArray(f, writeArray)
}
