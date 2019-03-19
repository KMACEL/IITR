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
	//	return strconv.Itoa(time.Now().Day()) + "_" + strconv.Itoa(int(time.Now().Month())) + "_" + strconv.Itoa(time.Now().Year()) + "___" + strconv.Itoa(time.Now().Hour()) + "_" + strconv.Itoa(time.Now().Minute())
	month := strconv.Itoa(int(time.Now().Month()))
	if int(time.Now().Month()) < 10 {
		month = "0" + strconv.Itoa(int(time.Now().Month()))
	}

	day := strconv.Itoa(time.Now().Day())
	if time.Now().Day() < 10 {
		day = "0" + strconv.Itoa(time.Now().Day())
	}

	hour := strconv.Itoa(time.Now().Hour())
	if time.Now().Hour() < 10 {
		hour = "0" + strconv.Itoa(time.Now().Hour())
	}

	minute := strconv.Itoa(time.Now().Minute())
	if time.Now().Minute() < 10 {
		minute = "0" + strconv.Itoa(time.Now().Minute())
	}

	return strconv.Itoa(time.Now().Year()) + "-" + month + "-" + day + "_" + hour + "-" + minute
}

// GetTimeNamesFormatDays is
func GetTimeNamesFormatDays() string {
	return strconv.Itoa(time.Now().Day()) + "_" + strconv.Itoa(int(time.Now().Month())) + "_" + strconv.Itoa(time.Now().Year())
}

//GetTimeNamesFormatDaysTYPE2 is
func GetTimeNamesFormatDaysTYPE2() string {
	return time.Now().Format("2006_01_02")
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
