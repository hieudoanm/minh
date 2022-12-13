package wards_service

import postgres "chatbot-functions/src/common/libs/postgres"

type Ward struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetWards() []Ward {
	database := postgres.GetDatabase()

	var wards []Ward
	result := database.Find(&wards)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return wards
}

func GetWard(id string) Ward {
	database := postgres.GetDatabase()

	var ward Ward
	database.First(&ward, "id = ?", id)

	return ward
}
