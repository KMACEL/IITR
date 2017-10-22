package main

import (
	"fmt"
	"time"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/cases/operations"
)

func main() {
	// Major words at the beginning of functions:
	//http://patorjk.com/software/taag/#p=display&f=ANSI Shadow&t=Your Text

	start := time.Now()

	rest.Connect("", "")
	//user.Start()
	//user.LoginTenant()

	var otaOperation operations.OtaOp

	otaDeviceArray := operations.OtaDeviceArray{
		operations.OtaBaseOp{Imei: "867377020740787"},
		operations.OtaBaseOp{Imei: "867377020747089"}}

	otaOperation.Start(otaDeviceArray)
	//workingset.Workingset{}.PushApplications("DD76AFEA-E0A3-4B61-97CA-509B66A884E1", false, "867377020740787","867377020747089")

	var pressKey string
	fmt.Scan(&pressKey)

	elapsed := time.Since(start)
	fmt.Printf("Total Time : %s \n", elapsed)
}
