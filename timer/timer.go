package timer

import (
	"git.garena.com/russell.chanxl/dinner/utils"
	"log"
	"time"
)

const (
	bookingHour int = 12
	bookingMin  int = 30
	bookingSec  int = 00
)

// SetRudimentaryTimer means you will have to leave the program running; if you close the terminal or terminate the program, it's over -- no food for you!
func SetRudimentaryTimer() {

	s := utils.NewSpinner()
	s.Prefix = "Waiting to book at 1230pm (please don't close this window) "
	s.Start()

	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), bookingHour, bookingMin, bookingSec, 0, time.Local)
	if nextTick.Before(now) {
		log.Fatalf("booking time has passed, do it manually. Chances are, there's probably nothing tasty left\n")
	}
	t := time.NewTimer(nextTick.Sub(now))

	// blocks until 12:30pm
	<-t.C

	s.Stop()
}
