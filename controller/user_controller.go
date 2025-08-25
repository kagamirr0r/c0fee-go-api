package controller

import (
	"c0fee-api/common"
	"c0fee-api/model"
	"c0fee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	ListUserBeans(c echo.Context) error
	// LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
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
		return c.JSON(http.StatusInternalServerError, err.Error())
		// var fieldErrors []common.FieldError
		// if errors.Is(err, repository.ErrDuplicateId) {
		// 	fieldErrors = []common.FieldError{{Field: "id", Message: "ID already exists"}}
		// 	return c.JSON(http.StatusConflict, common.GenerateErrorResponse("CONFLICT", "Validation failed", fieldErrors))
		// }
		// if errors.Is(err, repository.ErrDuplicateName) {
		// 	fieldErrors = []common.FieldError{{Field: "name", Message: "Name already exists"}}
		// 	return c.JSON(http.StatusConflict, common.GenerateErrorResponse("CONFLICT", "Validation failed", fieldErrors))
		// }
		// fieldErrors = []common.FieldError{{Field: "", Message: err.Error()}}
		// return c.JSON(http.StatusInternalServerError, common.GenerateErrorResponse("INTERNAL_SERVER_ERROR", "Something went wrong", fieldErrors))
	}
	return c.JSON(http.StatusCreated, resUser)
}

func (uc *userController) Read(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resUser, err := uc.uu.Read(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resUser)
}

func (uc *userController) ListUserBeans(c echo.Context) error {
	var params common.QueryParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resBeans, err := uc.uu.GetUserBeans(user, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resBeans)
}

// TODO: Delete User
// func (uc *userController) LogOut(c echo.Context) error {
// 	cookie := new(http.Cookie)
// 	cookie.Name = "token"
// 	cookie.Value = ""
// 	cookie.Expires = time.Now()
// 	cookie.Path = "/"
// 	cookie.Domain = os.Getenv("API_DOMEIN")

// 	cookie.HttpOnly = true
// 	cookie.SameSite = http.SameSiteNoneMode
// 	c.SetCookie(cookie)
// 	return c.NoContent(http.StatusOK)
// }
