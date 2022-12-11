package districts_service

import postgres "chatbot-functions/src/libs/postgres"

type District struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetDistricts() []District {
	database := postgres.GetDatabase()

	var districts []District
	result := database.Find(&districts)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return districts
}

func GetDistrict(id string) District {
	database := postgres.GetDatabase()

	var district District
	database.First(&district, "id = ?", id)

	return district
}
