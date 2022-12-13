package ministers_service

import postgres "chatbot-functions/src/common/libs/postgres"

type Minister struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetMinisters() []Minister {
	database := postgres.GetDatabase()

	var ministers []Minister
	result := database.Find(&ministers)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return ministers
}

func GetMinister(id string) Minister {
	database := postgres.GetDatabase()

	var minister Minister
	database.First(&minister, "id = ?", id)

	return minister
}
