package service

import (
	"github.com/p-pawel/go-challenge/database"
	"log"
)

func FindAllBookings() interface{} {
	var bookings []database.Booking
	database.DB.Find(&bookings)
	return bookings
}


// TODO I guess there must be a better way to more "inline" initialization of validator instances
var d destinationValidator
var l launchpadAvailabilityValidator
var validators = []BookingValidator{&d, &l}

func TryToCreateBooking(newBooking *database.Booking) []string {

	for _, v := range validators {
		result := v.isValid(*newBooking)
		if result != "" {
			log.Printf(result)
			return []string{result}
		}

	}
	database.DB.Create(&newBooking)
	return nil
}


