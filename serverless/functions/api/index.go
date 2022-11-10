package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/slices"
)

var API_KEY_FIXER string = os.Getenv("API_KEY_FIXER")

func GetQueryParameter(request *http.Request, key string, defaultValue string) string {
	var value string = request.URL.Query().Get("limit")
	if defaultValue == "" {
		return value
	}
	if defaultValue != "" && value == "" {
		return defaultValue
	}
	return value
}

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

type RatesResponseBody struct {
	Success   bool               `json:"success"`
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

func GetForexRates(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// HTTP Request
	var url string = fmt.Sprintf(
		"http://data.fixer.io/api/latest?access_key=%s",
		API_KEY_FIXER,
	)
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
	var ratesResponseBody RatesResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &ratesResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetForexRates", jsonUnmarshalError)
	}

	json.NewEncoder(writer).Encode(ratesResponseBody)
}

type HealthResponse struct {
	Status string `json:"status"`
}

func GetHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

var SERVICES []string = []string{
	"atlassian",
	"bitbucket",
	"circleci",
	"confluence",
	"discord",
	"github",
	"hedera",
	"jira-software",
	"solana",
	"trello",
	"vercel",
}

var SERVICE_NAMES map[string]string = map[string]string{
	"atlassian":     "Atlassian",
	"bitbucket":     "Bitbucket",
	"circleci":      "Circle CI",
	"confluence":    "Confluence",
	"discord":       "Discord",
	"github":        "GitHub",
	"hedera":        "Hedera",
	"jira-software": "Jira Software",
	"solana":        "Solana",
	"trello":        "Trello",
	"vercel":        "Vercel",
}

var SERVICE_URLS map[string]string = map[string]string{
	"atlassian":     "https://status.atlassian.com/api/v2/status.json",
	"bitbucket":     "https://bitbucket.status.atlassian.com/api/v2/status.json",
	"circleci":      "https://status.circleci.com/api/v2/status.json",
	"confluence":    "https://confluence.status.atlassian.com/api/v2/status.json",
	"discord":       "https://discordstatus.com/api/v2/status.json",
	"github":        "https://www.githubstatus.com/api/v2/status.json",
	"hedera":        "https://status.hedera.com/api/v2/status.json",
	"jira-software": "https://jira-software.status.atlassian.com/api/v2/status.json",
	"solana":        "https://status.solana.com/api/v2/status.json",
	"trello":        "https://trello.status.atlassian.com/api/v2/status.json",
	"vercel":        "https://www.vercel-status.com/api/v2/status.json",
}

type StatusResponseBody struct {
	Page struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Url       string `json:"url"`
		TimeZone  string `json:"time_zone"`
		UpdatedAt string `json:"updated_at"`
	} `json:"page"`
	Status struct {
		Indicator   string `json:"indicator"`
		Description string `json:"description"`
	} `json:"status"`
}

type ServiceStatus struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func getStatuses(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var serviceStatuses map[string]ServiceStatus = map[string]ServiceStatus{}
	for _, service := range SERVICES {
		serviceStatuses[service] = GetStatusByService(service)
	}
	json.NewEncoder(writer).Encode(serviceStatuses)
}

// func getStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	var service string = params.ByName("service")
// 	var serviceStatus ServiceStatus = GetStatusByService(service)
// 	json.NewEncoder(writer).Encode(serviceStatus)
// }

func GetStatusByService(service string) ServiceStatus {
	if !slices.Contains(SERVICES, service) {
		log.Fatalln("Invalid Service")
	}
	// HTTP Request
	var url string = SERVICE_URLS[service]
	log.Printf("URL: %s", url)
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetStatusByService")
	}
	defer response.Body.Close()
	// Get Body
	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		log.Println("Fail to GetStatusByService")
	}
	// Parse JSON
	var statusResponseBody StatusResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &statusResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetStatusByService")
	}

	var status bool = statusResponseBody.Status.Indicator == "none"
	var name string = SERVICE_NAMES[service]
	return ServiceStatus{Name: name, Status: status}
}

const TABLEBASE_URL = "https://raw.githubusercontent.com/hieudoanm/tablebase/master/json"

type CompaniesResponseBody struct {
	IssueShare  string `json:"issue_share"`
	ListedValue string `json:"listed_value"`
	Industry    string `json:"industry"`
	Sector      string `json:"sector"`
	Symbol      string `json:"symbol"`
	Supersector string `json:"supersector"`
	Subsector   string `json:"subsector"`
	MarketCap   string `json:"market_cap"`
	ListingDate string `json:"listing_date"`
	Name        string `json:"name"`
}

func GetVnindexCompanies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// HTTP Request
	var url string = fmt.Sprintf(
		"%s/vietnam/stock/companies.json",
		TABLEBASE_URL,
	)
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetVnindexCompanies", httpGetError)
	}
	defer response.Body.Close()
	// Get Body
	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		log.Println("Fail to GetVnindexCompanies", readBodyError)
	}
	// Parse JSON
	var companiesResponseBody []CompaniesResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &companiesResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetVnindexCompanies", jsonUnmarshalError)
	}

	json.NewEncoder(writer).Encode(companiesResponseBody)
}

type HistoryResponseBody struct {
	Date      string `json:"date"`
	Symbol    string `json:"symbol"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

func GetVnindexHistory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var symbol string = params.ByName("symbol")
	// HTTP Request
	var url string = fmt.Sprintf(
		"%s/vietnam/stock/history/%s.json",
		TABLEBASE_URL,
		strings.ToUpper(symbol),
	)
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetVnindexHistory", httpGetError)
	}
	defer response.Body.Close()
	// Get Body
	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		log.Println("Fail to GetVnindexHistory", readBodyError)
	}
	// Parse JSON
	var companiesResponseBody []HistoryResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &companiesResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetVnindexHistory", jsonUnmarshalError)
	}

	json.NewEncoder(writer).Encode(companiesResponseBody)
}

func GetYouTubeTrending(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func main() {
	// router := httprouter.New()
	// Router
	http.HandleFunc("/api/crypto/coins", GetCryptoCoins)
	http.HandleFunc("/api/forex/rates", GetForexRates)
	http.HandleFunc("/api/health", GetHealth)
	http.HandleFunc("/api/status", getStatuses)
	// router.GET("/api/status/:service", getStatus)
	http.HandleFunc("/api/vnindex/companies", GetVnindexCompanies)
	// router.GET("/api/vnindex/history/:symbol", GetVnindexHistory)
	http.HandleFunc("/api/youtube/trending", GetYouTubeTrending)
	// Start
	log.Println("🚀 Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
