package main

import (
	"fmt"
	"time"

	"github.com/KMACEL/IITR/cases/tests"
	"github.com/KMACEL/IITR/rest"
)

func main() {
	// Major words at the beginning of functions:
	//http://patorjk.com/software/taag/#p=display&f=ANSI Shadow&t=Your Text

	start := time.Now()

	rest.Connect("", "")

	//user.Start()
	//user.LoginTenant()

	//var reportAll reports.DetailAllReport2
	//reportAll.Start("TestReportSX.xlsx", nil)

	/*
		var ping reports.PingControl

		ping.OutputFileName = "test"
		ping.ControlLogFileName = "ping_access.log-20171130"
		ping.Start()
	*/

	var mapsTest tests.MapsTest
	//mapsTest.Start("867377020915728", "com.google.android.apps.maps") // itaksi
	mapsTest.Start("867377020746784", "com.google.android.apps.maps") // test

	var pressKey string
	fmt.Scan(&pressKey)

	elapsed := time.Since(start)
	fmt.Printf("Total Time : %s \n", elapsed)
}
