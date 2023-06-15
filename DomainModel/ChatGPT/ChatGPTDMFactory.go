package DomainModelChatGPT

type DomainModelChatGPTFactory struct {
}

type MessageVO struct {
	Content string `json:"content"`
}

func (factory DomainModelChatGPTFactory) GetMessageVO(responseContent string) MessageVO {
	return MessageVO{Content: responseContent}
}

func (m MessageVO) GetContent() string {
	return m.Content
}
