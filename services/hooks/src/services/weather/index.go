package weather

import (
	airvisual "chatbot-webhook/src/clients/air-visual"
	"chatbot-webhook/src/clients/weather"
	"chatbot-webhook/src/utils"
	"fmt"
	"net/url"
)

var API_KEY_AIR_VISUAL = utils.Getenv("API_KEY_AIR_VISUAL", "")
var API_KEY_OPEN_WEATHER_MAP = utils.Getenv("API_KEY_OPEN_WEATHER_MAP", "")

func GetAirVisualMessage() string {
	var country string = "Vietnam"
	var state string = "Ho Chi Minh City"
	var city string = "Ho Chi Minh City"
	airQuality, airQualityError := airvisual.GetAirQuality(API_KEY_AIR_VISUAL, country, state, city)
	if airQualityError != nil {
		fmt.Println("airQualityError", airQualityError)
		return "Air Quality: N/A"
	}
	return fmt.Sprintf("Air Quality: %d", airQuality.Data.Current.Pollution.Aqius)
}

func GetWeatherMessage(city string) string {
	if API_KEY_OPEN_WEATHER_MAP == "" {
		return "Weather: N/A"
	}

	var units string = "metric"
	weather, weatherError := weather.GetWeather(
		API_KEY_OPEN_WEATHER_MAP,
		url.QueryEscape(city),
		units,
	)
	if weatherError != nil {
		return "Weather: N/A"
	}
	fmt.Println(weather.Weather)
	var name string = weather.Name
	var main string = weather.Weather[0].Main
	var description string = weather.Weather[0].Description
	var temperature float64 = weather.Main.Temperature
	var feelsLike float64 = weather.Main.FeelsLike
	return fmt.Sprintf("%s\n%s (%s)\n%f°C - %f°C", name, main, description, temperature, feelsLike)
}
