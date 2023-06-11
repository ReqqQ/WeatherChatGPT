package UI

import (
	"WeatherAPI/UI/API"
	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
	WeatherController API.WeatherController
}

func (co Controllers) InitGetRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Get("/weather", func(c *fiber.Ctx) error {
		return c.JSON(co.WeatherController.GetWeatherInfo(c))
	})
}
