package controller

import (
	dto "B51-MICRO-FEATURE-GO/dto/result"
	"B51-MICRO-FEATURE-GO/services"
	service "B51-MICRO-FEATURE-GO/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type paslonController struct {
	PaslonRepository service.PaslonService
}

func PaslonController(paslonServices services.PaslonService) *paslonController {
	return &paslonController{paslonServices}
}

func (cp *paslonController) FindPaslons(c echo.Context) error {
	paslon, err := cp.PaslonRepository.GetPaslons()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: paslon,
	})
}

func (cp *paslonController) GetPaslon(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	paslon, err := cp.PaslonRepository.GetPaslonById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: paslon,
	})
}
