package vnindex_client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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

func GetVnindexCompanies() []CompaniesResponseBody {
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

	return companiesResponseBody
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

func GetVnindexHistory(symbol string) []HistoryResponseBody {
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
	var historyResponseBody []HistoryResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &historyResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetVnindexHistory", jsonUnmarshalError)
	}

	return historyResponseBody
}
