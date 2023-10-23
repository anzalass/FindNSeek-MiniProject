package controller

import (
	"findnseek/middleware"
	_ "findnseek/middleware"
	"findnseek/model"
	"fmt"
	"strconv"
	"strings"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ItemControllerInterface interface {
	CreateItem() echo.HandlerFunc
	GetItemsWithPaginationAndSearch() echo.HandlerFunc
	UpdateStatusItem() echo.HandlerFunc
	UpdateItems() echo.HandlerFunc
	GetItemsByIdWithPengajuanAndPersetujuan() echo.HandlerFunc
	DeleteItemsById() echo.HandlerFunc
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
				"message": fmt.Sprintf("invalid file, %s", err.Error()),
			})
		}
		url, err := middleware.ImageUploader(src)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("invalid file, %s", err.Error()),
			})
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("invalid input, %s", err.Error()),
			})
		}

		token := c.Request().Header.Get("Authorization")
		tokenWithoutBearer := strings.TrimPrefix(token, "Bearer ")
		id_user, err := middleware.ExtractToken(tokenWithoutBearer)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("cant get id_user, %s", err.Error()),
			})
		}

		input.Foto = url
		input.ID = uuid.NewString()
		input.Id_User = id_user["id"].(string)
		fmt.Println(id_user["id"].(string))

		if input.ID == "" || input.Id_User == "" || input.Judul == "" || input.Kategori == "" || input.Tanggal == "" || input.Lokasi == "" || input.Foto == "" || input.Alamat == "" || input.Deskripsi == "" || input.NoHp == "" || input.Email == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "invalid input",
			})
		}

		res, err := ic.mdl.CreateItem(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal membuat item, %s", err.Error()),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sukses membuat item",
			"data":    res,
		})
	}
}

func (ic *ItemController) GetItemsWithPaginationAndSearch() echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.QueryParam("search")
		kategori := c.QueryParam("kategori")
		page := c.QueryParam("page")
		pagee, err := strconv.Atoi(page)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "invalid page parameter",
			})
		}
		perPage := 2

		var res = ic.mdl.GetItemsWithPaginationAndSearch(search, kategori, pagee, perPage)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("something went wrong %s", err.Error()),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "sukses mendapatkan data",
			"data":    res,
			"meta": map[string]any{
				"page":     pagee,
				"total":    len(res),
				"kategori": kategori,
				"search":   search,
			},
		})

	}
}

func (ic *ItemController) UpdateStatusItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		_, res := ic.mdl.GetItemsByID(id)
		if res != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		ic.mdl.UpdateStatusItem(id)
		return c.JSON(http.StatusOK, map[string]any{
			"message": "sukses update status",
		})
	}
}

func (ic *ItemController) GetItemsByIdWithPengajuanAndPersetujuan() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		res1, err1 := ic.mdl.GetItemsByID(id)
		if err1 != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("gagal mendapatkan data, %s", err1.Error()),
				"data":    res1,
			})
		}
		res2, err2 := ic.mdl.GetPengajuanByItemId(id)
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("gagal mendapatkan data, %s", err2.Error()),
				"data":    res2,
			})
		}
		res3, err3 := ic.mdl.GetPersetujuanByID(id)
		if err3 != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("gagal mendapatkan data, %s", err3.Error()),
				"data":    res3,
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "sukses mendapatkan item",
			"meta": map[string]interface{}{
				"id items": id,
			},
			"data": map[string]any{
				"items":       res1,
				"pengajuan":   res2,
				"persetujuan": res3,
			},
		})

	}
}

func (ic *ItemController) UpdateItems() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.Item{}
		id := c.Param("id")
		file, err := c.FormFile("foto")
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("invalid file, %s", err.Error()),
			})
		}

		url, err := middleware.ImageUploader(src)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("invalid file, %s", err.Error()),
			})
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("invalid input, %s", err.Error()),
			})
		}

		//
		db := model.InitModel()
		item := model.Item{}
		if err := db.First(&item, "id = ?", id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("invalid input, %s", err.Error()),
			})
		}

		if item.Status == 1 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "barang sudah ditemukan tidak boleh melakukan edit",
			})
		}

		//

		input.ID = id
		input.Foto = url
		res, errr := ic.mdl.UpdateItemsById(input)
		if errr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal update item, %s", err.Error()),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sukses edit item",
			"data":    res,
			"meta": map[string]interface{}{
				"id": id,
			},
		})
	}
}

func (ic *ItemController) DeleteItemsById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("id")

		db := model.InitModel()
		item := model.Item{}
		if err := db.First(&item, "id = ?", id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("gagal mendapatkan status %s", err.Error()),
			})
		}

		if item.Status == 1 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("data tidak bisa dihapus karena barang sudah ditemukan "),
			})
		}

		err := ic.mdl.DeleteItemsById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal menghapus data %s", err.Error()),
			})

		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "sukses menghapus data",
		})

	}
}
