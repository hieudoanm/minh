package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"chatbot-webhook/libs/http"
)

const NEWS_V2 = "https://newsapi.org/v2"

type Source struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type SourcesResponseBody struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Sources []Source `json:"sources"`
}

func GetSources(
	apiKey string,
	category string,
	language string,
	country string,
) ([]Source, error) {
	// Build query parameters
	var queryParameters []string = []string{}
	if apiKey != "" {
		queryParameters = append(queryParameters, "apiKey="+apiKey)
	}
	if category != "" {
		queryParameters = append(queryParameters, "category="+category)
	}
	if language != "" {
		queryParameters = append(queryParameters, "language="+language)
	}
	if country != "" {
		queryParameters = append(queryParameters, "country="+country)
	}

	// Build url
	var url = fmt.Sprintf(
		"%s/top-headlines/sources?%s",
		NEWS_V2,
		strings.Join(queryParameters, "&"),
	)

	body, getError := http.Get(url)
	if getError != nil {
		return nil, getError
	}

	var sourcesResponseBody SourcesResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &sourcesResponseBody)
	if jsonUnmarshalError != nil {
		return nil, jsonUnmarshalError
	}

	if sourcesResponseBody.Status != "ok" {
		return []Source{}, errors.New(sourcesResponseBody.Message)
	}

	return sourcesResponseBody.Sources, nil
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type ArticlesResponseBody struct {
	Status       string    `json:"status"`
	Message      string    `json:"message"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type EverythingRequest struct {
	ApiKey         string   `json:"apiKey"`
	Language       string   `json:"language"`
	From           string   `json:"from"`
	To             string   `json:"to"`
	Page           int      `json:"page"`
	PageSize       int      `json:"pageSize"`
	Q              string   `json:"q"`
	SortBy         string   `json:"sortBy"`
	Sources        []string `json:"sources"`
	SearchIn       []string `json:"searchIn"`
	Domains        []string `json:"domains"`
	ExcludeDomains []string `json:"excludeDomains"`
}

func GetEverything(everythingRequest EverythingRequest) ([]Article, error) {
	// Assign default value
	if everythingRequest.Language == "" {
		everythingRequest.Language = "en"
	}
	if everythingRequest.SortBy == "" {
		everythingRequest.SortBy = "publishedAt"
	}
	if everythingRequest.Page == 0 {
		everythingRequest.Page = 1
	}
	if everythingRequest.PageSize == 0 {
		everythingRequest.PageSize = 100
	}
	// Build query parameters
	var queryParameters []string = []string{}
	if everythingRequest.ApiKey != "" {
		queryParameters = append(queryParameters, "apiKey="+everythingRequest.ApiKey)
	}
	if everythingRequest.Language != "" {
		queryParameters = append(queryParameters, "language="+everythingRequest.Language)
	}
	if everythingRequest.From != "" {
		queryParameters = append(queryParameters, "from="+everythingRequest.From)
	}
	if everythingRequest.To != "" {
		queryParameters = append(queryParameters, "to="+everythingRequest.To)
	}
	if everythingRequest.Page > 0 {
		queryParameters = append(queryParameters, "page="+strconv.Itoa(everythingRequest.Page))
	}
	if everythingRequest.PageSize > 0 {
		queryParameters = append(queryParameters, "pageSize="+strconv.Itoa(everythingRequest.PageSize))
	}
	if everythingRequest.Q != "" {
		queryParameters = append(queryParameters, "q="+everythingRequest.Q)
	}
	if everythingRequest.SortBy != "" {
		queryParameters = append(queryParameters, "sortBy="+everythingRequest.SortBy)
	}
	if len(everythingRequest.Domains) > 0 {
		queryParameters = append(queryParameters, "domains="+strings.Join(everythingRequest.Domains, ","))
	}
	if len(everythingRequest.ExcludeDomains) > 0 {
		queryParameters = append(queryParameters, "excludeDomains="+strings.Join(everythingRequest.ExcludeDomains, ","))
	}
	if len(everythingRequest.Sources) > 0 {
		queryParameters = append(queryParameters, "sources="+strings.Join(everythingRequest.Sources, ","))
	}
	if len(everythingRequest.SearchIn) > 0 {
		queryParameters = append(queryParameters, "searchIn="+strings.Join(everythingRequest.SearchIn, ","))
	}
	// Build url
	var url = fmt.Sprintf(
		"%s/everything?%s",
		NEWS_V2,
		strings.Join(queryParameters, "&"),
	)
	// Get
	body, getError := http.Get(url)
	if getError != nil {
		return nil, getError
	}
	// Parse Response
	var articlesResponseBody ArticlesResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &articlesResponseBody)
	if jsonUnmarshalError != nil {
		return nil, jsonUnmarshalError
	}
	// Get Articles
	if articlesResponseBody.Status != "ok" {
		return []Article{}, errors.New(articlesResponseBody.Message)
	}

	return articlesResponseBody.Articles, nil

}

type TopHeadlinesRequest struct {
	ApiKey   string   `json:"apiKey"`
	Sources  []string `json:"sources"`
	Q        string   `json:"q"`
	Category string   `json:"category"`
	Country  string   `json:"country"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
}

func GetTopHeadlines(topHeadlinesRequest TopHeadlinesRequest) ([]Article, error) {
	// Assign default value
	if topHeadlinesRequest.Category == "" {
		topHeadlinesRequest.Category = "general"
	}
	if topHeadlinesRequest.Country == "" {
		topHeadlinesRequest.Country = "us"
	}
	if topHeadlinesRequest.Page == 0 {
		topHeadlinesRequest.Page = 1
	}
	if topHeadlinesRequest.PageSize == 0 {
		topHeadlinesRequest.PageSize = 20
	}
	// Build query parameters
	var queryParameters []string = []string{}
	if topHeadlinesRequest.ApiKey != "" {
		queryParameters = append(queryParameters, "apiKey="+topHeadlinesRequest.ApiKey)
	}
	if len(topHeadlinesRequest.Sources) == 0 {
		queryParameters = append(queryParameters, "sources="+strings.Join(topHeadlinesRequest.Sources, ","))
	}
	if topHeadlinesRequest.Q != "" {
		queryParameters = append(queryParameters, "q="+topHeadlinesRequest.Q)
	}
	if topHeadlinesRequest.Category != "" {
		queryParameters = append(queryParameters, "category="+topHeadlinesRequest.Category)
	}
	if topHeadlinesRequest.Country != "" {
		queryParameters = append(queryParameters, "country="+topHeadlinesRequest.Country)
	}
	if topHeadlinesRequest.PageSize > 0 {
		queryParameters = append(queryParameters, "pageSize="+strconv.Itoa(topHeadlinesRequest.PageSize))
	}
	if topHeadlinesRequest.Page > 0 {
		queryParameters = append(queryParameters, "page="+strconv.Itoa(topHeadlinesRequest.Page))
	}
	// Build url
	var url = fmt.Sprintf(
		"%s/top-headlines?%s",
		NEWS_V2,
		strings.Join(queryParameters, "&"),
	)
	// Get
	body, getError := http.Get(url)
	if getError != nil {
		return nil, getError
	}
	// Parse Response
	var articlesResponseBody ArticlesResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &articlesResponseBody)
	if jsonUnmarshalError != nil {
		return nil, jsonUnmarshalError
	}
	// Get Articles
	if articlesResponseBody.Status != "ok" {
		return []Article{}, errors.New(articlesResponseBody.Message)
	}

	return articlesResponseBody.Articles, nil
}
