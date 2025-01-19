package router

import (
	"c0fee-api/controller"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

var jst, _ = time.LoadLocation("Asia/Tokyo")

// Validate the input struct
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fieldErrors := make([]controller.FieldError, len(validationErrors))

		for i, fieldErr := range validationErrors {
			fieldErrors[i] = controller.FieldError{
				Field:   fieldErr.Field(),
				Message: fieldErr.Tag(),
			}
		}

		errorResponse := controller.Response{
			Code:      "BAD_REQUEST",
			Message:   "Invalid request parameters",
			Errors:    fieldErrors,
			Content:   nil,
			Timestamp: time.Now().In(jst).Format(time.RFC3339),
		}
		return echo.NewHTTPError(http.StatusBadRequest, errorResponse)
	}
	return nil
}
