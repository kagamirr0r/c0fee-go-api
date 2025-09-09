package controller

import (
	"c0fee-api/common"
	"c0fee-api/dto"
	"c0fee-api/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type IRoasterController interface {
	List(c echo.Context) error
	Read(c echo.Context) error
	Create(c echo.Context) error
}

type roasterController struct {
	ru usecase.IRoasterUsecase
}

func (rc *roasterController) List(c echo.Context) error {
	var params common.QueryParams

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resCountries, err := rc.ru.List(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resCountries)
}

func (rc *roasterController) Read(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	roaster, err := rc.ru.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, roaster)
}

func (rc *roasterController) Create(c echo.Context) error {
	// ユーザーIDを取得（認証ミドルウェアから）
	userID := c.Request().Header.Get("X-C0fee-User-ID")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "User ID is required",
		})
	}

	var req dto.RoasterFormInput
	// JSON文字列をBindで取得
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to bind request data: %v", err),
		})
	}

	// 画像ファイルをFormFileで取得
	file, err := c.FormFile("image")
	if err != nil {
		//ファイルがなくてもエラーにしない
		file = nil
	}
	req.ImageFile = file

	// JSONデータをパース
	var data dto.RoasterInput
	if err := json.Unmarshal([]byte(req.Data), &data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("invalid JSON data: %v", err),
		})
	}

	// Echo の Validation
	if err := c.Validate(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("validation failed: %v", err),
		})
	}

	// Roasterを作成
	roaster, err := rc.ru.Create(userID, data, req.ImageFile)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		if strings.Contains(err.Error(), "already exists") {
			return c.JSON(http.StatusConflict, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to create roaster: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, roaster)
}

func NewRoasterController(bu usecase.IRoasterUsecase) IRoasterController {
	return &roasterController{bu}
}
