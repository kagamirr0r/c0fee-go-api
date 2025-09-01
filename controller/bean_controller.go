package controller

import (
	"c0fee-api/dto"
	"c0fee-api/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type IBeanController interface {
	Read(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
}

type beanController struct {
	bu usecase.IBeanUsecase
}

func (bc *beanController) Read(c echo.Context) error {
	id := c.Param("id")

	// Convert string ID to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	resBean, err := bc.bu.Read(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resBean)
}

func (bc *beanController) Create(c echo.Context) error {
	// ユーザーIDを取得（認証ミドルウェアから）
	userID := c.Request().Header.Get("X-C0fee-User-ID")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "User ID is required",
		})
	}

	var req dto.BeanFormInput
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
	var data dto.BeanInput
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

	// Beanを作成
	bean, err := bc.bu.Create(userID, data, req.ImageFile)
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
			"error": fmt.Sprintf("Failed to create bean: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, bean)
}

func (bc *beanController) Update(c echo.Context) error {
	// Bean IDを取得
	id := c.Param("id")
	beanID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid bean ID format",
		})
	}

	// ユーザーIDを取得（認証ミドルウェアから）
	userID := c.Request().Header.Get("X-C0fee-User-ID")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "User ID is required",
		})
	}

	var req dto.BeanFormInput
	// JSON文字列をBindで取得
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to bind request data: %v", err),
		})
	}

	// 画像ファイルをFormFileで取得
	file, err := c.FormFile("image")
	if err != nil {
		// ファイルがなくてもエラーにしない
		file = nil
	}
	req.ImageFile = file

	// JSONデータをパース
	var data dto.BeanInput
	if err := json.Unmarshal([]byte(req.Data), &data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("invalid JSON data: %v", err),
		})
	}

	// Json Validation
	if err := c.Validate(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("validation failed: %v", err),
		})
	}

	// Beanを更新
	bean, err := bc.bu.Update(uint(beanID), userID, data, req.ImageFile)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}
		if strings.Contains(err.Error(), "access denied") {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		if strings.Contains(err.Error(), "validation failed") {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to update bean: %v", err),
		})
	}

	return c.JSON(http.StatusOK, bean)
}

func NewBeanController(bu usecase.IBeanUsecase) IBeanController {
	return &beanController{bu}
}
