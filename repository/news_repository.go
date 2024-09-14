package repository

import (
	"github.com/joshua468/news-aggregator-api/models"
	"gorm.io/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{db}
}

// SaveArticle saves an article in the database using GORM.
func (repo *NewsRepository) SaveArticle(article models.Article) error {
	result := repo.db.Create(&article)
	return result.Error
}

// GetArticles retrieves all saved articles using GORM.
func (repo *NewsRepository) GetArticles() ([]models.Article, error) {
	var articles []models.Article
	result := repo.db.Find(&articles)
	return articles, result.Error
}
