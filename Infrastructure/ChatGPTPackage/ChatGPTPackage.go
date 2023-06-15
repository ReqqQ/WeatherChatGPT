package InfrastructureChatGPTPackage

import (
	"encoding/json"

	AppChatGPT "github.com/ReqqQ/ChatGPT/App/ChatGPT"
)

type InfrastructureChatGPTPackage struct {
	AppChatGPTService AppChatGPT.AppChatGPTService
	AppChatGPTFactory AppChatGPT.AppChatGPTFactory
	Factory           InfrastructureChatGPTPackageFactory
}

func (gptPackage InfrastructureChatGPTPackage) GetChatGPTWheather(jsonData json.RawMessage) AppChatGPT.MessageDTO {
	return gptPackage.AppChatGPTService.GetDataFromChatGPT(gptPackage.Factory.GetChatGPTDTO(jsonData))
}
