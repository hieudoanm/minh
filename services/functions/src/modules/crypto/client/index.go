package crypto_client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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

var ORDER_BY map[string]string = map[string]string{
	"VOLUME_24H": "24hVolume",
	"CHANGE":     "change",
	"LISTED_AT":  "listedAt",
	"MARKET_CAP": "marketCap",
	"PRICE":      "price",
}

var ORDER_DIRECTION map[string]string = map[string]string{
	"ASC":  "asc",
	"DESC": "desc",
}

var TIME_PERIOD map[string]string = map[string]string{
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

type CoinsRequest struct {
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
	TimePeriod     string `json:"timePeriod"`
	Tier           string `json:"tier"`
	Tag            string `json:"tag"`
}

func GetCryptoCoins(coinsRequest CoinsRequest) CoinsResponseBody {
	var queryParameters = []string{}
	if coinsRequest.Limit != 0 {
		queryParameters = append(queryParameters, "limit="+strconv.Itoa(coinsRequest.Limit))
	}
	if coinsRequest.Offset != 0 {
		queryParameters = append(queryParameters, "offset="+strconv.Itoa(coinsRequest.Offset))
	}
	if coinsRequest.OrderBy != "" {
		queryParameters = append(queryParameters, "orderBy="+coinsRequest.OrderBy)
	}
	if coinsRequest.OrderDirection != "" {
		queryParameters = append(queryParameters, "orderDirection="+coinsRequest.OrderDirection)
	}
	if coinsRequest.TimePeriod != "" {
		queryParameters = append(queryParameters, "timePeriod="+coinsRequest.TimePeriod)
	}
	if coinsRequest.Tier != "" {
		queryParameters = append(queryParameters, "tier="+coinsRequest.Tier)
	}
	if coinsRequest.Tag != "" {
		queryParameters = append(queryParameters, "tags="+coinsRequest.Tag)
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

	return coinsResponseBody
}

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

func GetCryptoCoin(id string) CoinResponseBody {
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

	return coinResponseBody
}
