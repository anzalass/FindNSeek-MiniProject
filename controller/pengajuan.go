package controller

import (
	"findnseek/middleware"
	"findnseek/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PengajuanControllerInterface interface {
	CreatePengajuan() echo.HandlerFunc
}

type PengajuanController struct {
	mdl model.PengajuanInterface
}

func NewPengajuanModelController(m model.PengajuanInterface) PengajuanControllerInterface {
	return &PengajuanController{
		mdl: m,
	}
}

func (pc *PengajuanController) InitPengajuanController(pm model.PengajuanModel) {
	pc.mdl = &pm
}

func (pc *PengajuanController) CreatePengajuan() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.Pengajuan{}

		// id_item := c.FormValue("id_item")

		token := c.Request().Header.Get("Authorization")
		tokenWithoutBearer := strings.TrimPrefix(token, "Bearer ")
		id_user, err := middleware.ExtractToken(tokenWithoutBearer)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("cant get id_user, %s", err.Error()),
			})
		}

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
				"message": fmt.Sprintf("upload file failed, %s", err.Error()),
			})
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("invalid input, %s", err.Error()),
			})
		}

		input.Foto = url
		input.ID = uuid.NewString()
		input.Id_User = id_user["id"].(string)

		var res = pc.mdl.CreatePengajuan(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal membuat postingan, %s", err.Error()),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sukses membuat pengajuan",
			"data":    res,
		})

	}
}
