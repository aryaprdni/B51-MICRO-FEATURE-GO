package repositories

import (
	"B51-MICRO-FEATURE-GO/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindArticles() ([]models.Article, error)
	GetArticle(ID int) (models.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func RepositoryArticle(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) FindArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *articleRepository) GetArticle(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.First(&article, ID).Error
	return article, err
}
