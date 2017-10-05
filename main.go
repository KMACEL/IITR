package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/user"
	"github.com/KMACEL/IITR/writefile"
)

func main() {

	start := time.Now()
	user.Start()
	//user.LoginTenant()
	
	log.Println("Login Status : ", responseConnect)

	//var devices device.Device

	//devices.RefleshGatewayInfo()
	//devices.OSProfileInfo("867377020747089", true, true)
	writefile.CreateFile("test1")
	var openFile1 *os.File
	openFile1 = writefile.OpenFile2("test1", openFile1)

	writefile.CreateFile("test2")
	var openFile2 *os.File
	openFile2 = writefile.OpenFile2("test2", openFile2)

	for i := 0; i < 100; i++ {
		writefile.WriteText2(openFile1, strconv.Itoa(i))
		writefile.WriteText2(openFile2, strconv.Itoa(i+100))
	}

	var pressKey string
	fmt.Scan(&pressKey)

	elapsed := time.Since(start)
	fmt.Printf("Total Time : %s \n", elapsed)
}
