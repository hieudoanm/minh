package chess_router

import (
	chess_controller "chatbot-functions/src/modules/sports/chess/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router) {
	router.GET("/chess/teams", chess_controller.GetTeams)
	router.GET("/chess/teams/:id", chess_controller.GetTeam)
	router.GET("/chess/matches", chess_controller.GetMatches)
	router.GET("/chess/matches/:id", chess_controller.GetMatches)
}
