package controller

import (
	"c0fee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ICountryController interface {
	List(c echo.Context) error
}

type countryController struct {
	cu usecase.ICountryUsecase
}

func (cc *countryController) List(c echo.Context) error {
	resCountries, err := cc.cu.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resCountries)
}

func NewCountryController(bu usecase.ICountryUsecase) ICountryController {
	return &countryController{bu}
}
