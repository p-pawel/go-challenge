package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test(t *testing.T) {

	// given

	testCases := []struct {
		weekdayNo     uint
		launchpadId   uint
		destinationId uint
	}{
		{1, 1, 1},
		{2, 1, 2},
		{3, 1, 3},
		{4, 1, 4},
		{5, 1, 5},
		{6, 1, 6},
		{7, 1, 7},

		{1, 2, 7},
		{2, 2, 1},
		{3, 2, 2},
		{4, 2, 3},
		{5, 2, 4},
		{6, 2, 5},
		{7, 2, 6},

		{1, 3, 6},
		{2, 3, 7},
		{3, 3, 1},
		{4, 3, 2},
		{5, 3, 3},
		{6, 3, 4},
		{7, 3, 5},

		{1, 4, 5},
		{2, 4, 6},
		{3, 4, 7},
		{4, 4, 1},
		{5, 4, 2},
		{6, 4, 3},
		{7, 4, 4},

		{1, 5, 4},
		{2, 5, 5},
		{3, 5, 6},
		{4, 5, 7},
		{5, 5, 1},
		{6, 5, 2},
		{7, 5, 3},

		{1, 6, 3},
		{2, 6, 4},
		{3, 6, 5},
		{4, 6, 6},
		{5, 6, 7},
		{6, 6, 1},
		{7, 6, 2},

		{1, 7, 2},
		{2, 7, 3},
		{3, 7, 4},
		{4, 7, 5},
		{5, 7, 6},
		{6, 7, 7},
		{7, 7, 1},
	}


	for _, tt := range testCases {

		t.Run(fmt.Sprintf("Weekday #%d + Launchpad #%d => Destination #%d\n", tt.weekdayNo, tt.launchpadId, tt.destinationId), func(t *testing.T) {

			// when

			destinationId := calcDestinationForDayAndLaunchpad(tt.weekdayNo, tt.launchpadId)

			// then
			assert.Equal(t, tt.destinationId, destinationId)
		})
	}


}
