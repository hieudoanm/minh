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

	crypto_router "chatbot-functions/src/modules/currencies/crypto/router"
	forex_router "chatbot-functions/src/modules/currencies/forex/router"
	trends_router "chatbot-functions/src/modules/google/trends/router"
	youtube_router "chatbot-functions/src/modules/google/youtube/router"
	health_router "chatbot-functions/src/modules/health/router"
	proxy_router "chatbot-functions/src/modules/proxy/router"
	status_router "chatbot-functions/src/modules/status/router"
	vietnam_router "chatbot-functions/src/modules/vietnam/router"

	"chatbot-functions/src/common/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	var router *httprouter.Router = httprouter.New()
	// Router
	crypto_router.Router(router)  // Crypto
	forex_router.Router(router)   // Forex
	trends_router.Router(router)  // Google - Trends
	youtube_router.Router(router) // Google - YouTube
	health_router.Router(router)  // Health
	proxy_router.Router(router)   // Proxy
	status_router.Router(router)  // Status
	vietnam_router.Router(router) // Vietnam
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
