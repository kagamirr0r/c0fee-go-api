package router

import (
	"c0fee-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/signin", uc.SignIn)
	// e.POST("/logout", uc.LogOut)

	return e
}
