package router

import (
	"c0fee-api/controller"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	// For Debugging
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		headers := c.Request().Header
		for name, values := range headers {
			for _, value := range values {
				println(name, value)
			}
		}
		fmt.Printf("request body: %v, response body: %v", string(reqBody), string(resBody))
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/signin", uc.SignIn)
	// e.POST("/logout", uc.LogOut)

	return e
}
