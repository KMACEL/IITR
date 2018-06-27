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
	var (
		username string
		password string
	)

	fmt.Println("********************************************************************************")
	fmt.Println("Welcome to IITR...")
	fmt.Println("IITR is an Ardic project. This project was written to receive reports and test the devices from Ardic platform IoT-Ignite.")
	fmt.Println("To log in, enter the 'username' and 'password' of your IoT-Ignite account you want to process ...")
	fmt.Println("********************************************************************************")

retry:
	fmt.Print("Username : ")
	fmt.Scan(&username)

	fmt.Print("password: ")
	fmt.Print("\033[8m") // Hide input
	fmt.Scan(&password)
	fmt.Print("\033[28m")

	fmt.Println("********************************************************************************")
	fmt.Println("********************************************************************************")

	responseConnect, _ := rest.Connect(username, password)
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
