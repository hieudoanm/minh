package vnindex_controller

import (
	"chatbot-functions/src/common/constants"
	vnindex_client "chatbot-functions/src/modules/vietnam/client/vnindex"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetVnindexCompanies(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var companiesResponseBody []vnindex_client.CompaniesResponseBody = vnindex_client.GetVnindexCompanies()
	json.NewEncoder(writer).Encode(companiesResponseBody)
}

func GetVnindexHistory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var symbol string = params.ByName("symbol")
	var historyResponseBody []vnindex_client.HistoryResponseBody = vnindex_client.GetVnindexHistory(symbol)
	json.NewEncoder(writer).Encode(historyResponseBody)
}
