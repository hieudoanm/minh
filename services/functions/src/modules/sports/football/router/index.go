package football_router

import (
	football_controller "chatbot-functions/src/modules/sports/football/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router) {
	router.GET("/football/teams", football_controller.GetTeams)
	router.GET("/football/teams/:id", football_controller.GetTeam)
	router.GET("/football/matches", football_controller.GetMatches)
	router.GET("/football/matches/:id", football_controller.GetMatches)
}
