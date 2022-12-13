package vietnam_vnindex_router

import (
	"github.com/julienschmidt/httprouter"

	vnindex_controller "chatbot-functions/src/modules/vietnam/controller/vnindex"
)

func VietnamVnindexRouter(router *httprouter.Router) {
	router.GET("/vnindex/companies", vnindex_controller.GetVnindexCompanies)
	router.GET("/vnindex/history/:symbol", vnindex_controller.GetVnindexHistory)

}
