package router

import (
	"c0fee-api/controller"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupMiddleware(e *echo.Echo) {
	// Add body dump middleware for debugging
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		headers := c.Request().Header
		for name, values := range headers {
			for _, value := range values {
				fmt.Printf("Header: %s: %s\n", name, value)
			}
		}
		fmt.Printf("Request Body: %s\nResponse Body: %s\n", string(reqBody), string(resBody))
	}))

	// Add request logger middleware
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42) // Example custom value
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			customValue := c.Get("customValueFromContext")
			fmt.Printf("REQUEST: URI: %s, Status: %d, Custom Value: %v\n", v.URI, v.Status, customValue)
			return nil
		},
	}))

	e.Use(ValidateAuthorization)
}

func defineRoutes(e *echo.Echo, uc controller.IUserController, bc controller.IBeanController, cc controller.ICountryController, rc controller.IRoasterController, ac controller.IAreaController, fc controller.IFarmController, vc controller.IVarietyController) {
	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Read)
	e.GET("/users/:id/beans", uc.ListUserBeans)
	e.GET("/beans/:id", bc.Read)
	e.GET("/countries", cc.List)
	e.GET("/countries/:id", cc.Read)
	e.GET("/countries/:id", cc.Read)
	e.GET("/roasters", rc.List)
	e.GET("/areas/:id", ac.Read)
	e.GET("/farms/:id", fc.Read)
	e.GET("/varieties", vc.List)

	// e.POST("/logout", uc.LogOut)
}

func NewRouter(uc controller.IUserController, bc controller.IBeanController, cc controller.ICountryController, rc controller.IRoasterController, ac controller.IAreaController, fc controller.IFarmController, vc controller.IVarietyController) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()} //custom_validator.go

	// Setup Middleware
	setupMiddleware(e)

	// Define Routes
	defineRoutes(e, uc, bc, cc, rc, ac, fc, vc)

	return e
}
