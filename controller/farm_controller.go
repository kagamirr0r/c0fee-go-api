package controller

import (
	"c0fee-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IFarmController interface {
	Read(c echo.Context) error
}

type farmController struct {
	cu usecase.IFarmUsecase
}

func (cc *farmController) Read(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid farm ID")
	}

	resFarm, err := cc.cu.Read(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resFarm)
}

func NewFarmController(bu usecase.IFarmUsecase) IFarmController {
	return &farmController{bu}
}
