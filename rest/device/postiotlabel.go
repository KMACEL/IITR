package device

import (
	"encoding/json"
	"fmt"

	"github.com/KMACEL/IITR/errc"
)

func (d Device) AddIOTLabel(iotLabelREQ AddIOTLabelJSON, visualFlag bool) {
	setQueryAddress := addIOTLabelLink()
	setBody, _ := json.Marshal(iotLabelREQ)
	fmt.Println(string(setBody))
	_, errCreateIOTLabel := q.PostQuery(setQueryAddress, string(setBody), contentTypeJSON(), visualFlag)
	errc.ErrorCenter("AddIOTLabel", errCreateIOTLabel)
}
