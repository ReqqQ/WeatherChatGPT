package UIAPI

import (
	"github.com/gofiber/fiber/v2"
)

type WeatherController struct {
}
type WeatherControllerInterface interface {
	GetWeatherInfo(c *fiber.Ctx) string
}

func (wc WeatherController) GetWeatherInfo(c *fiber.Ctx) string {
	//service := Application.WeatherApp
	//command := service.WeatherFactory.BuildWeatherCommand(c)
	//
	//return service.WeatherService.GetWeatherInfo(command)

	return ""
}
