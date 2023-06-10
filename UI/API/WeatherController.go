package API

import (
	AppWeather "WeatherAPI/App/Weather"
	"github.com/gofiber/fiber/v2"
)

type WeatherController struct {
	WeatherService AppWeather.WeatherService
	WeatherFactory AppWeather.WeatherFactory
}
type WeatherControllerInterface interface {
	GetWeatherInfo(c *fiber.Ctx) string
}

func NewWeatherController(weatherService AppWeather.WeatherService, WeatherFactory AppWeather.WeatherFactory) *WeatherController {
	return &WeatherController{
		WeatherService: weatherService,
		WeatherFactory: WeatherFactory,
	}
}

func (wc WeatherController) GetWeatherInfo(c *fiber.Ctx) string {
	command := wc.WeatherFactory.BuildWeatherCommand(c)

	return wc.WeatherService.GetWeatherInfo(command)
}
