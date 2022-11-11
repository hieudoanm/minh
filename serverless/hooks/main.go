package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"chatbot-webhook/clients/telegram"
	"chatbot-webhook/utils"

	"github.com/julienschmidt/httprouter"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func GetHealth(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse StatusResponse = StatusResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func ProcessMessage(text string) string {
	return text
}

type WebhookRequestBody struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

func ProcessWebhookRequestBody(webhookRequestBody WebhookRequestBody) {
	var chatId int = webhookRequestBody.Message.Chat.Id
	var text string = webhookRequestBody.Message.Text
	var lowerText string = strings.ToLower(text)
	var message string = ProcessMessage(lowerText)
	telegram.SendMessage(chatId, message, "markdown")
}

func GetWebhook(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// Process Request Body
	var webhookRequestBody WebhookRequestBody
	jsonError := json.NewDecoder(request.Body).Decode(&webhookRequestBody)
	if jsonError != nil {
		http.Error(writer, jsonError.Error(), http.StatusBadRequest)
		return
	}
	ProcessWebhookRequestBody(webhookRequestBody)
	// Response
	var healthResponse StatusResponse = StatusResponse{"OK"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func main() {
	router := httprouter.New()
	// Router
	router.GET("/health", GetHealth)
	router.GET("/webhook", GetWebhook)
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
