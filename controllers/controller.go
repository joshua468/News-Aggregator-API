package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/news-aggregator-api/services"
)

type NewsController struct {
	newsService *services.NewsService
}

func NewNewsController(newsService *services.NewsService) *NewsController {
	return &NewsController{newsService}
}

func (ctrl *NewsController) FetchNews(c *gin.Context) {
	country := c.Query("country")
	if country == "" {
		country = "us" // default to US if no country is provided
	}
	err := ctrl.newsService.FetchAndSaveNews(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "News fetched and saved successfully"})
}

func (ctrl *NewsController) GetSavedNews(c *gin.Context) {
	articles, err := ctrl.newsService.GetSavedArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve saved news"})
		return
	}
	c.JSON(http.StatusOK, articles)
}
