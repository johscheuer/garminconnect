package garminconnect

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

type Weight struct {
	Date  int     `json:"date"`
	Value float64 `json:"weight"`
}

func (gc *GarminConnect) WeightByDate(date time.Time) []Weight {
	params := url.Values{}
	params.Set("from", strconv.FormatInt(time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).UnixNano(), 10))
	params.Set("until", strconv.FormatInt(time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 99, time.UTC).UnixNano(), 10))

	response, err := gc.client.Get("https://connect.garmin.com/modern/proxy/userprofile-service/userprofile/personal-information/weightWithOutbound/filterByDay?" + params.Encode())

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var weights []Weight

	json.NewDecoder(response.Body).Decode(&weights)

	return weights
}