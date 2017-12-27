package operations

import (
	"fmt"
	"time"

	"github.com/KMACEL/IITR/rest/workingset"
	"github.com/KMACEL/IITR/timop"
)

/*
var sendMessageOperation operations.SendMessages

sendMessageOperation.Message = "http://..."
sendMessageOperation.MessageType = "url"
sendMessageOperation.TimeType = "date"
//sendMessageOperation.Time = timop.CreateEpochTime(2017, 12, 27, 13, 48, 30, 0, +3)
sendMessageOperation.WorkingSetKey = "..."
sendMessageOperation.StartTime = timop.CreateEpochTime(2017, 12, 27, 15, 40, 00, 0, +3)
sendMessageOperation.EndTime = timop.CreateEpochTime(2017, 12, 27, 15, 50, 00, 0, +3)
sendMessageOperation.StepMinute = 2
sendMessageOperation.SendFrequencyMinute = 1

sendMessageOperation.Start()

*/

/*
███████╗███████╗███╗   ██╗██████╗         ███╗   ███╗███████╗███████╗███████╗ █████╗  ██████╗ ███████╗███████╗
██╔════╝██╔════╝████╗  ██║██╔══██╗        ████╗ ████║██╔════╝██╔════╝██╔════╝██╔══██╗██╔════╝ ██╔════╝██╔════╝
███████╗█████╗  ██╔██╗ ██║██║  ██║        ██╔████╔██║█████╗  ███████╗███████╗███████║██║  ███╗█████╗  ███████╗
╚════██║██╔══╝  ██║╚██╗██║██║  ██║        ██║╚██╔╝██║██╔══╝  ╚════██║╚════██║██╔══██║██║   ██║██╔══╝  ╚════██║
███████║███████╗██║ ╚████║██████╔╝        ██║ ╚═╝ ██║███████╗███████║███████║██║  ██║╚██████╔╝███████╗███████║
╚══════╝╚══════╝╚═╝  ╚═══╝╚═════╝         ╚═╝     ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚══════╝
*/

//SendMessages is
type SendMessages struct {
	SendMessageRequirements
}

//SendMessageRequirements is
type SendMessageRequirements struct {
	Message             string
	MessageType         string
	TimeType            string
	Time                int64
	StartTime           int64
	EndTime             int64
	StepMinute          int
	SendFrequencyMinute int
	WorkingSetKey       string
}

//Start is
func (s SendMessages) Start(deviceID ...string) {
	var sendMessage workingset.Workingset

	if s.StartTime != 0 && s.EndTime != 0 && s.StepMinute != 0 && s.SendFrequencyMinute != 0 {
		startTime := time.Unix(s.StartTime/1000, 0)
		endTime := time.Unix(s.EndTime/1000, 0)

		fmt.Println("Start Time : ", startTime)
		fmt.Println("End Time : ", endTime)
		fmt.Println("Step Minute : ", s.StepMinute)
		fmt.Println("Send Frequency Minute : ", s.SendFrequencyMinute)

		if endTime.Unix() > startTime.Unix() {
			if startTime.Unix() > time.Now().Unix() {
				fmt.Println("OK")

				totalDay := time.Unix(endTime.Unix()-startTime.Unix(), 0).Day() - 1
				totalHour := time.Unix(endTime.Unix()-startTime.Unix(), 0).Hour() - 2
				totalMinute := time.Unix(endTime.Unix()-startTime.Unix(), 0).Minute()
				fmt.Println("Total Day : ", totalDay)
				fmt.Println("Total Hour : ", totalHour)
				fmt.Println("Total Minute : ", totalMinute)

				if totalDay > 0 {
					totalMinute = totalMinute + (totalDay * 1440)
				}
				if totalHour > 0 {
					totalMinute = totalMinute + (totalHour * 60)
				}

				fmt.Println("Operation Total Time : ", totalMinute)

				totalMessageSize := totalMinute / s.StepMinute
				fmt.Println("Total Message : ", totalMessageSize)

				fmt.Println("Total Send Time (Minute) : ", totalMessageSize*s.SendFrequencyMinute)

				for i := 0; i <= totalMessageSize*s.StepMinute; i = i + s.StepMinute {
					sendTime := timop.CreateEpochTime(startTime.Year(), int(startTime.Month()), startTime.Day(), startTime.Hour(), startTime.Minute()+i, 00, 0, +3)
					sendMessage.SendRichMessage(s.Message, s.MessageType, s.TimeType, sendTime, s.WorkingSetKey, deviceID...)
					time.Sleep(time.Duration(s.SendFrequencyMinute) * time.Minute)
				}

			} else {
				fmt.Println("-Start Time- must be bigger than -Now-")
			}

		} else {
			fmt.Println("-End Time- must be bigger than -Start Time-")
		}
	} else {
		sendMessage.SendRichMessage(s.Message, s.MessageType, s.TimeType, s.Time, s.WorkingSetKey, deviceID...)
	}
}
