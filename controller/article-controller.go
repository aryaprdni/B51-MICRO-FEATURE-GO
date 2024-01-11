package controller

import (
	dto "B51-MICRO-FEATURE-GO/dto/result"
	service "B51-MICRO-FEATURE-GO/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type articleController struct {
	ArticleRepository service.ArticleService
}

func ArticleController(articleServices service.ArticleService) *articleController {
	return &articleController{articleServices}
}

func (cr *articleController) FindArticles(c echo.Context) error {
	articles, err := cr.ArticleRepository.GetArticles()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: articles})
}

func (cr *articleController) GetArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := cr.ArticleRepository.GetArticleById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: article})
}
