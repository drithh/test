package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

const (
	apiKey = "f6273965b4a43db14d02c2ddf9c2cd3f"
	city   = "Jakarta"
)

func displayResults(dayTempMap map[string]float64, city string, mode string) {
	// sort the keys, since the map is unsorted
	keys := make([]string, 0)
	for k := range dayTempMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Printf("5-Day Weather Forecast for %s (Mode: %s)\n", city, mode)
	for _, k := range keys {
		date, err := time.Parse("2006-01-02", k)
		if err != nil {
			log.Fatalf("Error parsing date: %v", err)
		}
		formattedDate := date.Format("Mon 02 Jan 2006")
		fmt.Printf("%s: %.2f°C\n", formattedDate, dayTempMap[k])
	}
}

type Mode int

const (
	StartDay Mode = iota
	Midday
	Average
)

func main() {
	mode := Midday
	weatherForecast, err := getWeatherForecast(apiKey, city)
	if err != nil {
		log.Fatalf("Could not get weather: %v", err)
	}
	dayTempMap := calculateTemperatures(weatherForecast.List, mode)

	strMode := ""
	switch mode {
	case StartDay:
		strMode = "StartDay"
	case Midday:
		strMode = "Midday"
	case Average:
		strMode = "Average"
	}

	displayResults(dayTempMap, city, strMode)
}
