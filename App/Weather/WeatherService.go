package AppWeather

import (
	DomainModelWeather "WeatherAPI/DomainModel/Weather"
	"encoding/json"
)

type WeatherService struct {
	WeatherFactory   WeatherFactory
	WeatherDMService DomainModelWeather.WeatherDMService
}

func (ws WeatherService) GetWeatherInfo(command BuildWeatherCommand) json.RawMessage {
	vo := ws.WeatherFactory.BuildWeatherDMVO(command.GetZipCode(), command.GetCountryCode())

	return ws.WeatherDMService.GetWeatherInfo(vo)
}
