package chess_controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"chatbot-functions/src/common/utils"
	chess_service "chatbot-functions/src/modules/sports/chess/service"
)

func GetTeam(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var team = chess_service.GetTeam()
	utils.ResponseWriter(writer, team)
}

func GetTeams(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var teams = chess_service.GetTeams()
	utils.ResponseWriter(writer, teams)
}

func GetMatch(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var match = chess_service.GetMatch()
	utils.ResponseWriter(writer, match)
}

func GetMatches(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var matches = chess_service.GetMatches()
	utils.ResponseWriter(writer, matches)
}
