package UI

import (
	"WeatherAPI/Config"
	"github.com/gofiber/fiber/v2"
)

func InitGetRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Get("/weather", func(c *fiber.Ctx) error {
		return c.JSON(Config.App.Controller.WeatherController.GetWeatherInfo(c))
	})
}
