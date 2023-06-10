package main

import (
	"WeatherAPI/Config"
	"WeatherAPI/UI"
)

func main() {
	app := Config.Init()
	UI.InitGetRoutes(app)
	Config.Start(app)
}
