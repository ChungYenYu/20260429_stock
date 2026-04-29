package main

import (
	"net/http"

	"github.com/ChungYenYu/20260429_stock/backend/database"
	"github.com/ChungYenYu/20260429_stock/backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.InitDB()
	database.SeedData()

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

	r.Run(":8080")
}
