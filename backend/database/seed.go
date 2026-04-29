package database

import (
	"log"

	"github.com/ChungYenYu/20260429_stock/backend/models"
)

func SeedData() {
	stocks := []models.Stock{
		{Code: "0056", Name: "Yuanta High Dividend", Price: 38.5},
		{Code: "00919", Name: "Capital Care Taiwan Select High Dividend ETF", Price: 25.3},
		{Code: "2330", Name: "TSMC", Price: 780.0},
		{Code: "2317", Name: "Hon Hai Precision Industry", Price: 150.0},
		{Code: "00878", Name: "Cathay MSCT Taiwan ESG Sustainability High Dividend Yield ETF", Price: 22.1},
		{Code: "0050", Name: "Yuanta/P-Shares Taiwan Top 50 ETF", Price: 155.0},
		{Code: "00713", Name: "Yuanta Taiwan High Dividend Low Volatility ETF", Price: 52.0},
	}

	for _, s := range stocks {
		var stock models.Stock
		if err := DB.Where("code = ?", s.Code).First(&stock).Error; err != nil {
			log.Printf("Seeding stock: %s", s.Code)
			if err := DB.Create(&s).Error; err != nil {
				log.Printf("Failed to seed stock %s: %v", s.Code, err)
			}
		} else {
			log.Printf("Stock %s already exists, skipping", s.Code)
		}
	}
}
