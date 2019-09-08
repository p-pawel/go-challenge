package spacex

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSpaceXConsumerAPI(t *testing.T) {
	t.Run("Should parse response from (test) server", func(t *testing.T) {

		// given
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			//_, _ = rw.Write([]byte(`{}`))
			w.WriteHeader(200)
			w.Write([]byte(`[{"flightNumber":123,"launch_date_utc":"2019-10-17T00:00:00.000Z","launch_site":{"site_id":"ccafs_slc_40"}},{"flightNumber":456, "launch_date_utc":"2019-10-19T00:00:00.000Z", "launch_site": {"site_id": "ksc_lc_39a"}}]`))
		}))
		defer server.Close()
		api := API{server.Client(), server.URL}

		// when
		launches, _ := api.GetUpcomingLaunches()

		// then
		assert.Equal(t, time.Date(2019, 10, 17, 0, 0, 0, 0, time.FixedZone("", 0)).Unix(), launches[0].LaunchDateUTC.Unix())
		assert.Equal(t, "ccafs_slc_40", launches[0].LaunchSite.SiteId)
		assert.Equal(t, time.Date(2019, 10, 19, 0, 0, 0, 0, time.FixedZone("", 0)).Unix(), launches[1].LaunchDateUTC.Unix())
		assert.Equal(t, "ksc_lc_39a", launches[1].LaunchSite.SiteId)

	})
}
