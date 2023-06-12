package InfrastructureWeatherAPI

import (
	"WeatherAPI/Infrastructure/Base"
)

const WEATHER_ZIP_ENDPOINT = "http://api.openweathermap.org/geo/1.0/zip?"
const WEATHER_CITY_ENDPOINT = "https://api.openweathermap.org/data/2.5/weather?"

type WeatherAPIRepository struct{}

func (api WeatherAPIRepository) GetCityCoordinate(vo WeatherRequest, baseRequest Base.BaseRequest) []byte {
	request, response := baseRequest.BaseSettings()
	request.SetRequestURI(WEATHER_ZIP_ENDPOINT + buildWeatherStatsParams(vo))
	baseRequest.CallAPI(request, response)

	return response.Body()
}

func (api WeatherAPIRepository) GetCityWeather(cityResponse WeatherCityResponse, baseRequest Base.BaseRequest) []byte {
	request, response := baseRequest.BaseSettings()
	request.SetRequestURI(WEATHER_CITY_ENDPOINT + buildCityWeatherParams(cityResponse))
	baseRequest.CallAPI(request, response)

	return response.Body()
}
