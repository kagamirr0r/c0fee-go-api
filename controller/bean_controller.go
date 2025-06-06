package controller

import (
	"c0fee-api/model"
	"c0fee-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IBeanController interface {
	Read(c echo.Context) error
}

type beanController struct {
	bu usecase.IBeanUsecase
}

func (bc *beanController) Read(c echo.Context) error {
	id := c.Param("id")

	// Convert string ID to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	bean := model.Bean{ID: uint(idUint)}
	resBean, err := bc.bu.Read(bean)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resBean)
}

func NewBeanController(bu usecase.IBeanUsecase) IBeanController {
	return &beanController{bu}
}
