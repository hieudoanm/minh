package main

import (
	"chatbot-webhook/src/clients/telegram"
	"chatbot-webhook/src/services/forex"
	"chatbot-webhook/src/utils"
)

var TELEGRAM_CHAT_ID = utils.Getenv("TELEGRAM_CHAT_ID", "")

func main() {
	var message = forex.GetForexMessage()
	telegram.SendMessage(TELEGRAM_CHAT_ID, message, "")
}
