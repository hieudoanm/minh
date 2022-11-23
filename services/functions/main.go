package main

import (
	"fmt"
	"log"
	"net/http"

	"chatbot-functions/src/clients/crypto"
	"chatbot-functions/src/clients/forex"
	"chatbot-functions/src/clients/health"
	"chatbot-functions/src/clients/status"
	"chatbot-functions/src/clients/vnindex"
	"chatbot-functions/src/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Router
	router.GET("/crypto/coins", crypto.GetCryptoCoins)
	router.GET("/crypto/coins/:id", crypto.GetCryptoCoin)
	router.GET("/forex/rates", forex.GetForexRates)
	router.GET("/health", health.GetHealth)
	router.GET("/status", status.GetStatuses)
	router.GET("/status/:service", status.GetStatus)
	router.GET("/vnindex/companies", vnindex.GetVnindexCompanies)
	router.GET("/vnindex/history/:symbol", vnindex.GetVnindexHistory)
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
