package controller

import (
	dto "B51-MICRO-FEATURE-GO/dto/result"
	service "B51-MICRO-FEATURE-GO/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type voteController struct {
	voteService service.VoteService
}

func VoteController(voteService service.VoteService) *voteController {
	return &voteController{voteService}
}

func (vc *voteController) FindVotes(c echo.Context) error {
	votes, err := vc.voteService.GetVotes()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: votes,
	})
}

func (vc *voteController) CreateVote(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	vote, err := vc.voteService.GetVoteById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: vote,
	})
}
