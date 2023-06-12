package DomainModelChatGPT

import (
	"encoding/json"
	"log"
)

func (s ChatGPTMessageResponse) getFirstChoice() Message {
	return s.Choices[0].Message
}

type ChatGPTMessageResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (m Message) GetContent() string {
	return m.Content
}

func getChatGPTMessageResponse(response []byte) Message {
	var chatGPTMessageResponse ChatGPTMessageResponse
	err := json.Unmarshal(response, &chatGPTMessageResponse)
	if err != nil {
		log.Fatal(err)
	}

	return chatGPTMessageResponse.getFirstChoice()
}
