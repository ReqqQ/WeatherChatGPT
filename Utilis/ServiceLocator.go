package Utilis

import (
	"WeatherAPI/App/Weather"
	"WeatherAPI/DomainModel/Weather"
	"WeatherAPI/UI/API"
)

type AppService struct {
	Controller  *Controller
	Application *AppServices
	DomainModel *DomainModelServices
}

type Controller struct {
	WeatherController *UIAPI.WeatherController
}

type AppServices struct {
	WeatherApp *weatherServiceGroup
}

type DomainModelServices struct {
	WeatherDomain *weatherDomainServiceGroup
}

type weatherServiceGroup struct {
	WeatherService *AppWeather.WeatherService
	WeatherFactory *AppWeather.WeatherFactory
}

type weatherDomainServiceGroup struct {
	WeatherService *DomainModelWeather.WeatherDMService
}

func initControllers() *Controller {
	weather := &UIAPI.WeatherController{}

	return &Controller{WeatherController: weather}
}

func initAppServices() *AppServices {
	return &AppServices{
		WeatherApp: initWeatherServiceGroup(),
	}
}

func initDomainModelServices() *DomainModelServices {
	return &DomainModelServices{
		WeatherDomain: initWeatherDomainServiceGroup(),
	}
}

func initWeatherDomainServiceGroup() *weatherDomainServiceGroup {
	weatherService := &DomainModelWeather.WeatherDMService{}

	return &weatherDomainServiceGroup{
		WeatherService: weatherService,
	}
}

func initWeatherServiceGroup() *weatherServiceGroup {
	weatherService := &AppWeather.WeatherService{}
	weatherFactory := &AppWeather.WeatherFactory{}

	return &weatherServiceGroup{
		WeatherService: weatherService,
		WeatherFactory: weatherFactory,
	}
}

func Init() AppService {
	return AppService{
		Controller:  initControllers(),
		Application: initAppServices(),
		DomainModel: initDomainModelServices(),
	}
}
