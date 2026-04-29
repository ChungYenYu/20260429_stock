package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ChungYenYu/20260429_stock/backend/database"
	"github.com/ChungYenYu/20260429_stock/backend/handlers"
	"github.com/ChungYenYu/20260429_stock/backend/models"
	"github.com/ChungYenYu/20260429_stock/backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func startPriceUpdater(db *gorm.DB) {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			updatePrices(db)
			<-ticker.C
		}
	}()
}

func updatePrices(db *gorm.DB) {
	var stocks []models.Stock
	if err := db.Find(&stocks).Error; err != nil {
		fmt.Printf("Error fetching stocks: %v\n", err)
		return
	}

	for _, stock := range stocks {
		price, err := services.FetchRealTimePrice(stock.Code)
		if err != nil {
			fmt.Printf("Error fetching price for %s: %v\n", stock.Code, err)
			continue
		}

		if err := db.Model(&stock).Update("price", price).Error; err != nil {
			fmt.Printf("Error updating price for %s: %v\n", stock.Code, err)
		} else {
			fmt.Printf("Updated price for %s: %.2f\n", stock.Code, price)
		}
	}
}

func main() {
	// Initialize Database
	database.InitDB()
	database.SeedData()

	// Start Background Price Updater
	startPriceUpdater(database.DB)

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Stock backend is running",
		})
	})

	// Stock routes
	r.GET("/api/stock/:code", handlers.GetStock(database.DB))
	r.GET("/api/stock/:code/history", handlers.GetStockHistory(database.DB))
	r.GET("/api/stocks/search", handlers.SearchStocks(database.DB))
	r.GET("/api/news", handlers.GetNews)

	r.Run(":8080")
}
