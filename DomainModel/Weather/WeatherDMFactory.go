package DomainModelWeather

import (
	"encoding/json"
	"fmt"

	"WeatherAPI/Infrastructure/WeatherAPI"
)

type WeatherDMVO struct {
	ZipCode     string
	CountryCode string
}

func buildWeatherResponse(responseData []byte) InfrastructureWeatherAPI.WeatherCityResponse {
	var response InfrastructureWeatherAPI.WeatherCityResponse
	err := json.Unmarshal(responseData, &response)
	if err != nil {
		return InfrastructureWeatherAPI.WeatherCityResponse{}
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

func buildWeatherRequest(vo WeatherDMVO) InfrastructureWeatherAPI.WeatherRequest {
	return InfrastructureWeatherAPI.WeatherRequest{
		ZipCode:     vo.ZipCode,
		CountryCode: vo.CountryCode,
	}
}
