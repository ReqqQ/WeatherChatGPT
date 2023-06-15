package DomainModelChatGPT

import "encoding/json"

type DomainModelChatGPTBuilder struct {
}

func (builder DomainModelChatGPTBuilder) BuildWheatherPrompt(jsonRaw json.RawMessage) json.RawMessage {
	content := "Describe the weather in human language use json(without talk about source data),describe only weather,prepare long describe in PL and EN : " + string(jsonRaw)
	encoded, _ := json.Marshal(content)

	return encoded
}
