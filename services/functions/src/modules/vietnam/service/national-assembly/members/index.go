package members_service

import postgres "chatbot-functions/src/common/libs/postgres"

type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetMembers() []Member {
	database := postgres.GetDatabase()

	var members []Member
	result := database.Find(&members)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return members
}

func GetMember(id string) Member {
	database := postgres.GetDatabase()

	var member Member
	database.First(&member, "id = ?", id)

	return member
}
