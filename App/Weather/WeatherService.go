package AppWeather

type WeatherService struct{}
type WeatherServiceInterface interface {
	GetWeatherInfo(command BuildWeatherCommand) string
}

func (ws WeatherService) GetWeatherInfo(command BuildWeatherCommand) string {
	//baseAPP := ServiceLocator.App
	//factory := baseAPP.Application.WeatherApp.WeatherFactory
	//weatherDomainService := baseAPP.DomainModel.WeatherDomain.WeatherService
	//
	//vo := factory.BuildWeatherDMVO(command.GetZipCode(), command.GetCountryCode())

	//return weatherDomainService.GetWeatherInfo(vo)
	return ""
}
