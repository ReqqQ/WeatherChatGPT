package WeatherAPI

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
	"os"
)

type WeatherRequest struct {
	ZipCode     string
	CountryCode string
}

type WeatherCityResponse struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

func buildGetWeatherStats(vo WeatherRequest) string {
	params := url.Values{}
	params.Set("zip", vo.ZipCode+","+vo.CountryCode)
	params.Set("appid", os.Getenv("WEATHER_API_KEY"))

	return params.Encode()
}
func buildGetCityWeather(cityResponse WeatherCityResponse) string {
	params := url.Values{}
	params.Set("units", "metric")
	params.Set("lat", fmt.Sprintf("%.2f", cityResponse.Lat))
	params.Set("lon", fmt.Sprintf("%.2f", cityResponse.Lon))
	params.Set("appid", os.Getenv("WEATHER_API_KEY"))

	return params.Encode()
}

func baseSettings() (*fasthttp.Request, *fasthttp.Response) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	return request, response
}

func callAPI(request *fasthttp.Request, response *fasthttp.Response) *fasthttp.Response {
	if err := client.Do(request, response); err != nil {
		fmt.Printf("Błąd podczas wysyłania żądania: %s\n", err)
	}

	return response
}
