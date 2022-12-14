package google_router

import (
	trends_controller "chatbot-functions/src/modules/google/trends/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router) {
	router.GET("/google/trends", trends_controller.GetGoogleTrends)
}
