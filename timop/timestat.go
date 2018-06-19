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
	return strconv.Itoa(time.Now().Day()) + "_" + strconv.Itoa(int(time.Now().Month())) + "_" + strconv.Itoa(time.Now().Year()) + "___" + strconv.Itoa(time.Now().Hour()) + "_" + strconv.Itoa(time.Now().Minute())
}

// GetTimeNamesFormat is
func GetTimeNamesFormatDays() string {
	return strconv.Itoa(time.Now().Day()) + "_" + strconv.Itoa(int(time.Now().Month())) + "_" + strconv.Itoa(time.Now().Year())
}

//CreateEpochTime is
func CreateEpochTime(year int, month int, day int, hour int, minute int, second int, miliSecond int, timeZone int) int64 {
	var (
		monthT time.Month
	)
	monthT = time.Month(month)
	hour = hour - timeZone
	return time.Date(year, monthT, day, hour, minute, second, miliSecond, time.UTC).Unix() * 1000
}
