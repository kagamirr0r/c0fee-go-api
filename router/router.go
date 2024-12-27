package router

import (
	"c0fee-api/controller"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			errorResponse := controller.Response{
				Code:      "BAD_REQUEST",
				Message:   "Invalid request parameters.",
				Errors:    []controller.FieldError{{Field: "", Message: ""}},
				Content:   nil,
				Timestamp: time.Now().In(jst).Format(time.RFC3339),
			}
			return c.JSON(http.StatusUnauthorized, errorResponse)
		}
		return next(c)
	}
}

func setupMiddleware(e *echo.Echo) {
	e.Use(CheckAuthorization)
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
}

func defineRoutes(e *echo.Echo, uc controller.IUserController) {
	e.POST("/signup", uc.SignUp)
	e.POST("/signin", uc.SignIn)
	// e.POST("/logout", uc.LogOut)
}

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()} //custom_validator.go

	// Setup Middleware
	setupMiddleware(e)

	// Define Routes
	defineRoutes(e, uc)

	return e
}
