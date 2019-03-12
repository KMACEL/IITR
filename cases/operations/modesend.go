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
	ModeName   string
	PolicyName string
}

//Start is
func (m ModeSendOperation) Start(sendProfile SendProfileWant) {
	var (
		profiles profile.Profile
	)

	profileCode, policyCode := profiles.GetSelectProfileInPolicyCode(sendProfile.ModeName, sendProfile.PolicyName)

	if len(sendProfile.ModeCode) != 0 {
		profileCode = profiles.GetProfile(sendProfile.ModeName).Code
	}

	if len(sendProfile.PolicyCode) != 0 {
		policyCode = sendProfile.PolicyCode
	}

	fmt.Println("Profile CODE : ", profileCode)
	fmt.Println("Policy CODE : ", policyCode)

	fmt.Println(profiles.PushModeAuto(profileCode, policyCode, sendProfile.DevicesID...))
}
