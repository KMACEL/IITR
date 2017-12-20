package operations

import (
	"fmt"
	"github.com/KMACEL/IITR/rest/profile"
)

//ModeSendOperation is
type ModeSendOperation struct{}

//SendProfileWant is
type SendProfileWant struct {
	ModeCode   string
	PolicyCode string
	DevicesID  []string
}

//Start is
func (m ModeSendOperation) Start(sendProfile SendProfileWant) {
	var (
		profiles profile.Profile
	)
	fmt.Println(profiles.PushModeAuto(sendProfile.ModeCode, sendProfile.PolicyCode, sendProfile.DevicesID...))
}
