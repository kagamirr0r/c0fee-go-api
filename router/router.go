package router

import (
	"c0fee-api/controller"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func generateErrorResponse(code string, message string) controller.Response {
	errorResponse := controller.Response{
		Code:      code,
		Message:   message,
		Errors:    []controller.FieldError{},
		Content:   nil,
		Timestamp: time.Now().In(jst).Format(time.RFC3339),
	}
	return errorResponse
}

func ValidateAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	jwtSrecret := os.Getenv("SUPABASE_JWT_SECRET")
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, generateErrorResponse("Unauthorized", "Invalid Request Parameters"))
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// HS256 アルゴリズムを使用しているか確認
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("予期しない署名方法: %v", token.Header["alg"])
			}
      // SupabaseのJWTシークレットを使用
      return []byte(jwtSrecret), nil
    })

		if err != nil {
			return c.JSON(http.StatusBadRequest, generateErrorResponse("BAD_REQUEST", "Invalid Request Parameters"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userMetadata, metaOk := claims["user_metadata"].(map[string]interface{})
			sub, subOk := userMetadata["sub"].(string)
			userId := c.Request().Header.Get("X-C0fee-User-ID")
			if sub != userId || !metaOk || !subOk {
				return c.JSON(http.StatusBadRequest, generateErrorResponse("BAD_REQUEST", "User ID mismatch"))
			}
		}
		return next(c)
	}
}

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

func defineRoutes(e *echo.Echo, uc controller.IUserController) {
	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Show)
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
