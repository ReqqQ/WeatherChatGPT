package DomainModelWeather

import (
	"encoding/json"

	"WeatherAPI/Infrastructure/Base"
	InfrastructureWeatherAPI "WeatherAPI/Infrastructure/WeatherAPI"
)

type WeatherDMService struct {
	BaseRequest Base.BaseRequest
	WeatherAPI  InfrastructureWeatherAPI.WeatherAPIRepository
}

func (ws WeatherDMService) GetWeatherInfo(vo WeatherDMVO) json.RawMessage {
	ws.BaseRequest.Init()
	cityCoordinateData := ws.WeatherAPI.GetCityCoordinate(buildWeatherRequest(vo), ws.BaseRequest)
	cityWeatherData := ws.WeatherAPI.GetCityWeather(buildWeatherResponse(cityCoordinateData), ws.BaseRequest)

	return buildRawWeatherCityJson(cityWeatherData)
}
