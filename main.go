package main

import (
	"log"

	"github.com/joshua468/news-aggregator-api/clients"
	"github.com/joshua468/news-aggregator-api/config"
	"github.com/joshua468/news-aggregator-api/controllers"
	"github.com/joshua468/news-aggregator-api/interfaces/http"
	"github.com/joshua468/news-aggregator-api/repository"
	"github.com/joshua468/news-aggregator-api/services"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database connection
	config.ConnectDatabase()

	// Set up dependencies
	newsRepo := repository.NewNewsRepository(config.DB)
	newsClient := clients.NewNewsAPIClient()
	newsService := services.NewNewsService(newsRepo, newsClient)
	newsController := controllers.NewNewsController(newsService)

	// Set up the router
	router := http.SetupRouter(newsController)

	// Start the server
	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}
