package DomainModelWeather

type WeatherDMService struct{}
type WeatherDMServiceInterface interface {
	GetWeatherInfo(vo WeatherDMVO) string
}

func (ws WeatherDMService) GetWeatherInfo(vo WeatherDMVO) string {

	return ""
}
