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

	var pressKey string
	fmt.Scan(&pressKey)

	elapsed := time.Since(start)
	fmt.Printf("Total Time : %s \n", elapsed)
}
