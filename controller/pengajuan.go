package controller

import (
	"findnseek/middleware"
	"findnseek/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PengajuanControllerInterface interface {
	CreatePengajuan() echo.HandlerFunc
	GetPengajuanByItemId() echo.HandlerFunc
	GetPengajuanById() echo.HandlerFunc
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
		var id_item = c.Param("id")

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

		item, err := pc.mdl.CekStatusItemFromPengajuan(id_item)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf(err.Error()),
			})
		}

		if item.Status == 1 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "barang sudah ditemukan tidak bisa lagi ada pengajuan",
			})

		}

		if id_user["id"].(string) == item.Id_User {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "anda tidak bisa membuat pengajuan barang di barang yang anda sendiri",
			})
		}
		input.Foto = url
		input.ID = uuid.NewString()
		input.Id_User = id_user["id"].(string)
		input.Email = id_user["email"].(string)
		input.Id_Item = id_item

		if input.ID == "" || input.Id_User == "" || input.Judul == "" || input.Kategori == "" || input.Id_Item == "" || input.Tanggal == "" || input.Lokasi == "" || input.Foto == "" || input.Alamat == "" || input.Deskripsi == "" || input.Email == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "invalid input",
			})
		}

		res, err := pc.mdl.CreatePengajuan(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal membuat postingan, %s", err.Error()),
			})
		}

		user, err := pc.mdl.GetUserNameForSend(item.Id_User)
		if err != nil {
			logrus.Error("error get email", err.Error())
		}

		//
		hasl := middleware.SendEmailPenngajuan(input.Foto, item.Email, fmt.Sprintf("http://http://13.210.70.71:8000/pengajuan/%s", input.ID), user.Name)
		if hasl != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal mengirim email, %s", err.Error()),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sukses membuat pengajuan",
			"data":    res,
		})

	}
}

func (pc *PengajuanController) GetPengajuanByItemId() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_item := c.Param("id")
		if id_item == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "id invalid",
			})
		}

		res, err := pc.mdl.GetPengajuanByItemId(id_item)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": fmt.Sprintf("gagal mendapatkan data %s", err.Error()),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "berhasil mendapatkan data",
			"data":    res,
		})

	}
}

func (pc *PengajuanController) GetPengajuanById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "id invalid",
			})
		}
		res, err := pc.mdl.GetPengajuanById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": fmt.Sprintf("gagal mendapatkan data %s", err.Error()),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "berhasil mendapatkan data",
			"data":    res,
		})
	}
}
