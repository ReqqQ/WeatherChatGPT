package DomainModelChatGPT

import (
	"encoding/json"

	"WeatherAPI/Infrastructure/Base"
	InfrastructureChatGPTPackage "WeatherAPI/Infrastructure/ChatGPTPackage"
)

type ChatGPTService struct {
	BaseRequest    Base.BaseRequest
	ChatGPTPackage InfrastructureChatGPTPackage.InfrastructureChatGPTPackage
	Factory        DomainModelChatGPTFactory
	Builder        DomainModelChatGPTBuilder
}

func (cg ChatGPTService) ParseCityWeatherJsonToHumanLang(json json.RawMessage) MessageVO {
	cg.BaseRequest.Init()

	messageDTO := cg.ChatGPTPackage.GetChatGPTWheather(cg.Builder.BuildWheatherPrompt(json))

	return cg.Factory.GetMessageVO(messageDTO.Content)
}
