package news_router

import (
	news_controller "chatbot-functions/src/modules/news/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router) {
	router.GET("/news/sources", news_controller.GetSources)
	router.GET("/news/top-headlines", news_controller.GetTopHeadlines)

}
