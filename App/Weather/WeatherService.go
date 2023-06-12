package AppWeather

import (
	DomainModelChatGPT "WeatherAPI/DomainModel/ChatGPT"
	DomainModelWeather "WeatherAPI/DomainModel/Weather"
)

type WeatherService struct {
	WeatherFactory   WeatherFactory
	WeatherDMService DomainModelWeather.WeatherDMService
	ChatGPTService   DomainModelChatGPT.ChatGPTService
}

func (ws WeatherService) GetWeatherInfo(command BuildWeatherCommand) WeatherDTO {
	vo := ws.WeatherFactory.BuildWeatherDMVO(command.GetZipCode(), command.GetCountryCode())
	cityWeatherJsonData := ws.WeatherDMService.GetWeatherInfo(vo)
	chatGPTMessage := ws.ChatGPTService.ParseCityWeatherJsonToHumanLang(cityWeatherJsonData)

	return ws.WeatherFactory.BuildWeatherDTO(chatGPTMessage)
}
