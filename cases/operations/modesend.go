package operations

import (
	"github.com/KMACEL/IITR/rest/profile"
	"fmt"
)

type ModeSendOperation struct{}

type SendProfileWant struct {
	ModeCode   string
	PolicyCode string
	DevicesID  []string
}

func (m ModeSendOperation) Start(sendProfile SendProfileWant) {
	var (
		profiles profile.Profile
	)
	fmt.Println(profiles.PushModeAuto(sendProfile.ModeCode, sendProfile.PolicyCode, sendProfile.DevicesID...))
}
