package presidents_service

import postgres "chatbot-functions/src/libs/postgres"

type President struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetPresidents() []President {
	database := postgres.GetDatabase()

	var presidents []President
	result := database.Find(&presidents)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return presidents
}

func GetPresident(id string) President {
	database := postgres.GetDatabase()

	var president President
	database.First(&president, "id = ?", id)

	return president
}
