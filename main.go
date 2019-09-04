package main

import (
	"github.com/p-pawel/go-challenge/server"

	"log"
	"net/http"
)


func main() {

	rocketServer := &server.RocketServer{}
	if err := http.ListenAndServe(":3000", rocketServer); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	}
}

