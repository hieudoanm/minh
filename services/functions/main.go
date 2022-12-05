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

	crypto_controller "chatbot-functions/src/controller/crypto"
	forex_controller "chatbot-functions/src/controller/forex"
	google_controller "chatbot-functions/src/controller/google"
	health_controller "chatbot-functions/src/controller/health"
	status_controller "chatbot-functions/src/controller/status"
	ministers_controller "chatbot-functions/src/controller/vietnam/government/ministers"
	ministries_controller "chatbot-functions/src/controller/vietnam/government/ministries"
	presidents_controller "chatbot-functions/src/controller/vietnam/government/presidents"
	districts_controller "chatbot-functions/src/controller/vietnam/maps/districts"
	provinces_controller "chatbot-functions/src/controller/vietnam/maps/provinces"
	wards_controller "chatbot-functions/src/controller/vietnam/maps/wards"
	committees_controller "chatbot-functions/src/controller/vietnam/national-assembly/committees"
	members_controller "chatbot-functions/src/controller/vietnam/national-assembly/members"
	vnindex_controller "chatbot-functions/src/controller/vietnam/vnindex"

	"chatbot-functions/src/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Router
	router.GET("/crypto/coins", crypto_controller.GetCryptoCoins)
	router.GET("/crypto/coins/:id", crypto_controller.GetCryptoCoin)
	router.GET("/forex/rates", forex_controller.GetForexRates)
	router.GET("/google/trends", google_controller.GetGoogleTrends)
	router.GET("/health", health_controller.GetHealth)
	router.GET("/status", status_controller.GetStatuses)
	router.GET("/status/:service", status_controller.GetStatus)
	router.GET("/vnindex/companies", vnindex_controller.GetVnindexCompanies)
	router.GET("/vnindex/history/:symbol", vnindex_controller.GetVnindexHistory)
	// Government
	router.GET("/government/ministers", ministers_controller.GetMinisters)
	router.GET("/government/ministers/:id", ministers_controller.GetMinister)
	router.GET("/government/ministries", ministries_controller.GetMinistries)
	router.GET("/government/ministries/:id", ministries_controller.GetMinistry)
	router.GET("/government/presidents", presidents_controller.GetPresidents)
	router.GET("/government/presidents/:id", presidents_controller.GetPresident)
	// Maps
	router.GET("/maps/provinces", provinces_controller.GetProvinces)
	router.GET("/maps/provinces/:id", provinces_controller.GetProvince)
	router.GET("/maps/districts", districts_controller.GetDistricts)
	router.GET("/maps/districts/:id", districts_controller.GetDistricts)
	router.GET("/maps/wards", wards_controller.GetWards)
	router.GET("/maps/wards/:id", wards_controller.GetWard)
	// National Assembly
	router.GET("/national-assembly/committees", committees_controller.GetCommittees)
	router.GET("/national-assembly/committees/:id", committees_controller.GetCommittee)
	router.GET("/national-assembly/members", members_controller.GetMembers)
	router.GET("/national-assembly/members/:id", members_controller.GetMember)
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
