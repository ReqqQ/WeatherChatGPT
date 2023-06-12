
# WeatherChatGPT

WeatherChatGPT is an innovative GitHub repository that combines the power of OpenWeather API and a ChatGPT model. It enables users to enter a postal code and country code in ISO format, and the endpoint responds with a human-language description of the weather conditions.

With WeatherChatGPT, users can easily retrieve weather information tailored to their specific location. By leveraging the OpenWeather API, the program fetches accurate and up-to-date weather data. The ChatGPT model adds a conversational element, allowing users to engage in natural language interactions while obtaining weather descriptions.



## Authors

- [@ReqqQ](https://www.github.com/ReqqQ)


## API Reference

#### Get info about weather

```http
  POST /api/v1/weather
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `zipCode` | `string` | **Required**. Zip code of country |
| `countryCode` | `string` | **Required**. ISO country |


## Demo
![alt text](https://i.ibb.co/RDJKW84/scr.png)

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`WEATHER_API_KEY`

`CHAT_GPT_KEY`


## Tech Stack

**Client API:** Fast HTTP

**Lang:** GoLang
