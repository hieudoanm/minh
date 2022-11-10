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

func Handler() {
	// Router
	// router.GET("/api/status/:service", getStatus)
	// router.GET("/api/vnindex/history/:symbol", GetVnindexHistory)
	// Start
	log.Println("🚀 Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
