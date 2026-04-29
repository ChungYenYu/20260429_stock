package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ChungYenYu/20260429_stock/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetStock(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Stock{})

	// Seed data
	testStock := models.Stock{Code: "AAPL", Name: "Apple Inc.", Price: 150.0}
	db.Create(&testStock)

	r := gin.Default()
	r.GET("/api/stock/:code", GetStock(db))

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/stock/AAPL", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var response models.Stock
		json.Unmarshal(w.Body.Bytes(), &response)

		if response.Code != "AAPL" {
			t.Errorf("Expected AAPL, got %s", response.Code)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/stock/MSFT", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", w.Code)
		}
	})
}

func TestSearchStocks(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Stock{})

	// Seed data
	stocks := []models.Stock{
		{Code: "AAPL", Name: "Apple Inc.", Price: 150.0},
		{Code: "GOOGL", Name: "Alphabet Inc.", Price: 2800.0},
		{Code: "MSFT", Name: "Microsoft Corporation", Price: 300.0},
		{Code: "AMZN", Name: "Amazon.com Inc.", Price: 3300.0},
		{Code: "TSLA", Name: "Tesla Inc.", Price: 700.0},
	}
	for _, s := range stocks {
		db.Create(&s)
	}

	r := gin.Default()
	r.GET("/api/stocks/search", SearchStocks(db))

	t.Run("SearchByCode", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/stocks/search?q=AA", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var response []models.Stock
		json.Unmarshal(w.Body.Bytes(), &response)

		if len(response) != 1 || response[0].Code != "AAPL" {
			t.Errorf("Expected only AAPL, got %v", response)
		}
	})

	t.Run("SearchByName", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/stocks/search?q=Micro", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var response []models.Stock
		json.Unmarshal(w.Body.Bytes(), &response)

		if len(response) != 1 || response[0].Code != "MSFT" {
			t.Errorf("Expected only MSFT, got %v", response)
		}
	})

	t.Run("EmptySearch", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/stocks/search?q=", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var response []models.Stock
		json.Unmarshal(w.Body.Bytes(), &response)

		// If q is empty, it might return all or none. Let's assume it returns max 10.
		// Given we have 5 stocks, it should return 5.
		if len(response) != 5 {
			t.Errorf("Expected 5 stocks, got %d", len(response))
		}
	})
}
