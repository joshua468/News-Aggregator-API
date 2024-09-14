package http

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/news-aggregator-api/controllers"
)

func SetupRouter(newsController *controllers.NewsController) *gin.Engine {
	router := gin.Default()

	router.GET("/news", newsController.FetchNews)
	router.GET("/saved-news", newsController.GetSavedNews)

	return router
}
