package provinces_service

import postgres "chatbot-functions/src/common/libs/postgres"

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetProvinces() []Province {
	database := postgres.GetDatabase()

	var provinces []Province
	result := database.Find(&provinces)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return provinces
}

func GetProvince(id string) Province {
	database := postgres.GetDatabase()

	var province Province
	database.First(&province, "id = ?", id)

	return province
}
