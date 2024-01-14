package main

import (
	"encoding/json"
	fl "flag"
	"fmt"
	f "fmt"
	"io"
	"net/http"
	"os"
	t "time"
)

type Weather struct {
	ResolvedAddress string `json:"resolvedAddress"`
	TimeZone        string `json:"timezone"`
	Days            []struct {
		DatetimeEpoch int64  `json:"datetimeEpoch"`
		Conditions    string `json:"conditions"`
		Descriptions  string `json:"descriptions"`
		Hours         []struct {
			DatetimeEpoch int64 `json:"datetimeEpoch"`

			Conditions string  `json:"conditions"`
			Temp       float64 `json:"temp"`
		} `json:"hours"`
	} `json:"days"`
	CurrentConditions struct {
		DatetimeEpoch int64   `json:"datetimeEpoch"`
		Conditions    string  `json:"conditions"`
		Temp          float64 `json:"temp"`
	} `json:"currentConditions"`
}

func convertToCelcius(fahrenheit float64) float64 {
	// (°F − 32) × 5/9 = °C
	return (fahrenheit - 32.0) * 5 / 9
}
func main() {
	q := "bath"
	var key string

	if len(os.Args) >= 3 {
		q = os.Args[1]
	}
	fl.StringVar(&key, "key", "NULL", "API KEY")
	fl.Parse()

	if key == "" {
		fmt.Println("No API key provided. Exiting.")
		os.Exit(1)
	}

	url := f.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=us&key=%s&contentType=json", q, key)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available or API Key was not supplied")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	location, timezone, currentCondition, hours := weather.ResolvedAddress, weather.TimeZone, weather.CurrentConditions, weather.Days[0].Hours
	temp := convertToCelcius(currentCondition.Temp)

	f.Printf("%s, %s: %.0f°C, %s\n",
		location,
		timezone,
		temp,
		currentCondition.Conditions,
	)
	for _, hour := range hours {
		timeNow := t.Unix(hour.DatetimeEpoch, 0)
		temp := convertToCelcius(hour.Temp)

		if timeNow.Before(t.Now()) {
			continue
		}

		f.Printf("%s - %.0f°C, %s\n",
			timeNow.Format("15:04"),
			temp,
			hour.Conditions,
		)
	}

}
