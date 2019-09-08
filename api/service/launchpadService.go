package service

import (
	"github.com/p-pawel/go-challenge/database"
)

func FindOneLaunchpad(id uint) *database.Launchpad {
	var launchpads []database.Launchpad
	database.DB.
		Where("id = ?", id).
		Find(&launchpads).
		Limit(1)

	if len(launchpads) > 0 {
		return &launchpads[0]
	} else {
		return nil
	}
}
