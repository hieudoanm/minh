package telegram

import (
	"bytes"
	"chatbot-webhook/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const TELEGRAM_BOT = "https://api.telegram.org/bot"

var TELEGRAM_TOKEN = utils.Getenv("TELEGRAM_TOKEN", "")

func Post(url string, requestBody map[string]string) ([]byte, error) {
	jsonData, jsonMarshalError := json.Marshal(requestBody)
	if jsonMarshalError != nil {
		return nil, jsonMarshalError
	}

	response, httpPostError := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if httpPostError != nil {
		return nil, httpPostError
	}
	defer response.Body.Close()

	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		return nil, readBodyError
	}

	return body, nil
}

func SendMessage(chatId string, message string, parseMode string) ([]byte, error) {
	var requestBody map[string]string = map[string]string{}
	requestBody["chat_id"] = chatId
	requestBody["text"] = message
	if parseMode == "" {
		requestBody["parse_mode"] = "markdown"
	} else {
		requestBody["parse_mode"] = parseMode
	}

	var url = fmt.Sprintf(
		"%s%s/sendMessage",
		TELEGRAM_BOT,
		TELEGRAM_TOKEN,
	)

	body, postError := Post(url, requestBody)
	if postError != nil {
		return nil, postError
	}

	return body, nil
}

func SetWebhook(webhook string) ([]byte, error) {
	var requestBody map[string]string = map[string]string{}
	requestBody["url"] = webhook

	var url = fmt.Sprintf(
		"%s%s/setWebhook",
		TELEGRAM_BOT,
		TELEGRAM_TOKEN,
	)

	body, postError := Post(url, requestBody)
	if postError != nil {
		return nil, postError
	}

	return body, nil
}

func DeleteWebhook(webhook string) ([]byte, error) {
	var requestBody map[string]string = map[string]string{}
	requestBody["url"] = webhook

	var url = fmt.Sprintf(
		"%s%s/deleteWebhook",
		TELEGRAM_BOT,
		TELEGRAM_TOKEN,
	)

	body, postError := Post(url, requestBody)
	if postError != nil {
		return nil, postError
	}

	return body, nil
}

type WebhookInfo struct {
	Ok     bool `json:"ok"`
	Result struct {
		Url                  string `json:"url"`
		HasCustomCertificate bool   `json:"has_custom_certificate"`
		PendingUpdateCount   int    `json:"pending_update_count"`
	} `json:"result"`
}

func GetWebhookInfo() (WebhookInfo, error) {
	var url = fmt.Sprintf(
		"%s%s/getWebhookInfo",
		TELEGRAM_BOT,
		TELEGRAM_TOKEN,
	)

	body, postError := Post(url, map[string]string{})
	if postError != nil {
		return WebhookInfo{}, postError
	}

	var webhookInfo WebhookInfo
	jsonUnmarshalError := json.Unmarshal(body, &webhookInfo)
	if jsonUnmarshalError != nil {
		return WebhookInfo{}, jsonUnmarshalError
	}

	return webhookInfo, nil
}
