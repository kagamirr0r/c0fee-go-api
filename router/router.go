package router

import (
	"c0fee-api/controller"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupMiddleware(e *echo.Echo) {
	// Add body dump middleware for debugging (only for non-multipart requests)
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		contentType := c.Request().Header.Get("Content-Type")

		// ヘッダー情報を表示
		fmt.Printf("=== REQUEST HEADERS ===\n")
		for name, values := range c.Request().Header {
			for _, value := range values {
				fmt.Printf("Header: %s: %s\n", name, value)
			}
		}

		// multipart/form-dataの場合は詳細なボディ出力をスキップ
		if contentType != "" && (strings.Contains(contentType, "multipart/form-data")) {
			fmt.Printf("Request Body: [multipart/form-data - binary content skipped]\n")
			fmt.Printf("Content-Type: %s\n", contentType)
		} else if contentType != "" {
			fmt.Printf("Content-Type: %s\n", contentType)
		}

		fmt.Printf("Response Body: %s\n", string(resBody))
		fmt.Printf("=====================\n")
	}))

	// Add request logger middleware
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("REQUEST: %s %s, Status: %d\n", v.Method, v.URI, v.Status)
			return nil
		},
	}))

	e.Use(ValidateAuthorization)
}

func defineRoutes(e *echo.Echo, uc controller.IUserController, bc controller.IBeanController, cc controller.ICountryController, rc controller.IRoasterController, ac controller.IAreaController, fc controller.IFarmController, vc controller.IVarietyController, pmc controller.IProcessMethodController, rlc controller.IRoastLevelController) {
	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Read)
	e.GET("/users/:id/beans", uc.ListUserBeans)
	e.GET("/beans/:id", bc.Read)
	e.POST("/beans", bc.Create)
	e.PUT("/beans/:id", bc.Update)
	e.GET("/countries", cc.List)
	e.GET("/countries/:id", cc.Read)
	e.GET("/countries/:id", cc.Read)
	e.GET("/roasters", rc.List)
	e.GET("/areas/:id", ac.Read)
	e.GET("/farms/:id", fc.Read)
	e.GET("/varieties", vc.List)
	e.GET("/process-methods", pmc.List)
	e.GET("/roast-levels", rlc.GetAllRoastLevels)

	// e.POST("/logout", uc.LogOut)
}

func NewRouter(uc controller.IUserController, bc controller.IBeanController, cc controller.ICountryController, rc controller.IRoasterController, ac controller.IAreaController, fc controller.IFarmController, vc controller.IVarietyController, pmc controller.IProcessMethodController, rlc controller.IRoastLevelController) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()} //custom_validator.go

	// Setup Middleware
	setupMiddleware(e)

	// Define Routes
	defineRoutes(e, uc, bc, cc, rc, ac, fc, vc, pmc, rlc)

	return e
}
