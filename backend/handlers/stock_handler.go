package handlers

import (
	"net/http"
	"time"

	"github.com/ChungYenYu/20260429_stock/backend/models"
	"github.com/ChungYenYu/20260429_stock/backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStock(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		var stock models.Stock
		if err := db.Where("code = ?", code).First(&stock).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update price if older than 5 minutes
		if time.Since(stock.UpdatedAt) > 5*time.Minute {
			newPrice, err := services.FetchRealTimePrice(stock.Code)
			if err == nil {
				stock.Price = newPrice
				db.Save(&stock)
			}
		}

		c.JSON(http.StatusOK, stock)
	}
}

func SearchStocks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Query("q")
		var stocks []models.Stock
		query := db.Limit(10)
		if q != "" {
			query = query.Where("code LIKE ? OR name LIKE ?", "%"+q+"%", "%"+q+"%")
		}
		if err := query.Find(&stocks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stocks)
	}
}

func GetStockHistory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")

		// We could verify if the stock exists in DB first, but FetchHistoricalData will fail if the code is invalid anyway.
		// However, it's better to check DB to ensure we only serve stocks we tracking or just let it fetch.
		// The spec says "Call service, return JSON array."

		history, err := services.FetchHistoricalData(code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, history)
	}
}
