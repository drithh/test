package main

import (
	"fmt"
	"time"
)

func calculateTemperatures(weatherData []WeatherData, mode Mode) map[string]float64 {
	dayTempMap := make(map[string]float64)

	switch mode {
	case StartDay:
		dayTempMap = calculateStartDay(weatherData)
	case Midday:
		dayTempMap = calculateMidday(weatherData)
	case Average:
		dayTempMap = calculateAverage(weatherData)
	default:
		fmt.Println("Invalid mode")
	}

	return dayTempMap
}

// Mode 1: StartDay - Calculate the first forecast entry of each day
func calculateStartDay(weatherData []WeatherData) map[string]float64 {
	dayTempMap := make(map[string]float64)

	for _, data := range weatherData {
		timestamp := time.Unix(data.Dt, 0)
		day := timestamp.Format("2006-01-02")

		// Store the first temperature entry of the day
		if _, exists := dayTempMap[day]; !exists {
			dayTempMap[day] = data.Main.Temp
		}
	}
	return dayTempMap
}

// Mode 2: Midday - Calculate the temperature closest to 12:00 PM each day
func calculateMidday(weatherData []WeatherData) map[string]float64 {
	dayTempMap := make(map[string]float64)
	middayHour := 12
	for _, data := range weatherData {
		timestamp := time.Unix(data.Dt, 0)
		hour := timestamp.Hour()

		if hour >= middayHour-1 && hour <= middayHour+1 {
			day := timestamp.Format("2006-01-02")
			dayTempMap[day] = data.Main.Temp
		}
	}
	return dayTempMap
}

// Mode 3: Average - Calculate the average temperature of the day
func calculateAverage(weatherData []WeatherData) map[string]float64 {
	dayTempMap := make(map[string]float64)
	dayCountMap := make(map[string]int)

	for _, data := range weatherData {
		timestamp := time.Unix(data.Dt, 0)
		day := timestamp.Format("2006-01-02")

		dayTempMap[day] += data.Main.Temp
		dayCountMap[day]++
	}

	for day, totalTemp := range dayTempMap {
		dayTempMap[day] = totalTemp / float64(dayCountMap[day])
	}
	return dayTempMap
}
