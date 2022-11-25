package main

import (
	"fmt"
	"log"
	"net/http"

	crypto_controller "chatbot-functions/src/controller/crypto"
	forex_controller "chatbot-functions/src/controller/forex"
	health_controller "chatbot-functions/src/controller/health"
	status_controller "chatbot-functions/src/controller/status"
	vnindex_controller "chatbot-functions/src/controller/vnindex"

	"chatbot-functions/src/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Router
	router.GET("/crypto/coins", crypto_controller.GetCryptoCoins)
	router.GET("/crypto/coins/:id", crypto_controller.GetCryptoCoin)
	router.GET("/forex/rates", forex_controller.GetForexRates)
	router.GET("/health", health_controller.GetHealth)
	router.GET("/status", status_controller.GetStatuses)
	router.GET("/status/:service", status_controller.GetStatus)
	router.GET("/vnindex/companies", vnindex_controller.GetVnindexCompanies)
	router.GET("/vnindex/history/:symbol", vnindex_controller.GetVnindexHistory)
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
