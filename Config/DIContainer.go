//go:generate wire
package Config

import (
	"WeatherAPI/App/Weather"
	"WeatherAPI/DomainModel/Weather"
	"WeatherAPI/UI/API"

	"github.com/google/wire"
)

type AppService struct {
	Controller  *Controller
	Application *AppServices
	DomainModel *DomainModelServices
}

type Controller struct {
	WeatherController *API.WeatherController
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
	weather := &API.WeatherController{}

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

var (
	AppServiceSet = wire.NewSet(initControllers, initAppServices, initDomainModelServices, provideAppService, provideWeatherController)
)

func InitializeAppService() (*AppService, error) {
	wire.Build(AppServiceSet)
	return &AppService{}, nil
}

func provideWeatherController(weatherService AppWeather.WeatherService, WeatherFactory AppWeather.WeatherFactory) *API.WeatherController {
	return API.NewWeatherController(weatherService, WeatherFactory)
}

func provideAppService(controller *Controller, appServices *AppServices, domainModel *DomainModelServices) *AppService {
	return &AppService{
		Controller:  controller,
		Application: appServices,
		DomainModel: domainModel,
	}
}
