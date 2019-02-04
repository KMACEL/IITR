package operations

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/KMACEL/IITR/rest/device"
)

/*
 █████╗ ██████╗ ██████╗         ██╗      █████╗ ██████╗ ███████╗██╗
██╔══██╗██╔══██╗██╔══██╗        ██║     ██╔══██╗██╔══██╗██╔════╝██║
███████║██║  ██║██║  ██║        ██║     ███████║██████╔╝█████╗  ██║
██╔══██║██║  ██║██║  ██║        ██║     ██╔══██║██╔══██╗██╔══╝  ██║
██║  ██║██████╔╝██████╔╝        ███████╗██║  ██║██████╔╝███████╗███████╗
╚═╝  ╚═╝╚═════╝ ╚═════╝         ╚══════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝
*/

// AddLabel is
type AddLabel struct{}

// Start is
func (a AddLabel) Start(deviceID ...string) {

	data, err := ioutil.ReadFile(initPath)
	if err != nil {
		fmt.Println("Can't read file:", initPath)
		panic(err)
	}

	dataList := strings.Split(string(data), "\n")

	for _, data := range dataList {
		if data != "" {
			parseData := strings.Split(data, ",")
			fmt.Println(parseData)
			fmt.Println(device.Device{}.SetLabel(parseData[0], parseData[1], true))
		}
	}

}
