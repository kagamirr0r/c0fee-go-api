package controller

import (
	"c0fee-api/model"
	"c0fee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IBeanController interface {
	ListByUser(c echo.Context) error
	// LogOut(c echo.Context) error
}

type beanController struct {
	bu usecase.IBeanUsecase
}

func NewBeanController(bu usecase.IBeanUsecase) IBeanController {
	return &beanController{bu}
}

func (bc *beanController) ListByUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resBean, err := bc.bu.ListByUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resBean)
}
