package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"c0fee-api/dto"
	"c0fee-api/usecase"
)

type RoastLevelController interface {
	GetAllRoastLevels(c echo.Context) error
}

type roastLevelController struct {
	rlu usecase.RoastLevelUsecase
}

func (rlc *roastLevelController) GetAllRoastLevels(c echo.Context) error {
	roastLevels := rlc.rlu.GetAllRoastLevels()

	res := dto.RoastLevelsOutput{
		RoastLevels: roastLevels,
		Count:       uint(len(roastLevels)),
	}

	return c.JSON(http.StatusOK, res)
}

func NewRoastLevelController(rlu usecase.RoastLevelUsecase) RoastLevelController {
	return &roastLevelController{rlu}
}
