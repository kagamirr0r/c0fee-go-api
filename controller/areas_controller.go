package controller

import (
	"c0fee-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IAreaController interface {
	Read(c echo.Context) error
}

type areaController struct {
	cu usecase.IAreaUsecase
}

func (cc *areaController) Read(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid area ID")
	}

	resArea, err := cc.cu.Read(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resArea)
}

func NewAreaController(bu usecase.IAreaUsecase) IAreaController {
	return &areaController{bu}
}
