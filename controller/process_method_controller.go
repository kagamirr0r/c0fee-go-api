package controller

import (
	"c0fee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IProcessMethodController interface {
	List(c echo.Context) error
}

type processMethodController struct {
	pmu usecase.IProcessMethodUsecase
}

func (pmc *processMethodController) List(c echo.Context) error {
	resProcessMethods, err := pmc.pmu.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resProcessMethods)
}

func NewProcessMethodController(pmu usecase.IProcessMethodUsecase) IProcessMethodController {
	return &processMethodController{pmu}
}
