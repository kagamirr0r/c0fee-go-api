package controller

import (
	"c0fee-api/model"
	"c0fee-api/repository"
	"c0fee-api/usecase"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Create(c echo.Context) error
	Show(c echo.Context) error
	// LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Response struct {
	Code      string       `json:"code"`
	Message   string       `json:"message"`
	Errors    []FieldError `json:"errors"`
	Content   interface{}  `json:"content"`
	Timestamp string       `json:"timestamp"`
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

var jst, _ = time.LoadLocation("Asia/Tokyo")

// 共通エラーレスポンス生成関数
func generateErrorResponse(code, message string, fieldErrors []FieldError) Response {
	if len(fieldErrors) == 0 {
			fieldErrors = []FieldError{} 
	}
	return Response{
		Code:      code,
		Message:   message,
		Errors:    fieldErrors,
		Content:   nil,
		Timestamp: time.Now().In(jst).Format(time.RFC3339),
	}
}

func (uc *userController) Create(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return err
	}

	resUser, err := uc.uu.Create(user)
	if err != nil {
		var fieldErrors []FieldError
		if errors.Is(err, repository.ErrDuplicateId) {
			fieldErrors = []FieldError{{Field: "id", Message: "ID already exists"}}
			return c.JSON(http.StatusConflict, generateErrorResponse("CONFLICT", "Validation failed", fieldErrors))
		}
		if errors.Is(err, repository.ErrDuplicateName) {
			fieldErrors = []FieldError{{Field: "name", Message: "Name already exists"}}
			return c.JSON(http.StatusConflict, generateErrorResponse("CONFLICT", "Validation failed", fieldErrors))
		}
		fieldErrors = []FieldError{{Field: "", Message: err.Error()}}
		return c.JSON(http.StatusInternalServerError, generateErrorResponse("INTERNAL_SERVER_ERROR", "Something went wrong", fieldErrors))
	}

	response := Response{
		Code:      "CREATED",
		Message:   "User created",
		Errors:    []FieldError{},
		Content:   resUser,
		Timestamp: time.Now().In(jst).Format(time.RFC3339),
	}
	return c.JSON(http.StatusCreated, response)
}

func (uc *userController) Show(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resUser, err := uc.uu.Show(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resUser)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMEIN")

	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
