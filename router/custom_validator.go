package router

import (
	"c0fee-api/common"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

// Validate the input struct
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fieldErrors := make([]common.FieldError, len(validationErrors))

		for i, fieldErr := range validationErrors {
			fieldErrors[i] = common.FieldError{
				Field:   fieldErr.Field(),
				Message: fieldErr.Tag(),
			}
		}

		errorResponse := common.Response{
			Code:      "BAD_REQUEST",
			Message:   "Invalid request parameters",
			Errors:    fieldErrors,
			Content:   nil,
			TimeStamp: time.Now().In(common.Jst).Format(time.RFC3339),
		}
		return echo.NewHTTPError(http.StatusBadRequest, errorResponse)
	}
	return nil
}
