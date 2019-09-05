package controllers

import (
	"encoding/json"
	"github.com/p-pawel/go-challenge/database"
	"net/http"
)

var GetBookings = func(w http.ResponseWriter, r *http.Request) {

	var bookings []database.Booking
	database.DB.Find(&bookings)

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(bookings)
}