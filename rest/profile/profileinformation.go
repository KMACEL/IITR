package profile

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ███████╗████████╗        ██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗        ██╗     ██╗███████╗████████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝        ██║     ██║██╔════╝╚══██╔══╝
██║  ███╗█████╗     ██║           ██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗          ██║     ██║███████╗   ██║
██║   ██║██╔══╝     ██║           ██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝          ██║     ██║╚════██║   ██║
╚██████╔╝███████╗   ██║           ██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗        ███████╗██║███████║   ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝        ╚══════╝╚═╝╚══════╝   ╚═╝
*/

/*
	profile.Profile{}.GetProfileList()
*/

// GetProfileList lists all the profiles that are registered.
func (p Profile) GetProfileList() ResponseProfileListJSON {
	setQueryAddress := getProfileListLink()
	visualFlag := rest.Invisible

	query, err := rest.Query{}.GetQuery(setQueryAddress, visualFlag)
	errc.ErrorCenter(errGetProfileListQueryTAG, err)
	var profileList ResponseProfileListJSON

	if query != nil {
		errJSON := json.Unmarshal(query, &profileList)
		errc.ErrorCenter(errGetProfileListUnmarshalTAG, errJSON)
		return profileList
	}
	return profileList
}

/*
 ██████╗ ███████╗████████╗        ██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝
██║  ███╗█████╗     ██║           ██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗
██║   ██║██╔══╝     ██║           ██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝
╚██████╔╝███████╗   ██║           ██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝
*/

/*
	profile.Profile{}.GetProfile("{YOUR_PROFILE_NAME}").Code
*/

// GetProfile retrieves a profile name and returns that profile information.
func (p Profile) GetProfile(setProfileName string) ResponseProfileJSON {
	setQueryAddress := getProfileLink(setProfileName)
	visualFlag := rest.Invisible

	query, errGetQuery := rest.Query{}.GetQuery(setQueryAddress, visualFlag)
	errc.ErrorCenter(errGetProfileQueryTAG, errGetQuery)

	var getProfileArray ResponseProfileJSONArray
	var getProfile ResponseProfileJSON

	if query != nil {
		errJSON := json.Unmarshal(query, &getProfileArray)
		errc.ErrorCenter(errGetProfileUnmarshalTAG, errJSON)
		getProfile = getProfileArray[0]
		return getProfile
	}
	return getProfile
}

/*
 ██████╗ ███████╗████████╗        ██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗        ██╗███╗   ██╗        ██████╗  ██████╗ ██╗     ██╗ ██████╗██╗   ██╗        ██╗     ██╗███████╗████████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝        ██║████╗  ██║        ██╔══██╗██╔═══██╗██║     ██║██╔════╝╚██╗ ██╔╝        ██║     ██║██╔════╝╚══██╔══╝
██║  ███╗█████╗     ██║           ██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗          ██║██╔██╗ ██║        ██████╔╝██║   ██║██║     ██║██║      ╚████╔╝         ██║     ██║███████╗   ██║
██║   ██║██╔══╝     ██║           ██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝          ██║██║╚██╗██║        ██╔═══╝ ██║   ██║██║     ██║██║       ╚██╔╝          ██║     ██║╚════██║   ██║
╚██████╔╝███████╗   ██║           ██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗        ██║██║ ╚████║        ██║     ╚██████╔╝███████╗██║╚██████╗   ██║           ███████╗██║███████║   ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝        ╚═╝╚═╝  ╚═══╝        ╚═╝      ╚═════╝ ╚══════╝╚═╝ ╚═════╝   ╚═╝           ╚══════╝╚═╝╚══════╝   ╚═╝
*/
//profile.Profile{}.GetProfileInPolicyList(profile.Profile{}.GetProfile("debug").Code)

// GetProfileInPolicyList is
func (p Profile) GetProfileInPolicyList(setProfileCode string) ResponseGetProfileInPolicyListJSON {
	setQueryAddress := getProfileInPolicyListLink(setProfileCode)
	visualFlag := rest.Invisible

	query, errGetQuery := rest.Query{}.GetQuery(setQueryAddress, visualFlag)
	errc.ErrorCenter(errGetProfileQueryTAG, errGetQuery)

	var responseGetProfileInPolicyListJSON ResponseGetProfileInPolicyListJSON
	if query != nil {
		errJSON := json.Unmarshal(query, &responseGetProfileInPolicyListJSON)
		errc.ErrorCenter(errGetProfileUnmarshalTAG, errJSON)

		return responseGetProfileInPolicyListJSON
	}
	return responseGetProfileInPolicyListJSON
}

/*
 ██████╗ ███████╗████████╗        ███████╗███████╗██╗     ███████╗ ██████╗████████╗        ██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗        ██╗███╗   ██╗        ██████╗  ██████╗ ██╗     ██╗ ██████╗██╗   ██╗         ██████╗ ██████╗ ██████╗ ███████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔════╝██╔════╝██║     ██╔════╝██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝        ██║████╗  ██║        ██╔══██╗██╔═══██╗██║     ██║██╔════╝╚██╗ ██╔╝        ██╔════╝██╔═══██╗██╔══██╗██╔════╝
██║  ███╗█████╗     ██║           ███████╗█████╗  ██║     █████╗  ██║        ██║           ██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗          ██║██╔██╗ ██║        ██████╔╝██║   ██║██║     ██║██║      ╚████╔╝         ██║     ██║   ██║██║  ██║█████╗
██║   ██║██╔══╝     ██║           ╚════██║██╔══╝  ██║     ██╔══╝  ██║        ██║           ██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝          ██║██║╚██╗██║        ██╔═══╝ ██║   ██║██║     ██║██║       ╚██╔╝          ██║     ██║   ██║██║  ██║██╔══╝
╚██████╔╝███████╗   ██║           ███████║███████╗███████╗███████╗╚██████╗   ██║           ██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗        ██║██║ ╚████║        ██║     ╚██████╔╝███████╗██║╚██████╗   ██║           ╚██████╗╚██████╔╝██████╔╝███████╗
 ╚═════╝ ╚══════╝   ╚═╝           ╚══════╝╚══════╝╚══════╝╚══════╝ ╚═════╝   ╚═╝           ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝        ╚═╝╚═╝  ╚═══╝        ╚═╝      ╚═════╝ ╚══════╝╚═╝ ╚═════╝   ╚═╝            ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝
*/
// 	profileCode, policyCode := profile.Profile{}.GetSelectProfileInPolicyCode(ModeName, PolicyName)

// GetSelectProfileInPolicyCode is
func (p Profile) GetSelectProfileInPolicyCode(setProfileName string, setPolicyName string) (profileCode string, policyCode string) {
	profileCode = p.GetProfile(setProfileName).Code
	responseGetProfileInPolicyListJSON := p.GetProfileInPolicyList(profileCode)
	for _, policy := range responseGetProfileInPolicyListJSON {
		if policy.PolicyProfile.Name == setPolicyName {
			return profileCode, policy.Code
		}
	}
	return profileCode, "Not Found"
}
