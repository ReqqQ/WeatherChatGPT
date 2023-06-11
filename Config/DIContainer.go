//go:generate wire
package Config

import (
	"WeatherAPI/UI"
	"WeatherAPI/UI/API"

	"github.com/google/wire"
)

type AppService struct{}

var (
	controllersSet = wire.NewSet(provideControllers)
	AppServiceSet  = wire.NewSet(provideAppService, controllersSet)
)

func InitDIContainer() (*AppService, error) {
	wire.Build(AppServiceSet)

	return &AppService{}, nil
}

func provideControllers(wc API.WeatherController) *UI.Controllers {
	return &UI.Controllers{
		WeatherController: wc,
	}
}

func provideAppService() *AppService {
	return &AppService{}
}
