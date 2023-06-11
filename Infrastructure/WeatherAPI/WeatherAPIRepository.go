package WeatherAPI

import (
	"github.com/valyala/fasthttp"
	"time"
)

const API_WEATHER_ZIP = "http://api.openweathermap.org/geo/1.0/zip?"
const API_CITY_WEATHER = "https://api.openweathermap.org/data/2.5/weather?"

type WeatherAPIRepository struct {
}

var client *fasthttp.Client

func (api WeatherAPIRepository) Init() {
	client = &fasthttp.Client{
		MaxConnsPerHost:     1000,
		ReadBufferSize:      4096,
		WriteBufferSize:     4096,
		ReadTimeout:         5 * time.Second,
		WriteTimeout:        5 * time.Second,
		MaxIdleConnDuration: 10 * time.Second,
	}
}

func (api WeatherAPIRepository) GetCityCoordinate(vo WeatherRequest) []byte {
	request, response := baseSettings()
	request.SetRequestURI(API_WEATHER_ZIP + buildGetWeatherStats(vo))
	callAPI(request, response)

	return response.Body()
}

func (api WeatherAPIRepository) GetCityWeather(cityResponse WeatherCityResponse) []byte {
	request, response := baseSettings()
	request.SetRequestURI(API_CITY_WEATHER + buildGetCityWeather(cityResponse))
	callAPI(request, response)

	return response.Body()
}
