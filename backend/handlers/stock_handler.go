package handlers

import (
	"net/http"

	"github.com/ChungYenYu/20260429_stock/backend/models"
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
		c.JSON(http.StatusOK, stock)
	}
}
