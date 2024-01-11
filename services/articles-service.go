package services

import (
	"B51-MICRO-FEATURE-GO/models"

	"gorm.io/gorm"
)

type ArticleService interface {
	GetArticles() ([]models.Article, error)
	GetArticleById(ID int) (models.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func ArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) GetArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *articleRepository) GetArticleById(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.First(&article, ID).Error
	return article, err
}
