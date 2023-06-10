package main

import (
	"WeatherAPI/App"
	"WeatherAPI/UI"
	"WeatherAPI/Utilis"
)

func main() {
	app := App.Init()
	DIContainer := Utilis.Init()
	App.SetDependency(DIContainer)
	UI.InitGetRoutes(app)
	App.Start(app)
}
