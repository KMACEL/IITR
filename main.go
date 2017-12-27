package main

import (
	"fmt"
	"gerrit/DeviceReportApp/user"
	"time"
)

func main() {
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
