package main

import (
	"fmt"
	"github.com/KMACEL/IITR/user"
	"time"
)

func main() {
	// Major words at the beginning of functions:
	//http://patorjk.com/software/taag/#p=display&f=ANSI Shadow&t=Your Text

	start := time.Now()

	user.Start()
	user.LoginTenant()

	/*var reportAll reports.DetailAllReport2
	reportAll.Start("TestReportSX.xlsx", "tr.com.innology.itaksi.taxi.preprod", "tr.com.innology.itaksi.taxi")
	*/

	var pressKey string
	fmt.Scan(&pressKey)

	elapsed := time.Since(start)
	fmt.Printf("Total Time : %s \n", elapsed)
}
