package App

import (
	"WeatherAPI/Utilis"
)

var DomainModelServices Utilis.DomainModelServices
var Application Utilis.AppServices
var Controller Utilis.Controller

func SetDependency(dependency Utilis.AppService) {
	DomainModelServices = *dependency.DomainModel
	Application = *dependency.Application
	Controller = *dependency.Controller
}
