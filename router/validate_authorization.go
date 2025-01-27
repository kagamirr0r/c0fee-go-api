package router

import (
	"c0fee-api/controller"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func ValidateAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	jwtSrecret := os.Getenv("SUPABASE_JWT_SECRET")
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, controller.GenerateErrorResponse("Unauthorized", "Invalid Request Parameters", []controller.FieldError{}))
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

		fmt.Println("err", err)
		if err != nil {
			return c.JSON(http.StatusBadRequest, controller.GenerateErrorResponse("BAD_REQUEST", "Invalid Request Parameters", []controller.FieldError{}))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userMetadata, metaOk := claims["user_metadata"].(map[string]interface{})
			sub, subOk := userMetadata["sub"].(string)
			userId := c.Request().Header.Get("X-C0fee-User-ID")
			if sub != userId || !metaOk || !subOk {
				return c.JSON(http.StatusBadRequest, controller.GenerateErrorResponse("BAD_REQUEST", "User ID mismatch", []controller.FieldError{}))
			}
		}
		return next(c)
	}
}
