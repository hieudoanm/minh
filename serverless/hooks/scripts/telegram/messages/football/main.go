package main

import (
	"chatbot-webhook/clients/telegram"
	"chatbot-webhook/services/football"
	"chatbot-webhook/utils"
)

var TELEGRAM_CHAT_ID = utils.Getenv("TELEGRAM_CHAT_ID", "")

func main() {
	var message = football.GetMatchesMessage(64)
	telegram.SendMessage(TELEGRAM_CHAT_ID, message, "")
}
