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

type TopHeadlinesRequest struct {
	ApiKey   string   `json:"apiKey"`
	Sources  []string `json:"sources"`
	Q        string   `json:"q"`
	Category string   `json:"category"`
	Country  string   `json:"country"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
}

type TopHeadlinesResponseBody struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Articles []Article `json:"articles"`
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

	body, getError := http.Get(url)
	if getError != nil {
		return nil, getError
	}

	var topHeadlinesResponseBody TopHeadlinesResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &topHeadlinesResponseBody)
	if jsonUnmarshalError != nil {
		return nil, jsonUnmarshalError
	}

	if topHeadlinesResponseBody.Status != "ok" {
		return []Article{}, errors.New(topHeadlinesResponseBody.Message)
	}

	return topHeadlinesResponseBody.Articles, nil
}
