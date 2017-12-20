package tests

import (
	"fmt"
	"log"

	"github.com/KMACEL/IITR/rest/profile"
	"github.com/KMACEL/IITR/timop"
)

// ModeSwitch is
type ModeSwitch struct{}

// Start is
func (m ModeSwitch) Start() {
	serviceMode := "5AD8896B-6B74-48C8-B0BF-FDD693590539"
	servicePolicy := "703EE61C-6123-472F-A567-B6FE9EBCE5C2"

	productionMode := "C4633E1E-3F4D-4BB6-8F36-82E25AEDA054"
	productionPolicy := "962538F4-D668-47B0-AB08-8D918FEFDAFF"

	devices := "867377020915728"

	for {
		m.changeMode(serviceMode, servicePolicy, devices)
		log.Println("Change Service Mode")
		timop.Delay{}.M(timop.Random(0, 6))
		m.changeMode(productionMode, productionPolicy, devices)
		log.Println("Change Production Mode")
		timop.Delay{}.M(timop.Random(0, 6))
	}
}

func (m ModeSwitch) changeMode(modeCode string, policyCode string, deviceID ...string) {
	var (
		profiles profile.Profile
	)
	fmt.Println(profiles.PushModeAuto(modeCode, policyCode, deviceID...))
}
