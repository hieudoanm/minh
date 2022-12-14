package main

import (
	"chatbot-webhook/src/clients/telegram"
	"chatbot-webhook/src/services/news"
	"chatbot-webhook/src/utils"
)

// Environment Variables
var TELEGRAM_CHAT_ID = utils.Getenv("TELEGRAM_CHAT_ID", "")

func main() {
	var topHeadlinesMessage = news.GetTopHeadlinesMessage()
	telegram.SendMessage(TELEGRAM_CHAT_ID, topHeadlinesMessage, "")
	var blockchainCryptoMessage = news.GetBlockchainCryptoMessage()
	telegram.SendMessage(TELEGRAM_CHAT_ID, blockchainCryptoMessage, "")
}
