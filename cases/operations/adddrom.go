package operations

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/KMACEL/IITR/rest/drom"
)

// AddDrom is
type AddDrom struct {
	ConfigurationName string
}

// Start is
func (a AddDrom) Start(deviceID ...string) {
	var configurationID string
	if len(initPath) > 0 {
		data, err := ioutil.ReadFile(initPath)
		if err != nil {
			fmt.Println("Can't read file:", initPath)
			panic(err)
		}

		split := strings.Split(string(data), "\n")
		configurationID = drom.Drom{}.GetDromConfiguration(split[0], false, false)
		fmt.Println("ConfigurationName : ", split[0])
		deviceID = split[1:]
	} else {
		configurationID = drom.Drom{}.GetDromConfiguration(a.ConfigurationName, false, false)
		fmt.Println("ConfigurationName : ", a.ConfigurationName)
	}

	fmt.Println("configurationID : ", configurationID)
	fmt.Println("deviceID : ", deviceID)

	for _, data := range deviceID {
		if data != "" {
			fmt.Println(data)
			fmt.Println(drom.Drom{}.AddDevice(data, configurationID))
		}
	}
}
