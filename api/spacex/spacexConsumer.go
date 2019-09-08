package spacex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type API struct {
	Client *http.Client
	URL    string
}

type LaunchSite struct {
	SiteId string `json:"site_id"`
}

type LaunchesUpcomingResponseItem struct {
	LaunchDateUTC time.Time`json:"launch_date_utc"`
	LaunchSite   LaunchSite `json:"launch_site"`
}

func (api *API) GetUpcomingLaunches() ([]LaunchesUpcomingResponseItem, error) {
	url := api.URL + "/launches/upcoming"
	resp, err := api.Client.Get(url)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	var items []LaunchesUpcomingResponseItem
	err = json.Unmarshal(body, &items)

	return items, err
}
