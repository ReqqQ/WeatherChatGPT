package main

import (
	"WeatherAPI/Config"
	"WeatherAPI/UI"
)

func main() {
	app := Config.Init()
	controller := new(UI.Controllers)
	controller.InitGetRoutes(app)
	Config.Start(app)
}
