package DomainModelChatGPT

import (
	"encoding/json"

	"WeatherAPI/Infrastructure/Base"
	InfrastructureChatGPTAPI "WeatherAPI/Infrastructure/ChatGPTAPI"
)

type ChatGPTService struct {
	BaseRequest Base.BaseRequest
	ChatGPTApi  InfrastructureChatGPTAPI.ChatGPTAPIRepository
}

func (cg ChatGPTService) ParseCityWeatherJsonToHumanLang(json json.RawMessage) Message {
	cg.BaseRequest.Init()
	response := cg.ChatGPTApi.TranslateWeatherJsonToHumanLanguage(json, cg.BaseRequest)

	return getChatGPTMessageResponse(response)
}
