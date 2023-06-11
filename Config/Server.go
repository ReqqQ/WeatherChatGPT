package Config

import (
	"WeatherAPI/Config/DI"
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	_, err := DI.InitDIContainer()
	if err != nil {
		return nil
	}

	return fiber.New()
}

func Start(app *fiber.App) {
	app.Listen(":3000")
}
