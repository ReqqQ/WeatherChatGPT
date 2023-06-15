package AppWeather

import (
	"github.com/gofiber/fiber/v2"

	DomainModelChatGPT "WeatherAPI/DomainModel/ChatGPT"
	DomainModelWeather "WeatherAPI/DomainModel/Weather"
)

type WeatherFactory struct{}

type BuildWeatherCommand struct {
	ZipCode     string `json:"zipCode" form:"zipCode"`
	CountryCode string `json:"countryCode" form:"countryCode"`
}

type WeatherDTO struct {
	Content string `json:"content"`
}

func (wf WeatherFactory) BuildWeatherDTO(message DomainModelChatGPT.MessageVO) WeatherDTO {
	return WeatherDTO{Content: message.GetContent()}
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
