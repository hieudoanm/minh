// MINH API
//
//	MINH API
//
// Version: 0.0.1
// Schemes: http, https
// Host: localhost:8080
// BasePath: /
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// License: MIT
// Contact: Hieu Doan<hieumdoan@example.com> http://hieudoanm.github.io
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"

	crypto_router "chatbot-functions/src/router/crypto"
	forex_router "chatbot-functions/src/router/forex"
	google_router "chatbot-functions/src/router/google"
	health_router "chatbot-functions/src/router/health"
	status_router "chatbot-functions/src/router/status"
	vietnam_router "chatbot-functions/src/router/vietnam"

	"chatbot-functions/src/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	var router *httprouter.Router = httprouter.New()
	// Crypto
	crypto_router.CryptoRouter(router)
	// Forex
	forex_router.ForexRouter(router)
	// Google
	google_router.GoogleRouter(router)
	// Health
	health_router.HealthRouter(router)
	// Status
	status_router.StatusRouter(router)
	// Vietnam
	vietnam_router.VietnamRouter(router)
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
