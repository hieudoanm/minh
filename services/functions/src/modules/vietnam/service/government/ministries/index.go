package ministries_service

import postgres "chatbot-functions/src/common/libs/postgres"

type Ministry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetMinistries() []Ministry {
	database := postgres.GetDatabase()

	var ministries []Ministry
	result := database.Find(&ministries)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return ministries
}

func GetMinistry(id string) Ministry {
	database := postgres.GetDatabase()

	var ministry Ministry
	database.First(&ministry, "id = ?", id)

	return ministry
}
