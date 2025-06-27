package controller

import (
	"c0fee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IVarietyController interface {
	List(c echo.Context) error
}

type varietyController struct {
	vu usecase.IVarietyUsecase
}

func (vc *varietyController) List(c echo.Context) error {
	resVarieties, err := vc.vu.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resVarieties)
}

func NewVarietyController(vu usecase.IVarietyUsecase) IVarietyController {
	return &varietyController{vu}
}