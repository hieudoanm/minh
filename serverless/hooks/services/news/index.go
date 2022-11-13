package news

import (
	"fmt"
	"log"
	"strings"

	"chatbot-webhook/clients/news"
	"chatbot-webhook/utils"
)

var API_KEY_NEWS = utils.Getenv("API_KEY_NEWS", "")

func GetMessage(articles []news.Article) string {
	var border string = "--------------------------------"
	var newsArticles []string = []string{}
	for i, article := range articles {
		var index = i + 1
		var source = article.Source.Name
		var title = article.Title
		var url = article.Url
		var newsArticle = fmt.Sprintf("%d. [%s - %s](%s)", index, source, title, url)
		newsArticles = append(newsArticles, newsArticle)
	}
	var message = strings.Join(newsArticles, fmt.Sprintf("\n%s\n", border))
	return fmt.Sprintf("%s\n%s\n%s", border, message, border)
}

func GetTopHeadlinesMessage() string {
	articles, articlesError := news.GetTopHeadlines(
		news.TopHeadlinesRequest{
			ApiKey:   API_KEY_NEWS,
			Sources:  []string{},
			Q:        "",
			Category: "",
			Country:  "",
			Page:     0,
			PageSize: 0,
		},
	)
	if articlesError != nil {
		return "News: N/A"
	}
	return GetMessage(articles)
}

func GetBlockchainCryptoMessage() string {
	articles, articlesError := news.GetEverything(
		news.EverythingRequest{
			ApiKey:   API_KEY_NEWS,
			Sources:  []string{},
			Q:        "blockchain+crypto",
			Page:     1,
			PageSize: 10,
		},
	)
	if articlesError != nil {
		log.Println("articlesError", articlesError)
		return "News: N/A"
	}
	log.Println("articles", len(articles))
	return GetMessage(articles)
}
