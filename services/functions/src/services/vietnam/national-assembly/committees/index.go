package committees_service

import postgres "chatbot-functions/src/libs/postgres"

type Committee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetCommittees() []Committee {
	database := postgres.GetDatabase()

	var committees []Committee
	result := database.Find(&committees)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return committees
}

func GetCommittee(id string) Committee {
	database := postgres.GetDatabase()

	var committee Committee
	database.First(&committee, "id = ?", id)

	return committee
}
