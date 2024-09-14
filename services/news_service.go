package services

import (
	"errors"

	"github.com/joshua468/news-aggregator-api/clients"
	"github.com/joshua468/news-aggregator-api/models"
	"github.com/joshua468/news-aggregator-api/repository"
)

type NewsService struct {
	newsRepo   *repository.NewsRepository
	newsClient *clients.NewsAPIClient
}

func NewNewsService(newsRepo *repository.NewsRepository, newsClient *clients.NewsAPIClient) *NewsService {
	return &NewsService{
		newsRepo:   newsRepo,
		newsClient: newsClient,
	}
}

// FetchAndSaveNews fetches news from the external API and saves the articles to the database.
func (s *NewsService) FetchAndSaveNews(country string) error {
	if country == "" {
		return errors.New("country parameter is required")
	}

	// Fetch news articles from the external API
	articles, err := s.newsClient.FetchNews(country)
	if err != nil {
		return err
	}

	// Save each article to the database
	for _, article := range articles {
		newsArticle := models.Article{
			Title:  article.Title,
			Author: article.Author,
			URL:    article.URL,
		}

		err := s.newsRepo.SaveArticle(newsArticle)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetSavedArticles retrieves all saved articles from the database.
func (s *NewsService) GetSavedArticles() ([]models.Article, error) {
	return s.newsRepo.GetArticles()
}
