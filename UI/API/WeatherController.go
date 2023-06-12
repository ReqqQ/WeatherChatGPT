package API

import (
	"github.com/gofiber/fiber/v2"

	AppWeather "WeatherAPI/App/Weather"
)

type WeatherController struct {
	WeatherService AppWeather.WeatherService
	WeatherFactory AppWeather.WeatherFactory
}

func (wc WeatherController) GetWeatherInfo(c *fiber.Ctx) AppWeather.WeatherDTO {
	command := wc.WeatherFactory.BuildWeatherCommand(c)

	return wc.WeatherService.GetWeatherInfo(command)
}
