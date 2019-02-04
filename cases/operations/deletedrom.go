package operations

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/KMACEL/IITR/rest/drom"
)

// DeleteDrom is
type DeleteDrom struct {
}

// Start is
func (d DeleteDrom) Start(deviceID ...string) {
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
				fmt.Println(drom.Drom{}.DeleteDrom(data, true))
			}
		}
	}

}
