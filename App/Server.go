package App

import (
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	return fiber.New()
}

func Start(app *fiber.App) {
	app.Listen(":3000")
}
