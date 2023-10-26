package controller

import (
	"bytes"
	"encoding/json"
	"findnseek/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type PersetujuanMock struct{}

func (pm *PersetujuanMock) CreatePersetujuan(data model.Persetujuan) (*model.Persetujuan, error) {
	return &data, nil
}
func (pm *PersetujuanMock) GetPersetujuanByID(id string) ([]model.Res, error) {
	res := []model.Res{}
	return res, nil
}

func TestCreatePersetujuan(t *testing.T) {
	t.Run("create persertujuan", func(t *testing.T) {
		var mdl = PersetujuanMock{}
		var ctl = NewPersetujuanModelController(&mdl)
		var e = echo.New()

		e.PUT("/persetujuan", ctl.CreatePersetujuan())
		req := httptest.NewRequest(http.MethodPut, "/persetujuan", bytes.NewReader([]byte(`{"id_item":"123fdf" , "id_pengajuan":"123fdf"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)

		res := httptest.NewRecorder()
		e.NewContext(req, res)

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		resData.Decode(&tmp)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)

	})
	t.Run("bad request create persetujuan", func(t *testing.T) {
		var mdl = PersetujuanMock{}
		var ctl = NewPersetujuanModelController(&mdl)
		var e = echo.New()

		e.PUT("/persetujuan", ctl.CreatePersetujuan())
		req := httptest.NewRequest(http.MethodPut, "/persetujuan", bytes.NewReader([]byte(``)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)

		res := httptest.NewRecorder()
		e.NewContext(req, res)
		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "id not found", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)

	})
}

func TestGetPersetujuanByidItem(t *testing.T) {
	t.Run("get persetujuan by id item sukses", func(t *testing.T) {
		var mdl = PersetujuanMock{}
		var ctl = NewPersetujuanModelController(&mdl)
		var e = echo.New()

		uri := "/persetujuan/:id"
		e.GET(uri, ctl.GetPersetujuanByID())
		req := httptest.NewRequest(http.MethodGet, uri, nil)
		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)

		res := httptest.NewRecorder()
		e.ServeHTTP(res, req)

		c := e.NewContext(req, res)
		c.SetPath(":id")
		c.SetParamNames("id")
		c.SetParamValues("6a30ca5b-bd31-42e8-946f-f9194c7a8981")

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		resData.Decode(&tmp)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "sukses mendapatkan data persetujuan by id", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)

	})

	t.Run("get persetujuan by id item gagal", func(t *testing.T) {
		var mdl = PersetujuanMock{}
		var ctl = NewPersetujuanModelController(&mdl)
		var e = echo.New()

		uri := "/persetujuan"
		e.GET(uri, ctl.GetPersetujuanByID())
		req := httptest.NewRequest(http.MethodGet, uri, nil)
		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)

		res := httptest.NewRecorder()
		e.ServeHTTP(res, req)
		e.NewContext(req, res)

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "invalid id", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)

	})
}
