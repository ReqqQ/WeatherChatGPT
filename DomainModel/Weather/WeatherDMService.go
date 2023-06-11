package DomainModelWeather

import (
	"WeatherAPI/Infrastructure/WeatherAPI"
	"encoding/json"
)

type WeatherDMService struct {
	WeatherAPI WeatherAPI.WeatherAPIRepository
}

func (ws WeatherDMService) GetWeatherInfo(vo WeatherDMVO) json.RawMessage {
	ws.WeatherAPI.Init()
	data := ws.WeatherAPI.GetCityCoordinate(buildWeatherRequest(vo))
	weatherResponse := buildWeatherResponse(data)

	data = ws.WeatherAPI.GetCityWeather(weatherResponse)

	return buildRawWeatherCityJson(data)
}
