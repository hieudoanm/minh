package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golang.org/x/exp/slices"
)

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

type ErrorResponse struct {
	Message string `json:"message"`
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

func GetStatus(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var service string = GetQueryParameter(request, "service", "")
	log.Printf("Service: %s", service)
	if service == "" {
		var errorResponse ErrorResponse = ErrorResponse{Message: "Missing Service"}
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}
	if !slices.Contains(SERVICES, service) {
		var errorResponse ErrorResponse = ErrorResponse{Message: "Invalid Service"}
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}
	var serviceStatus ServiceStatus = GetStatusByService(service)
	json.NewEncoder(writer).Encode(serviceStatus)
}

func Handler() {
	http.HandleFunc("/", GetStatus)
	log.Println("???? Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
