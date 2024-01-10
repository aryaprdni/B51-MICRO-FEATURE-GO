package routes

import (
	"B51-MICRO-FEATURE-GO/handlers"
	"B51-MICRO-FEATURE-GO/pkg/mysql"
	"B51-MICRO-FEATURE-GO/repositories"

	"github.com/labstack/echo/v4"
)

func ArticleRoutes(e *echo.Group) {
	r := repositories.RepositoryArticle(mysql.DB)
	h := handlers.ArticleHandler(r)

	e.GET("/articles", h.FindArticles)
	e.GET("/article/:id", h.GetArticle)
}
