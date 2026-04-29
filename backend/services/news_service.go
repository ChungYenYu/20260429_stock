package services

import (
	"strings"
	"time"
	"github.com/ChungYenYu/20260429_stock/backend/models"
)

func FetchLatestNews() []models.News {
	newsList := []models.News{
		{
			Title:       "TSMC Announces Record Quarterly Earnings",
			Summary:     "TSMC has reported a significant increase in its quarterly earnings, driven by strong demand for AI chips.",
			URL:         "https://example.com/tsmc-earnings",
			Source:      "Market Watch",
			PublishedAt: time.Now().Add(-1 * time.Hour),
		},
		{
			Title:       "Taiwan Inflation Rate Steady in April",
			Summary:     "The inflation rate in Taiwan remained steady in April, meeting economist expectations.",
			URL:         "https://example.com/taiwan-inflation",
			Source:      "Financial Times",
			PublishedAt: time.Now().Add(-5 * time.Hour),
		},
		{
			Title:       "Foxconn Expands EV Production in Southeast Asia",
			Summary:     "Foxconn is increasing its investment in electric vehicle production facilities in Southeast Asia.",
			URL:         "https://example.com/foxconn-ev",
			Source:      "Reuters",
			PublishedAt: time.Now().Add(-10 * time.Hour),
		},
	}

	for i := range newsList {
		AnalyzeNewsImpact(&newsList[i])
	}

	return newsList
}

func AnalyzeNewsImpact(news *models.News) {
	title := strings.ToLower(news.Title)
	if strings.Contains(title, "tsmc") {
		news.ImpactAnalysis = "Positive"
		news.ImpactedStocks = "[\"2330\"]"
	} else if strings.Contains(title, "foxconn") || strings.Contains(title, "hon hai") {
		news.ImpactAnalysis = "Positive"
		news.ImpactedStocks = "[\"2317\"]"
	} else if strings.Contains(title, "inflation") {
		news.ImpactAnalysis = "Neutral"
		news.ImpactedStocks = "[]"
	} else {
		news.ImpactAnalysis = "Neutral"
		news.ImpactedStocks = "[]"
	}
}
