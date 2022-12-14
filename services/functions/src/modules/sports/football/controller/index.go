package football_controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"chatbot-functions/src/common/utils"
	football_service "chatbot-functions/src/modules/sports/football/service"
)

func GetTeam(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var team = football_service.GetTeam()
	utils.ResponseWriter(writer, team)
}

func GetTeams(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var teams = football_service.GetTeams()
	utils.ResponseWriter(writer, teams)
}

func GetMatch(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var match = football_service.GetMatch()
	utils.ResponseWriter(writer, match)
}

func GetMatches(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var matches = football_service.GetMatches()
	utils.ResponseWriter(writer, matches)
}
