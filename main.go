package main

import (
	"fmt"
	"time"

	"github.com/KMACEL/IITR/user"
)

func main() {

	start := time.Now()
	user.Start()
	user.LoginTenant()

	//operations.TimeOp{}.Start()
	//cases.DetailReport{}.Start("ApplicationPackage_"+timop.GetTimeNamesFormat()+".xlsx", "tr.com.innology.taksipager", "com.ardic.android.iot.appblocker", "com.android.launcher3", "com.google.android.apps.maps", "com.streamaxtech.mdvr.direct")
	//cases.DeviceInformation{}.Start()
	//device.Device{}.ApplicationInfo("867377021052356", true, true)
	//cases.OfflineLog{}.Start("867377020740787")

	/*
		var offlineLog cases.OfflineLog
		offlineLog.DeviceID = []string{"867377020747089"}
		cases.Test(offlineLog)
	*/

	//cases.DromTest{}.Start("867377020746784")

	var pressKey string
	fmt.Scan(&pressKey)

	elapsed := time.Since(start)
	fmt.Printf("Total Time : %s \n", elapsed)
}
