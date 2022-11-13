package main

import (
	"chatbot-webhook/clients/telegram"
	"chatbot-webhook/utils"
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
