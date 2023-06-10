package Config

import (
	"WeatherAPI/Config/DI"
	"github.com/gofiber/fiber/v2"
)

var App *DI.AppService

func Init() *fiber.App {
	App, _ = DI.InitializeAppService()

	return fiber.New()
}

func Start(app *fiber.App) {
	app.Listen(":3000")
}
