package handlers

import (
	articlesdto "B51-MICRO-FEATURE-GO/dto/articles"
	dto "B51-MICRO-FEATURE-GO/dto/result"
	"B51-MICRO-FEATURE-GO/models"
	"B51-MICRO-FEATURE-GO/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	ArticleRepository repositories.ArticleRepository
}

func ArticleHandler(articleRepository repositories.ArticleRepository) *articleHandler {
	return &articleHandler{articleRepository}
}

func (h *articleHandler) FindArticles(c echo.Context) error {
	articles, err := h.ArticleRepository.FindArticles()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: articles})
}

func (h *articleHandler) GetArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.ArticleRepository.GetArticle(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertArticleRes(article)})
}

func convertArticleRes(c models.Article) articlesdto.ArticleRes {
	return articlesdto.ArticleRes{
		ID:          c.ID,
		Title:       c.Title,
		Image:       c.Image,
		Description: c.Description,
	}
}
