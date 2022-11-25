package google_client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetTrends() map[string][]string {
	// HTTP Request
	var url = "https://trends.google.com/trends/hottrends/visualize/internal/data"
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetForexRates", httpGetError)
	}
	defer response.Body.Close()
	// Get Body
	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		log.Println("Fail to GetForexRates", readBodyError)
	}
	// Parse JSON
	var trendsResponseBody map[string][]string
	jsonUnmarshalError := json.Unmarshal(body, &trendsResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetForexRates", jsonUnmarshalError)
	}

	return trendsResponseBody
}
