package Base

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type BaseRequest struct{}

var client *fasthttp.Client

func (req BaseRequest) Init() {
	if client == nil {
		client = &fasthttp.Client{}
	}
}

func (req BaseRequest) BaseSettings() (*fasthttp.Request, *fasthttp.Response) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	return request, response
}

func (req BaseRequest) CallAPI(request *fasthttp.Request, response *fasthttp.Response) *fasthttp.Response {
	if err := client.Do(request, response); err != nil {
		fmt.Printf("Błąd podczas wysyłania żądania: %s\n", err)
	}

	return response
}

func (req BaseRequest) GetClient() *fasthttp.Client {
	return client
}
