package main

import (
	"context"
	"fmt"
	"log"
)

type WeatherForecastResponse struct {
	Cod     string        `json:"cod"`
	Message int           `json:"message"`
	Cnt     int           `json:"cnt"`
	List    []WeatherData `json:"list"`
	City    CityData      `json:"city"`
}

type WeatherData struct {
	Dt         int64         `json:"dt"`
	Main       MainData      `json:"main"`
	Weather    []WeatherInfo `json:"weather"`
	Clouds     CloudsData    `json:"clouds"`
	Wind       WindData      `json:"wind"`
	Visibility int           `json:"visibility"`
	Pop        float64       `json:"pop"`
	Rain       RainData      `json:"rain"`
	Sys        SysData       `json:"sys"`
	DtTxt      string        `json:"dt_txt"`
}

type MainData struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type WeatherInfo struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CloudsData struct {
	All int `json:"all"`
}

type WindData struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type RainData struct {
	Hours3 float64 `json:"3h"`
}

type SysData struct {
	Pod string `json:"pod"`
}

type CityData struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Coord      CoordData `json:"coord"`
	Country    string    `json:"country"`
	Population int       `json:"population"`
	Timezone   int       `json:"timezone"`
	Sunrise    int64     `json:"sunrise"`
	Sunset     int64     `json:"sunset"`
}

type CoordData struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// getWeatherForecast fetches the weather data from OpenWeatherMap API.
func getWeatherForecast(
	apiKey string,
	city string,
) (*WeatherForecastResponse, error) {
	ctx := context.Background()

	// api.openweathermap.org/data/2.5/forecast?q={city}&appid={API key}
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?units=metric&q=%s&appid=%s", city, apiKey)

	weatherForecastResponses, err := fetch[WeatherForecastResponse](ctx, url)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return &weatherForecastResponses, nil
}
