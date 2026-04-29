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
