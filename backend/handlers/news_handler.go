package handlers

import (
	"net/http"
	"github.com/ChungYenYu/20260429_stock/backend/services"
	"github.com/gin-gonic/gin"
)

func GetNews(c *gin.Context) {
	newsList := services.FetchLatestNews()
	c.JSON(http.StatusOK, newsList)
}
