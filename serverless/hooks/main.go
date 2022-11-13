package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"chatbot-webhook/clients/telegram"
	"chatbot-webhook/services/football"
	"chatbot-webhook/services/news"
	"chatbot-webhook/services/weather"
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

var INTENTS []string = []string{
	"blockchain news",
	"forex",
	"liverpool",
	"news",
	"uuid",
	"weather",
}

func ProcessMessage(text string) string {
	if utils.Contains(INTENTS, text) {
		if text == "weather" {
			var airVisualMessage = weather.GetAirVisualMessage()
			var weatherMessage = weather.GetWeatherMessage("ho chi minh city")
			return fmt.Sprintf("%s\n%s", airVisualMessage, weatherMessage)
		} else if text == "football" {
			return football.GetMatchesMessage(64)
		} else if text == "news" {
			return news.GetTopHeadlinesMessage()
		} else if text == "blockchain news" {
			return news.GetBlockchainCryptoMessage()
		} else {
			return "N/A"
		}
	} else if text == "help" {
		return strings.Join(INTENTS, "\n")
	}
	return "N/A"
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
	telegram.SendMessage(strconv.Itoa(chatId), message, "markdown")
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
