package AppWeather

type WeatherService struct {
	WeatherFactory WeatherFactory
}
type WeatherServiceInterface interface {
	GetWeatherInfo(command BuildWeatherCommand) string
}

func (ws WeatherService) GetWeatherInfo(command BuildWeatherCommand) string {
	ws.WeatherFactory.BuildWeatherDMVO(command.GetZipCode(), command.GetCountryCode())

	//return weatherDomainService.GetWeatherInfo(vo)
	return ""
}
