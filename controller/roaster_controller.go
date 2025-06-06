package controller

import (
	"c0fee-api/common"
	"c0fee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IRoasterController interface {
	List(c echo.Context) error
}

type roasterController struct {
	ru usecase.IRoasterUsecase
}

func (rc *roasterController) List(c echo.Context) error {
	var params common.QueryParams

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resCountries, err := rc.ru.List(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resCountries)
}

func NewRoasterController(bu usecase.IRoasterUsecase) IRoasterController {
	return &roasterController{bu}
}
