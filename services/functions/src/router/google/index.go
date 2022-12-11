package google_router

import (
	google_controller "chatbot-functions/src/controller/google"

	"github.com/julienschmidt/httprouter"
)

func GoogleRouter(router *httprouter.Router) {
	router.GET("/google/trends", google_controller.GetGoogleTrends)
}
