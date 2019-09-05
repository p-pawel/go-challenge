package database

type Booking struct {
	Id uint `gorm:"type:serial;not null;auto_increment;primary_key"`
	FirstName     string
	LastName      string
	Gender        string
	Birthday      string
	LaunchpadId   uint
	DestinationId uint
	LaunchDate    string
}
