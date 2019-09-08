package database

import "time"

type Booking struct {
	Id            uint `gorm:"type:serial;not null;auto_increment;primary_key"`
	FirstName     string
	LastName      string
	Gender        string
	Birthday      string
	LaunchpadId   uint
	//Launchpad     Launchpad `gorm:"ForeignKey:LaunchpadId;AssociationForeignKey:Id"`
	DestinationId uint
	//Destination   Destination `gorm:"ForeignKey:DestinationId;AssociationForeignKey:Id"`
	LaunchDate    time.Time
}
// TODO
// I gave up with these many-to-one relationships - definitely what is described in http://gorm.io/docs/has_one.html or http://jinzhu.me/gorm/associations.html#has-one
// is not enough to generate schema by CreateTable, and with uncommented uses of structs they are not joined


type Destination struct {
	Id   uint `gorm:"type:serial;not null;auto_increment;primary_key"`
	Name string
}

type Launchpad struct {
	Id           uint `gorm:"type:serial;not null;auto_increment;primary_key"`
	Name         string
	SpacexSiteId string
}
