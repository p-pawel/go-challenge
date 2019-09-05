package controllers

import (
	"encoding/json"
	"github.com/p-pawel/go-challenge/database"
	"io/ioutil"
	"net/http"
)

func GetBookings(w http.ResponseWriter, r *http.Request) {

	var bookings []database.Booking
	database.DB.Find(&bookings)

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(bookings)
}


func PostBooking(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newBooking database.Booking

	err2 := json.Unmarshal(reqBody, &newBooking)
	if err2 != nil {
		return
	}

	database.DB.Create(&newBooking)

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newBooking)
}