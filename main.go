package main

import (
	"fmt"
	"time"

	"github.com/KMACEL/IITR/user"
)

func init() {
	fmt.Println("On Init")
}

func main() {
	defer onExit()
	// Major words at the beginning of functions:
	//http://patorjk.com/software/taag/#p=display&f=ANSI Shadow&t=Your Text

	start := time.Now()

	user.Start()
	user.LoginTenant()

	var pressKey string
	fmt.Scan(&pressKey)

	elapsed := time.Since(start)
	fmt.Printf("Total Time : %s \n", elapsed)
}

func onExit() {
	fmt.Println("On Exit")
}
