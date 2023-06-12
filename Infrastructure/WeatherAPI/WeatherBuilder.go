package InfrastructureWeatherAPI

import (
	"fmt"
	"net/url"
	"os"
)

type WeatherRequest struct {
	ZipCode     string
	CountryCode string
}

type WeatherCityResponse struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

func buildWeatherStatsParams(vo WeatherRequest) string {
	params := url.Values{}
	params.Set("zip", vo.ZipCode+","+vo.CountryCode)
	params.Set("appid", os.Getenv("WEATHER_API_KEY"))

	return params.Encode()
}
func buildCityWeatherParams(cityResponse WeatherCityResponse) string {
	params := url.Values{}
	params.Set("units", "metric")
	params.Set("lat", fmt.Sprintf("%.2f", cityResponse.Lat))
	params.Set("lon", fmt.Sprintf("%.2f", cityResponse.Lon))
	params.Set("appid", os.Getenv("WEATHER_API_KEY"))

	return params.Encode()
}
