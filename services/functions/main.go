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

	crypto_router "chatbot-functions/src/modules/crypto/router"
	forex_router "chatbot-functions/src/modules/forex/router"
	google_router "chatbot-functions/src/modules/google/router"
	health_router "chatbot-functions/src/modules/health/router"
	proxy_router "chatbot-functions/src/modules/proxy/router"
	status_router "chatbot-functions/src/modules/status/router"
	vietnam_router "chatbot-functions/src/modules/vietnam/router"

	"chatbot-functions/src/common/utils"

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
	// Proxy
	proxy_router.ProxyRouter(router)
	// Status
	status_router.StatusRouter(router)
	// Vietnam
	vietnam_router.VietnamRouter(router)
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
