package DomainModelWeather

import (
	"WeatherAPI/Infrastructure/WeatherAPI"
	"encoding/json"
	"fmt"
)

type WeatherDMVO struct {
	ZipCode     string
	CountryCode string
}

func buildWeatherResponse(responseData []byte) WeatherAPI.WeatherCityResponse {
	var response WeatherAPI.WeatherCityResponse
	err := json.Unmarshal(responseData, &response)
	if err != nil {
		return WeatherAPI.WeatherCityResponse{}
	}

	return response
}

func buildRawWeatherCityJson(data []byte) json.RawMessage {
	var jsonData json.RawMessage
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Błąd podczas przetwarzania danych JSON:", err)
		return json.RawMessage("")
	}

	return jsonData
}

func buildWeatherRequest(vo WeatherDMVO) WeatherAPI.WeatherRequest {
	return WeatherAPI.WeatherRequest{
		ZipCode:     vo.ZipCode,
		CountryCode: vo.CountryCode,
	}
}
