package service

import (
	"fmt"
	"github.com/p-pawel/go-challenge/database"
)

type BookingValidator interface {
	isValid(booking database.Booking) string
}

type DestinationValidator struct{}

func (d *DestinationValidator) isValid(booking database.Booking) string {

	booking.LaunchDate.Weekday()
	weekdayNo := (uint(booking.LaunchDate.Weekday()) + 6) % 7

	allowedDestination := calcDestinationForDayAndLaunchpad(weekdayNo, booking.LaunchpadId)
	if allowedDestination == booking.Destination.Id {
		return ""
	}
	return fmt.Sprintf("Destinations mismatch, requested destination #%d while for day #%d and pad #%d only destination #%d is allowed", booking.Destination.Id, weekdayNo, booking.LaunchpadId, allowedDestination)

}
