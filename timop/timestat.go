package timop

import (
	"strconv"
	"time"
)

// GetHour is
func GetHour() int {
	return time.Now().Hour()
}

//GetMinute is
func GetMinute() int {
	return time.Now().Minute()
}

//GetSecond is
func GetSecond() int {
	return time.Now().Second()
}

// GetTimeNamesFormat is
func GetTimeNamesFormat() string {
	return strconv.Itoa(time.Now().Day()) + "_" + time.Now().Month().String() + "___" + strconv.Itoa(time.Now().Hour()) + "_" + strconv.Itoa(time.Now().Minute())
}
