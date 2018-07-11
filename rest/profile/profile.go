package profile

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/workingset"
)

/*
██████╗ ██╗   ██╗███████╗██╗  ██╗        ███╗   ███╗ ██████╗ ██████╗ ███████╗
██╔══██╗██║   ██║██╔════╝██║  ██║        ████╗ ████║██╔═══██╗██╔══██╗██╔════╝
██████╔╝██║   ██║███████╗███████║        ██╔████╔██║██║   ██║██║  ██║█████╗
██╔═══╝ ██║   ██║╚════██║██╔══██║        ██║╚██╔╝██║██║   ██║██║  ██║██╔══╝
██║     ╚██████╔╝███████║██║  ██║        ██║ ╚═╝ ██║╚██████╔╝██████╔╝███████╗
╚═╝      ╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝
*/

/*
	profile.Profile{}.PushMode("{YOUR_WORKING_GROUP_KEY}", profile.Profile{}.GetProfile("{YOUR_PROFILE_NAME}").Code, profile.Profile{}.GetProfile("{YOUR_PROFILE_NAME}").DefaultPolicy.Code)
*/

// PushMode takes workingset key, mode key policy key and performs mode sending to specified devices
func (p Profile) PushMode(workingset string, setMode string, setPolicy string) string {
	setAddress := pushProfileLink(setMode, workingset)
	setBody := pushProfileBody(setPolicy)

	header := contentTypeJSON()
	query, errcPostQuery := rest.Query{}.PostQuery(setAddress, setBody, header, rest.Invisible)
	errc.ErrorCenter(errPushModeQueryTAG, errcPostQuery)
	var modeResponseJSONVariable ModeResponseJSON

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			errJSON := json.Unmarshal(query, &modeResponseJSONVariable)
			errc.ErrorCenter(errPushModeUnmarshalTAG, errJSON)
			return modeResponseJSONVariable[0].Status
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil
}

/*
██████╗ ██╗   ██╗███████╗██╗  ██╗        ███╗   ███╗ ██████╗ ██████╗ ███████╗         █████╗ ██╗   ██╗████████╗ ██████╗
██╔══██╗██║   ██║██╔════╝██║  ██║        ████╗ ████║██╔═══██╗██╔══██╗██╔════╝        ██╔══██╗██║   ██║╚══██╔══╝██╔═══██╗
██████╔╝██║   ██║███████╗███████║        ██╔████╔██║██║   ██║██║  ██║█████╗          ███████║██║   ██║   ██║   ██║   ██║
██╔═══╝ ██║   ██║╚════██║██╔══██║        ██║╚██╔╝██║██║   ██║██║  ██║██╔══╝          ██╔══██║██║   ██║   ██║   ██║   ██║
██║     ╚██████╔╝███████║██║  ██║        ██║ ╚═╝ ██║╚██████╔╝██████╔╝███████╗        ██║  ██║╚██████╔╝   ██║   ╚██████╔╝
╚═╝      ╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝        ╚═╝  ╚═╝ ╚═════╝    ╚═╝    ╚═════╝
*/

/*
	profile.Profile{}.PushModeAuto(profile.Profile{}.GetProfile("{YOUR_PROFILE_NAME}").Code, profile.Profile{}.GetProfile("{YOUR_PROFILE_NAME}").DefaultPolicy.Code, "{YOUR_DEVICE_ID}")
	OR
	var p profile.Profile
	p.PushModeAuto(p.GetProfile("{YOUR_PROFILE_NAME}").Code, p.GetProfile("{YOUR_PROFILE_NAME}").DefaultPolicy.Code, "{YOUR_DEVICE_ID}")
*/

// PushModeAuto takes a mode key, a policy key, and retrieves a device list. send profile to specified devices in specified settings
func (p Profile) PushModeAuto(setMode string, setPolicy string, devicesID ...string) string {
	var (
		workingsets              workingset.Workingset
		devices                  device.Device
		modeResponseJSONVariable ModeResponseJSON
	)

	workingsetKey := workingsets.CreateWorkingset()

	for _, deviceID := range devicesID {
		workingsets.AddDeviceWorkingSet(workingsetKey, devices.DeviceID2Code(deviceID))
	}

	setAddress := pushProfileLink(setMode, workingsetKey)

	query, errPostQuery := rest.Query{}.PostQuery(setAddress, pushProfileBody(setPolicy), contentTypeJSON(), rest.Invisible)
	errc.ErrorCenter(errPushModeAutoQueryTAG, errPostQuery)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			errJSON := json.Unmarshal(query, &modeResponseJSONVariable)
			errc.ErrorCenter(errPushModeAutoUnmarshalTAG, errJSON)
			return modeResponseJSONVariable[0].Status
		}
		return rest.ResponseNotFound
	}
	return rest.ResponseNil
}
