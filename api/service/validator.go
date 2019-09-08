package service

import (
	"fmt"
	"github.com/p-pawel/go-challenge/database"
	"github.com/p-pawel/go-challenge/spacex"
	"log"
	"net/http"
)

type BookingValidator interface {
	isValid(booking database.Booking) string
}

type destinationValidator struct{}

func (v *destinationValidator) isValid(booking database.Booking) string {

	booking.LaunchDate.Weekday()
	weekdayNo := (uint(booking.LaunchDate.Weekday()) + 6) % 7

	allowedDestination := calcDestinationForDayAndLaunchpad(weekdayNo, booking.LaunchpadId)
	if allowedDestination != booking.DestinationId {
		return fmt.Sprintf("Destinations mismatch, requested destination #%v while for day #%v and pad #%v only destination #%v is allowed", booking.DestinationId, weekdayNo, booking.LaunchpadId, allowedDestination)
	} else {
		return ""
	}

}

type launchpadAvailabilityValidator struct{}

func (v *launchpadAvailabilityValidator) isValid(booking database.Booking) string {

	var bookings []database.Booking
	database.DB.Where("date(launch_date) = date(?) and launchpad_id = ?", booking.LaunchDate, booking.LaunchpadId).Find(&bookings)

	if len(bookings) > 0 {
		return fmt.Sprintf("Launchpad #%d for day %v is not available", booking.LaunchpadId, booking.LaunchDate)
	}

	return ""
}

type overlapWithSpaceXValidator struct{}

var api = spacex.API{
	Client: &http.Client{},
	URL:    "https://api.spacexdata.com/v3",
}

func (v *overlapWithSpaceXValidator) isValid(booking database.Booking) string {

	launches, _ := api.GetUpcomingLaunches()
	launchpad := FindOneLaunchpad(booking.LaunchpadId)

	for _, l := range launches {
		if l.LaunchSite.SiteId != (*launchpad).SpacexSiteId {
			continue
		}
		if l.LaunchDateUTC.Year() != booking.LaunchDate.Year() {
			continue
		}
		fmt.Println(l.LaunchDateUTC.YearDay())
		fmt.Println(booking.LaunchDate.YearDay())
		if l.LaunchDateUTC.YearDay() != booking.LaunchDate.YearDay() {
			continue
		}

		return fmt.Sprintf("Launchpad #%d for day %v is used by SpaceX", booking.LaunchpadId, booking.LaunchDate)

	}

	log.Println(launches)

	return ""
}
