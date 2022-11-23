package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type CoinsResponseBody struct {
	Status string `json:"status"`
	Data   struct {
		Stats struct {
			Total          int    `json:"total"`
			TotalCoins     int    `json:"totalCoins"`
			TotalMarkets   int    `json:"totalMarkets"`
			TotalExchanges int    `json:"totalExchanges"`
			TotalMarketCap string `json:"totalMarketCap"`
			Total24hVolume string `json:"Total24hVolume"`
		} `json:"stats"`
		Coins []struct {
			Uuid           string   `json:"uuid"`
			Symbol         string   `json:"symbol"`
			Name           string   `json:"name"`
			Color          string   `json:"color"`
			IconUrl        string   `json:"iconUrl"`
			MarketCap      string   `json:"marketCap"`
			Price          string   `json:"price"`
			ListedAt       int      `json:"listedAt"`
			Tier           int      `json:"tier"`
			Change         string   `json:"change"`
			Rank           int      `json:"rank"`
			Sparkline      []string `json:"sparkline"`
			LowVolume      bool     `json:"lowVolume"`
			CoinrankingUrl string   `json:"coinrankingUrl"`
			BtcPrice       string   `json:"btcPrice"`
			Volume24h      string   `json:"24hVolume"`
		} `json:"coins"`
	} `json:"data"`
}

var OrderBy map[string]string = map[string]string{
	"VOLUME_24H": "24hVolume",
	"CHANGE":     "change",
	"LISTED_AT":  "listedAt",
	"MARKET_CAP": "marketCap",
	"PRICE":      "price",
}

var OrderDirection map[string]string = map[string]string{
	"ASC":  "asc",
	"DESC": "desc",
}

var TimePeriod map[string]string = map[string]string{
	"1h":  "1h",
	"3h":  "3h",
	"12h": "12h",
	"24h": "24h",
	"7d":  "7d",
	"30d": "30d",
	"3m":  "3m",
	"1y":  "1y",
	"3y":  "3y",
	"5y":  "5y",
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

func GetCryptoCoins(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// Query Parameters
	limit := GetQueryParameter(request, "limit", "100")
	offset := GetQueryParameter(request, "offset", "0")
	orderBy := GetQueryParameter(request, "orderBy", OrderBy["MARKET_CAP"])
	orderDirection := GetQueryParameter(request, "orderDirection", OrderDirection["DESC"])
	timePeriod := GetQueryParameter(request, "timePeriod", TimePeriod["24h"])
	tier := GetQueryParameter(request, "tier", "")
	tag := GetQueryParameter(request, "tag", "")
	var queryParameters = []string{}
	if limit != "" {
		queryParameters = append(queryParameters, "limit="+limit)
	}
	if offset != "" {
		queryParameters = append(queryParameters, "offset="+offset)
	}
	if orderBy != "" {
		queryParameters = append(queryParameters, "orderBy="+orderBy)
	}
	if orderDirection != "" {
		queryParameters = append(queryParameters, "orderDirection="+orderDirection)
	}
	if timePeriod != "" {
		queryParameters = append(queryParameters, "timePeriod="+timePeriod)
	}
	if tier != "" {
		queryParameters = append(queryParameters, "tier="+tier)
	}
	if tag != "" {
		queryParameters = append(queryParameters, "tags="+tag)
	}
	// HTTP Request
	var url string = fmt.Sprintf(
		"https://api.coinranking.com/v2/coins?%s",
		strings.Join(queryParameters, "&"),
	)
	log.Printf("URL: %s", url)
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetCryptoCoins")
	}
	defer response.Body.Close()
	// Get Body
	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		log.Println("Fail to GetCryptoCoins")
	}
	// Parse JSON
	var coinsResponseBody CoinsResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &coinsResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetCryptoCoins")
	}

	json.NewEncoder(writer).Encode(coinsResponseBody)
}

func Handler() {
	http.HandleFunc("/", GetCryptoCoins)
	log.Println("🚀 Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
