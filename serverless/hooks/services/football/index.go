package football

import (
	"chatbot-webhook/clients/football"
	"chatbot-webhook/utils"
	"fmt"
)

var API_KEY_FOOTBALL_DATA = utils.Getenv("API_KEY_FOOTBALL_DATA", "")

func GetLastMatchMessage(matches []football.Match) string {
	var finishedMatches = []football.Match{}
	for _, match := range matches {
		if match.Status == "FINISHED" {
			finishedMatches = append(finishedMatches, match)
		}
	}
	var lastMatch = finishedMatches[len(finishedMatches)-1]
	var utcDate = lastMatch.UtcDate
	var homeScore = lastMatch.Score.FullTime.Home
	var awayScore = lastMatch.Score.FullTime.Away
	var homeTeam = lastMatch.HomeTeam.Name
	var awayTeam = lastMatch.AwayTeam.Name
	return fmt.Sprintf("%s\n%s %d - %d %s", utcDate, homeTeam, homeScore, awayScore, awayTeam)
}

func GetNextMatchMessage(matches []football.Match) string {
	var scheduledMatches = []football.Match{}
	for _, match := range matches {
		if match.Status == "TIMED" {
			scheduledMatches = append(scheduledMatches, match)
		}
	}
	var nextMatch = scheduledMatches[0]
	var utcDate = nextMatch.UtcDate
	var homeScore = nextMatch.Score.FullTime.Home
	var awayScore = nextMatch.Score.FullTime.Away
	var homeTeam = nextMatch.HomeTeam.Name
	var awayTeam = nextMatch.AwayTeam.Name
	return fmt.Sprintf("%s\n%s %d - %d %s", utcDate, homeTeam, homeScore, awayScore, awayTeam)
}

func GetMatchesMessage(teamId int) string {
	matchesResponse, matchesError := football.GetMatchesByTeam(API_KEY_FOOTBALL_DATA, teamId)
	if matchesError != nil {
		return "Football: N/A"
	}
	var lastMatchMessage = GetLastMatchMessage(matchesResponse.Matches)
	var nextMatchMessage = GetNextMatchMessage(matchesResponse.Matches)
	var border = "------------------------------"
	return fmt.Sprintf("%s\n%s\n%s", lastMatchMessage, border, nextMatchMessage)
}
