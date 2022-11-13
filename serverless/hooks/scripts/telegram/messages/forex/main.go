package main

import (
	"chatbot-webhook/clients/telegram"
	"chatbot-webhook/services/forex"
	"chatbot-webhook/utils"
)

var TELEGRAM_CHAT_ID = utils.Getenv("TELEGRAM_CHAT_ID", "")

func main() {
	var message = forex.GetForexMessage()
	telegram.SendMessage(TELEGRAM_CHAT_ID, message, "")
}
