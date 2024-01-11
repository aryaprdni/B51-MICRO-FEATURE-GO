package controller

import (
	dto "B51-MICRO-FEATURE-GO/dto/result"
	service "B51-MICRO-FEATURE-GO/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userController struct {
	UserRepository service.UserService
}

func UserController(userRepository service.UserService) *userController {
	return &userController{userRepository}
}

func (cu *userController) GetUsers(c echo.Context) error {
	users, err := cu.UserRepository.GetUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users})
}

func (cu *userController) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := cu.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: user,
	})
}
