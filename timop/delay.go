package timop

import "time"

// Delay is
type Delay struct{}

//M is
func (d Delay) M(delayValue int) {
	for delayTime := 0; delayTime < delayValue; delayTime++ {
		time.Sleep(1 * time.Minute)
	}
}

//S is
func (d Delay) S(delayValue int) {
	for delayTime := 0; delayTime < delayValue; delayTime++ {
		time.Sleep(1 * time.Second)
	}
}
