package profile

import (
	"fmt"

	"github.com/KMACEL/IITR/rest"
)

//GetProfileList is
func (p Profile) GetProfileList() {
	setQueryAddress := getProfileListLink()
	visualFlag := rest.Invisible
	query, _ := queryVariable.GetQuery(setQueryAddress, visualFlag)
	fmt.Println(string(query))
}

//GetProfile is
func (p Profile) GetProfile(setProfileName string) string {
	setQueryAddress := getProfileLink(setProfileName)
	visualFlag := rest.Invisible
	query, _ := queryVariable.GetQuery(setQueryAddress, visualFlag)
	fmt.Println(string(query))

	return string(query)
}
