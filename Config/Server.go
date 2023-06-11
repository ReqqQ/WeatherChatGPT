package Config

import (
	"WeatherAPI/Config/DI"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Init() *fiber.App {
	_, err := DI.InitDIContainer()
	if err != nil {
		return nil
	}
	err = godotenv.Load(".env")

	return fiber.New()
}

func Start(app *fiber.App) {
	app.Listen(":3000")
}
