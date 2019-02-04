package operations

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/KMACEL/IITR/rest/license"
)

// ClearLicense is
type ClearLicense struct {
}

// Start is
func (d ClearLicense) Start(deviceID ...string) {

	if len(initPath) > 0 {
		data, err := ioutil.ReadFile(initPath)
		if err != nil {
			fmt.Println("Can't read file:", initPath)
			panic(err)
		}
		dataList := strings.Split(string(data), "\n")

		for _, data := range dataList {
			if data != "" {
				fmt.Println(data)
				fmt.Println(license.License{}.ClearLicense(data, true))
			}
		}
	}

}
