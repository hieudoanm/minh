package youtube_router

import (
	"github.com/julienschmidt/httprouter"

	youtube_controller "chatbot-functions/src/modules/google/youtube/controller"
)

func Router(router *httprouter.Router) {
	router.GET("/google/youtube/trending", youtube_controller.GetTrending)
}
