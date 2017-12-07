package reports

import (
	"encoding/json"
	"os"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/writefile"
)

type Profile struct{}

func (p Profile) Start() {

}

func (p Profile) profilePolicy(devicesID ...string) {

	var devices device.Device
	var modePolicyOS *os.File
	writefile.CreateFile("ModePolicy.xlsx")
	writefile.OpenFile(modePolicyOS, "ModePolicy.xlsx")
	writefile.WriteText(modePolicyOS, "Device ID", "Active Mode", "Active Policy", "Current Mode", "Current Policy")

	query := devices.LocationMap(rest.NOMarshal, rest.Invisible)

	// This assignment is aimed at resetting the variable
	deviceCode := device.LocationJSON{}
	json.Unmarshal(query, &deviceCode)

	for _, devicessID := range deviceCode.Extras {
		deviceID := devicessID.DeviceID
		if deviceID != "" {
			var (
				activeProfile string
				activePolicy  string

				currentProfile string
				currentPolicy  string
			)

			profilePolicyQuery := devices.ActiveProfilePolicy(deviceID, rest.NOMarshal, rest.Invisible)
			if profilePolicyQuery != nil {
				if string(profilePolicyQuery) != rest.ResponseNotFound {
					activeProfilePolicy := device.ActiveProfilePolicyJSON{}
					json.Unmarshal(profilePolicyQuery, &activeProfilePolicy)

					activeProfile = activeProfilePolicy.ActiveProfile
					activePolicy = activeProfilePolicy.ActivePolicy

					currentProfile = activeProfilePolicy.CurrentProfile
					currentPolicy = activeProfilePolicy.CurrentProfile

					if len(activeProfile) == 0 {
						activeProfile = rest.ResponseNil
					}
					if len(activePolicy) == 0 {
						activePolicy = rest.ResponseNil
					}

					if len(currentProfile) == 0 {
						currentProfile = rest.ResponseNil
					}
					if len(currentPolicy) == 0 {
						currentPolicy = rest.ResponseNil
					}
				} else {
					activeProfile = rest.ResponseNotFound
					activePolicy = rest.ResponseNotFound
					currentProfile = rest.ResponseNotFound
					currentPolicy = rest.ResponseNotFound
				}
			} else {
				activeProfile = rest.ResponseNil
				activePolicy = rest.ResponseNil
				currentProfile = rest.ResponseNil
				currentPolicy = rest.ResponseNil
			}

			if activeProfile == "" {
				activeProfile = rest.ResponseNil
			}

			if activePolicy == "" {
				activePolicy = rest.ResponseNil
			}

			if currentProfile == "" {
				currentProfile = rest.ResponseNil
			}

			if currentPolicy == "" {
				currentPolicy = rest.ResponseNil
			}
			writefile.WriteText(modePolicyOS, deviceID, activeProfile, activePolicy, currentProfile, currentPolicy)
		}
	}
}
