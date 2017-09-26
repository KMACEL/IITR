package user

import (
	"fmt"

	"github.com/KMACEL/IITR/rest"
)

//Start is
func Start() {
	welcome()
}

//LoginTenant is
func LoginTenant() {
	var username string
	var password string
retry:

	fmt.Print("Username : ")
	fmt.Scan(&username)

	fmt.Print("password: ")
	fmt.Print("\033[8m") // Hide input
	fmt.Scan(&password)
	fmt.Print("\033[28m")

	responseConnect := rest.Connect(username, password)
	if responseConnect {
		entryParameter()
	} else {
		goto retry
	}
}

func entryParameter() {
	var operation string
	for {
		fmt.Print("Entry Operation Parameter : ")
		fmt.Scan(&operation)
		if operation == "report" {
			report()
		} else if operation == "test" {
			test()
		}
	}
}
