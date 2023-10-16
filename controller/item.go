package controller

import (
	"findnseek/middleware"
	_ "findnseek/middleware"
	"findnseek/model"
	"fmt"
	"strconv"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ItemControllerInterface interface {
	CreateItem() echo.HandlerFunc
	GetItemsWithPaginationAndSearch() echo.HandlerFunc

	// Login() echo.HandlerFunc
}

type ItemController struct {
	mdl model.ItemInterface
}

func NewItemModelControllerInstance(m model.ItemInterface) ItemControllerInterface {
	return &ItemController{
		mdl: m,
	}
}

func (ic *ItemController) InitItemController(im model.ItemModel) {
	ic.mdl = &im
}

func (ic *ItemController) CreateItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.Item{}
		file, err := c.FormFile("file")
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err.Error(),
				"message": "Invalid file",
			})
		}
		url, err := middleware.ImageUploader(src)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "Invalid file",
			})
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Invalid user input",
			})
		}

		input.Foto = url
		input.ID = uuid.NewString()
		var res = ic.mdl.CreateItem(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "error register user",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sukses publish item",
			"data":    res,
			"url":     url,
		})
	}
}

func (ic *ItemController) GetItemsWithPaginationAndSearch() echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.FormValue("search")
		page := c.Param("page")
		pagee, err := strconv.Atoi(page)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err.Error(),
				"message": "Invalid page parameter",
			})
		}
		perPage := 2

		var res = ic.mdl.GetItemsWithPaginationAndSearch(search, pagee, perPage)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "server error",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": fmt.Sprintf("sukses with page: %s and search : %s", page, search),
			"data":    res,
		})

	}
}
