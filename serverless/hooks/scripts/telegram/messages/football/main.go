package main

import (
	"chatbot-webhook/src/clients/telegram"
	"chatbot-webhook/src/services/football"
	"chatbot-webhook/src/utils"
)

var TELEGRAM_CHAT_ID = utils.Getenv("TELEGRAM_CHAT_ID", "")

func main() {
	var message = football.GetMatchesMessage(64)
	telegram.SendMessage(TELEGRAM_CHAT_ID, message, "")
}
