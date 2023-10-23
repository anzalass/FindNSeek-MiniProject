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

type PersetujuanControllerInterface interface {
	CreatePersetujuan() echo.HandlerFunc
	GetPersetujuanByID() echo.HandlerFunc
}

type PersetujuanController struct {
	mdl model.PersetujuanInterface
}

func NewPersetujuanModelController(m model.PersetujuanInterface) PersetujuanControllerInterface {
	return &PersetujuanController{
		mdl: m,
	}
}

func (pc *PersetujuanController) InitPersetujuanController(pm model.PersetujuanModel) {
	pc.mdl = &pm
}

func (pc *PersetujuanController) CreatePersetujuan() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input = model.Persetujuan{}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Invalid user input",
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

		db := model.InitModel()

		id := id_user["id"].(string)
		var item = model.Item{}
		if err := db.First(&item, "id = ?", input.Id_Item).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("cant get item email, %s", err.Error()),
			})
		}

		if item.Status == 1 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "anda sudah melakukan persetujuan sebelumnya",
			})
		}

		var to = model.Pengajuan{}
		if err := db.First(&to, "id = ?", input.Id_Pengajuan).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("cant get user email, %s", err.Error()),
			})
		}

		sendEmail := middleware.SendEmailPersetujuan(to.Email, item.NoHp, item.Email)
		if sendEmail != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal mengirim email, %s", err.Error()),
			})
		}

		input.Id_User = id
		input.ID = uuid.NewString()

		res, err := pc.mdl.CreatePersetujuan(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("anda sudah melakukan persetujuan sebelumnya"),
			})
		}

		if err := db.Model(&model.Item{}).Where("id = ?", input.Id_Item).Update("status", 1).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal update status, %s", err.Error()),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sukses membuat persetujuan",
			"data":    res,
			"meta": map[string]interface{}{
				"email_owner": item.Email,
				"email_to":    to.Email,
			},
		})
	}
}

func (pc *PersetujuanController) GetPersetujuanByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_items := c.Param("id")
		if id_items == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "invalid id",
			})
		}

		res, err := pc.mdl.GetPersetujuanByID(id_items)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": fmt.Sprintf("gagal mendapatkan persetujuan by id, %s", err.Error()),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "sukses mendapatkan data persetujuan by id",
			"data":    res,
		})

	}
}
