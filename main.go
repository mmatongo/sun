package main

import (
	f "fmt"
	"io"
	net "net/http"
)

type Weather struct {
	ResolvedAddress string `json:"resolvedAddress"`
	TimeZone        string `json:"timezone"`
	Days            []struct {
		DateTime     string `json:"datetime"`
		Conditions   string `json:"conditions"`
		Descriptions string `json:"descriptions"`
		Hours        []struct {
			DateTime   string `json:"datetime"`
			Conditions string `json:"conditions"`
		} `json:"hours"`
	} `json:"days"`
	CurrentConditions struct {
		DateTime   string `json:"datetime"`
		Conditions string `json:"conditions"`
	} `json:"currentConditions"`
}

func main() {
	res, err := net.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/bath?unitGroup=us&key=&contentType=json")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	f.Println(string(body))
}
