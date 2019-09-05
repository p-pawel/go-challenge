package server

import (
	"github.com/gorilla/mux"
	"github.com/p-pawel/go-challenge/controllers"
	"log"
	"net/http"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/booking", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/booking", controllers.PostBooking).Methods("POST")
	return router
}

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
