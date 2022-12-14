package main

import (
	"chatbot-webhook/src/clients/telegram"
	"chatbot-webhook/src/utils"
	"fmt"
)

var TELEGRAM_WEBHOOK = utils.Getenv("TELEGRAM_WEBHOOK", "")

func main() {
	telegram.DeleteWebhook(TELEGRAM_WEBHOOK)
	telegram.SetWebhook(TELEGRAM_WEBHOOK)
	webhookInfo, webhookInfoError := telegram.GetWebhookInfo()
	if webhookInfoError != nil {
		fmt.Println("webhookInfoError", webhookInfoError)
	}
	fmt.Println("webhookInfo", webhookInfo)
}
