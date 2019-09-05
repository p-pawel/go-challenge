package controllers

import (
	"encoding/json"
	"net/http"
)

var GetBookings = func(w http.ResponseWriter, r *http.Request) {

	var data = make([]string, 0)

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}