package routes

import (
	"B51-MICRO-FEATURE-GO/controller"
	"B51-MICRO-FEATURE-GO/pkg/mysql"
	"B51-MICRO-FEATURE-GO/services"

	"github.com/labstack/echo/v4"
)

func ArticleRoutes(e *echo.Group) {
	r := services.ArticleRepository(mysql.DB)
	cr := controller.ArticleController(r)

	e.GET("/articles", cr.FindArticles)
	e.GET("/article/:id", cr.GetArticle)
}
