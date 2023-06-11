//go:generate wire
package Config

import (
	"github.com/google/wire"
)

type AppService struct{}

var (
	AppServiceSet = wire.NewSet(provideAppService)
)

func InitDIContainer() (*AppService, error) {
	wire.Build(AppServiceSet)

	return &AppService{}, nil
}

func provideAppService() *AppService {
	return &AppService{}
}
