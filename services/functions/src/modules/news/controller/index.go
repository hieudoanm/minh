package news_controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetSources(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {}

func GetTopHeadlines(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {}
