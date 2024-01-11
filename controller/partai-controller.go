package controller

import (
	dto "B51-MICRO-FEATURE-GO/dto/result"
	service "B51-MICRO-FEATURE-GO/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type partaiController struct {
	PartaiRepository service.PartaiService
}

func PartaiController(partaiServices service.PartaiService) *partaiController {
	return &partaiController{partaiServices}
}

func (cpt *partaiController) FindPartais(c echo.Context) error {
	partais, err := cpt.PartaiRepository.GetPartais()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: partais})
}

func (cpt *partaiController) GetPartai(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	partai, err := cpt.PartaiRepository.GetPartaiById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: partai})
}
