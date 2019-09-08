package service

import (
	"fmt"
	"github.com/p-pawel/go-challenge/database"
)

type BookingValidator interface {
	isValid(booking database.Booking) string
}

type destinationValidator struct{}

func (v *destinationValidator) isValid(booking database.Booking) string {

	booking.LaunchDate.Weekday()
	weekdayNo := (uint(booking.LaunchDate.Weekday()) + 6) % 7

	allowedDestination := calcDestinationForDayAndLaunchpad(weekdayNo, booking.LaunchpadId)
	if allowedDestination != booking.Destination.Id {
		return fmt.Sprintf("Destinations mismatch, requested destination #%v while for day #%v and pad #%v only destination #%v is allowed", booking.Destination.Id, weekdayNo, booking.LaunchpadId, allowedDestination)
	} else {
		return ""
	}

}

type launchpadAvailabilityValidator struct{}

func (v *launchpadAvailabilityValidator) isValid(booking database.Booking) string {

	var bookings []database.Booking
	database.DB.Where("date(launch_date) = date(?) and launchpad_id = ?", booking.LaunchDate, booking.LaunchpadId).Find(&bookings)

	if len(bookings)>0 {
		return fmt.Sprintf("Launchpad #%d for day %v is not available", booking.LaunchpadId, booking.LaunchDate)
	}

	return ""
}
