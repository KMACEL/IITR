package profile

import (
	"fmt"

	"github.com/KMACEL/IITR/rest"
)

//GetProfileList is
func (p Profile) GetProfileList() {
	setQueryAdress := getProfileListLink()
	vasualFlag := rest.Invisible
	query, _ := queryVariable.GetQuery(setQueryAdress, vasualFlag)
	fmt.Println(string(query))
}

//GetProfile is
func (p Profile) GetProfile(setProfileName string) {
	setQueryAdress := getProfileLink(setProfileName)
	vasualFlag := rest.Invisible
	query, _ := queryVariable.GetQuery(setQueryAdress, vasualFlag)
	fmt.Println(string(query))
}
