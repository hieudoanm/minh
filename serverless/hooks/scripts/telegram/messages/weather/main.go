package main

import (
	"chatbot-webhook/clients/telegram"
	"chatbot-webhook/services/weather"
	"chatbot-webhook/utils"
	"fmt"
)

// Environment Variables
var TELEGRAM_CHAT_ID = utils.Getenv("TELEGRAM_CHAT_ID", "")

func main() {
	var city string = "ho chi minh city"
	var airVisualMessage string = weather.GetAirVisualMessage()
	var weatherMessage string = weather.GetWeatherMessage(city)
	var message string = fmt.Sprintf("%s\n%s", airVisualMessage, weatherMessage)
	fmt.Println("Message", message)

	sendMessageResponse, sendMessageError := telegram.SendMessage(TELEGRAM_CHAT_ID, message, "")
	if sendMessageError != nil {
		fmt.Println("sendMessageError", sendMessageError)
		return
	}
	fmt.Println("sendMessageResponse", sendMessageResponse)

}
