package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func Handler() {
	http.HandleFunc("/", GetVnindexCompanies)
	log.Println("🚀 Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
