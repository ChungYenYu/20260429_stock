package models

import "time"

type News struct {
	Title           string    `json:"title"`
	Summary         string    `json:"summary"`
	URL             string    `json:"url"`
	Source          string    `json:"source"`
	PublishedAt     time.Time `json:"published_at"`
	ImpactAnalysis  string    `json:"impact_analysis"`
	ImpactedStocks  string    `json:"impacted_stocks"` // JSON string of impacted codes
}
