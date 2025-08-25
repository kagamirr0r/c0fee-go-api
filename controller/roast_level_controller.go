package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"c0fee-api/common"
	"c0fee-api/usecase"
)

type IRoastLevelController interface {
	GetAllRoastLevels(c echo.Context) error
}

type roastLevelController struct {
	rlu usecase.IRoastLevelUsecase
}

func (rlc *roastLevelController) GetAllRoastLevels(c echo.Context) error {
	roastLevels, err := rlc.rlu.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.GenerateErrorResponse("INTERNAL_ERROR", "Failed to get roast levels", nil))
	}

	return c.JSON(http.StatusOK, roastLevels)
}

func NewRoastLevelController(rlu usecase.IRoastLevelUsecase) IRoastLevelController {
	return &roastLevelController{rlu}
}
