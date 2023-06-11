package API

import (
	AppWeather "WeatherAPI/App/Weather"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type WeatherController struct {
	WeatherService AppWeather.WeatherService
	WeatherFactory AppWeather.WeatherFactory
}

func (wc WeatherController) GetWeatherInfo(c *fiber.Ctx) json.RawMessage {
	command := wc.WeatherFactory.BuildWeatherCommand(c)

	return wc.WeatherService.GetWeatherInfo(command)
}
