package InfrastructureChatGPTAPI

import (
	"encoding/json"
	"os"

	"github.com/valyala/fasthttp"
)

func buildCompletionsRequestData(request *fasthttp.Request) {
	request.Header.Set("Authorization", "Bearer "+os.Getenv("CHAT_GPT_KEY"))
	request.Header.Set("Content-Type", "application/json")

	request.Header.SetMethod(REQUEST_TYPE_POST)
	request.SetRequestURI(CHATGPT_COMPLETIONS_ENDPOINT)
}

func transformRequestBodyToJson(jsonMessage json.RawMessage) []byte {
	jsonData, err := json.Marshal(getRequestBody(jsonMessage))
	if err != nil {
	}

	return jsonData
}
