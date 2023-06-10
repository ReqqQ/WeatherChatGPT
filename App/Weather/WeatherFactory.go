package AppWeather

import (
	DomainModelWeather "WeatherAPI/DomainModel/Weather"
	"github.com/gofiber/fiber/v2"
)

type WeatherFactory struct{}
type WeatherFactoryInterface interface {
	BuildWeatherCommand(c *fiber.Ctx)
}

type BuildWeatherCommand struct {
	ZipCode     string `json:"zipCode" form:"zipCode"`
	CountryCode string `json:"countryCode" form:"countryCode"`
}

func (receiver BuildWeatherCommand) GetZipCode() string {
	return receiver.ZipCode
}

func (receiver BuildWeatherCommand) GetCountryCode() string {
	return receiver.CountryCode
}

func (wf WeatherFactory) BuildWeatherCommand(c *fiber.Ctx) BuildWeatherCommand {
	command := new(BuildWeatherCommand)

	if err := c.BodyParser(command); err != nil {
		panic(err)
	}

	return *command
}

func (wf WeatherFactory) BuildWeatherDMVO(zipCode string, countryCode string) DomainModelWeather.WeatherDMVO {
	return DomainModelWeather.WeatherDMVO{
		ZipCode:     zipCode,
		CountryCode: countryCode,
	}
}
