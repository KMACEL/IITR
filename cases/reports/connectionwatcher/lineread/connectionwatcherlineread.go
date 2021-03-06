package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/KMACEL/IITR/writefile"
)

var writeArray []string

func main() {
	var f *os.File

	if len(os.Args) < 3 {
		fmt.Println("Missing parameter, provide file name!")
		panic("Please entry file name...")
	}

	log.Println("Reading...")

	//-----------------------------------------
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("err : ", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	log.Println("Read.")
	log.Println("Processing...")

	writeFileName := "ConnectionWatcher-" + os.Args[1] + "-" + os.Args[2] + ".csv"
	writefile.CreateFile(writeFileName)
	f = writefile.OpenFile(f, writeFileName)
	writefile.WriteText(f, " ", "IP", "Server Time", "Device ID", "Send Time", "Counter", "Connection Type", "DataState", "Signal DB", "New Time", "Long", "Lat", "GPS Time", "Satellite", "Power State")

	writeArray = append(writeArray, "")

	var i int
	for scanner.Scan() {
		i++
		u, err := url.Parse(scanner.Text())
		if err != nil {
			panic(err)
		}
		m, _ := url.ParseQuery(u.RawQuery)
		if m["id"][0] == os.Args[2] {

			pingFileColumn := strings.Split(scanner.Text(), " ")
			ip := pingFileColumn[0]
			serverTime := pingFileColumn[3] + pingFileColumn[4]
			writeArray = append(writeArray, ip, serverTime)

			ControlAdd(m["id"])
			ControlAdd(m["time"])
			ControlAdd(m["counter"])

			//ControlAdd(m["connType"])
			if m["connType"] != nil {
				connTypeSplit := strings.Split(m["connType"][0], " ")
				if len(connTypeSplit) != 0 {
					var connType string

					switch connTypeSplit[0] {
					case "11":
						connType = "2G"
					case "12", "14", "15", "17":
						connType = "3G"
					case "13", "18", "19":
						connType = "4G"
					default:
						connType = "?"
					}
					writeArray = append(writeArray, connType)
				}
			} else {
				writeArray = append(writeArray, " ")
			}

			//ControlAdd(m["dataState"])
			if m["dataState"] != nil {
				dataStateSplit := strings.Split(m["dataState"][0], " ")
				if len(dataStateSplit) != 0 {
					var dataState string

					switch dataStateSplit[0] {
					case "0":
						dataState = "DISCONNECTED"
					case "1":
						dataState = "CONNECTING"
					case "2":
						dataState = "CONNECTED"
					case "3":
						dataState = "SUSPENDED"
					default:
						dataState = "?"
					}
					writeArray = append(writeArray, dataState)
				}
			} else {
				writeArray = append(writeArray, " ")
			}

			ControlAdd(m["signalDb"])
			ControlAdd(m["NewTime"])

			if m["J"] != nil {
				counterSplit := strings.Split(m["J"][0], " ")
				if len(counterSplit) != 0 {
					long, _ := strconv.ParseFloat(counterSplit[0], 64)
					long = long / 600000
					writeArray = append(writeArray, strconv.FormatFloat(long, 'E', -1, 64))
				}
			} else {
				writeArray = append(writeArray, " ")
			}

			if m["W"] != nil {
				counterSplit := strings.Split(m["W"][0], " ")
				if len(counterSplit) != 0 {
					long, _ := strconv.ParseFloat(counterSplit[0], 64)
					long = long / 600000
					writeArray = append(writeArray, strconv.FormatFloat(long, 'E', -1, 64))
				}
			} else {
				writeArray = append(writeArray, " ")
			}

			//ControlAdd(m["J"])
			//ControlAdd(m["W"])
			ControlAdd(m["T"])
			ControlAdd(m["PN"])
			ControlAdd(m["ACC"])

			writeArray = append(writeArray, "\n")

		}

		if i%100000 == 0 {
			fmt.Println("Size :  ", i)
		}
	}

	log.Println("The data was processed...")
	log.Println("Preparing Report.")
	writefile.WriteArray(f, writeArray)

}

//ControlAdd is
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
