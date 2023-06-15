package InfrastructureChatGPTPackage

import (
	"encoding/json"

	AppChatGPT "github.com/ReqqQ/ChatGPT/App/ChatGPT"
)

type InfrastructureChatGPTPackageFactory struct {
}

func (factory InfrastructureChatGPTPackageFactory) GetChatGPTDTO(jsonData json.RawMessage) AppChatGPT.ChatGPTDTO {
	return AppChatGPT.ChatGPTDTO{Json: jsonData}
}
