package controller

import (
	"c0fee-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICountryController interface {
	List(c echo.Context) error
	Read(c echo.Context) error
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

func (cc *countryController) Read(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid country ID")
	}

	resCountry, err := cc.cu.Read(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resCountry)
}

func NewCountryController(bu usecase.ICountryUsecase) ICountryController {
	return &countryController{bu}
}
