package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CoinResponseBody struct {
	Status string `json:"status"`
	Data   struct {
		Coin struct {
			Uuid        string `json:"uuid"`
			Symbol      string `json:"symbol"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Color       string `json:"color"`
			IconUrl     string `json:"iconUrl"`
			WebsiteUrl  string `json:"websiteUrl"`
			Links       []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				Url  string `json:"url"`
			} `json:"links"`
			Supply struct {
				Confirmed   bool   `json:"confirmed"`
				SupplyAt    int    `json:"supplyAt"`
				Max         string `json:"max"`
				Total       string `json:"total"`
				Circulating string `json:"circulating"`
			} `json:"supply"`
			NumberOfMarkets       int      `json:"numberOfMarkets"`
			NumberOfExchanges     int      `json:"numberOfExchanges"`
			Volume24h             string   `json:"24hVolume"`
			MarketCap             string   `json:"marketCap"`
			FullyDilutedMarketCap string   `json:"fullyDilutedMarketCap"`
			Price                 string   `json:"price"`
			BtcPrice              string   `json:"btcPrice"`
			PriceAt               int      `json:"priceAt"`
			Change                string   `json:"change"`
			Rank                  int      `json:"rank"`
			Sparkline             []string `json:"sparkline"`
			AllTimeHigh           struct {
				Price     string `json:"price"`
				Timestamp int    `json:"timestamp"`
			} `json:"allTimeHigh"`
			CoinrankingUrl string   `json:"coinrankingUrl"`
			Tier           int      `json:"tier"`
			LowVolume      bool     `json:"lowVolume"`
			ListedAt       int      `json:"listedAt"`
			Tags           []string `json:"tags"`
		} `json:"coin"`
	} `json:"data"`
}

func GetQueryParameter(request *http.Request, key string, defaultValue string) string {
	var value string = request.URL.Query().Get(key)
	if defaultValue == "" {
		return value
	}
	if defaultValue != "" && value == "" {
		return defaultValue
	}
	return value
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetCryptoCoin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// Get Query Parameter
	var id string = GetQueryParameter(request, "id", "")
	if id == "" {
		var errorResponse ErrorResponse = ErrorResponse{Message: "Missing Id"}
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}
	// HTTP Request
	var url string = fmt.Sprintf(
		"https://api.coinranking.com/v2/coin/%s",
		id,
	)
	log.Printf("URL: %s", url)
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetCryptoCoin")
	}
	defer response.Body.Close()
	// Get Body
	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		log.Println("Fail to GetCryptoCoin")
	}
	// Parse JSON
	var coinResponseBody CoinResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &coinResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetCryptoCoin")
	}

	json.NewEncoder(writer).Encode(coinResponseBody)
}

func Handler() {
	http.HandleFunc("/", GetCryptoCoin)
	log.Println("🚀 Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
