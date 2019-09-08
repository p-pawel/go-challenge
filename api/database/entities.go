package database

import "time"

type Booking struct {
	Id            uint `gorm:"type:serial;not null;auto_increment;primary_key"`
	FirstName     string
	LastName      string
	Gender        string
	Birthday      string
	LaunchpadId   uint
	Destination   Destination
	LaunchDate    time.Time
}

type Destination struct {
	Id   uint `gorm:"type:serial;not null;auto_increment;primary_key"`
	Name string
}
