package controller

import (
	"bytes"
	"encoding/json"
	"findnseek/model"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type PengajuanMockModel struct{}

func (cmm *PengajuanMockModel) CreatePengajuan(data model.Pengajuan) (*model.Pengajuan, error) {
	return &data, nil
}
func (cmm *PengajuanMockModel) GetPengajuanByItemId(itemId string) ([]model.Result, error) {
	var result []model.Result
	return result, nil
}
func (cmm *PengajuanMockModel) GetPengajuanById(id string) (*model.Pengajuan, error) {
	var result *model.Pengajuan
	return result, nil
}
func (cmm *PengajuanMockModel) CekStatusItemFromPengajuan(id string) (*model.Item, error) {
	var result *model.Item
	return result, nil
}
func (cmm *PengajuanMockModel) GetUserNameForSend(id string) (*model.User, error) {
	var result *model.User
	return result, nil
}

func TestCreatePengajuan(t *testing.T) {
	t.Run("create pengajuan", func(t *testing.T) {
		var mdl = PengajuanMockModel{}
		var ctl = NewPengajuanModelController(&mdl)
		var e = echo.New()

		var fakeFileContents = []byte("file.jpg")
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("id", "abcd123")
		writer.WriteField("id_user", "abcd123")
		writer.WriteField("judul", "abcd123")
		writer.WriteField("kategori", "abcd123")
		writer.WriteField("tanggal", "abcd123")
		writer.WriteField("lokasi", "abcd123")
		writer.WriteField("foto", "abcd123")
		writer.WriteField("no_hp", "abcd123")
		writer.WriteField("email", "abcd123")
		writer.WriteField("alamat", "abcd123")
		writer.WriteField("deskripsi", "abcd123")
		part, _ := writer.CreateFormFile("file", "file.jpg")
		part.Write(fakeFileContents)
		writer.Close()

		e.POST("/pengajuan/:id", ctl.CreatePengajuan())
		req := httptest.NewRequest(http.MethodPost, "/pengajuan/:id", body)
		req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id") // Perbaiki set path ke "/:id"
		c.SetParamNames("id")
		c.SetParamValues("e0003061-5366-4aae-9c4b-8c0d7ef2970b")
		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)
	})

}

func TestGetPengajuanByItemId(t *testing.T) {
	t.Run("GetPengajuanByItemId", func(t *testing.T) {
		var model = PengajuanMockModel{}
		var ctl = NewPengajuanModelController(&model)
		var e = echo.New()

		e.GET("/pengajuans/:id", ctl.GetPengajuanByItemId())
		req := httptest.NewRequest(http.MethodGet, "/pengajuans/:id", bytes.NewReader([]byte(`{"id_items":"6a30ca5b-bd31-42e8-946f-f9194c7a8981"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath(":id") // Perbaiki set path ke "/:id"
		c.SetParamNames("id")
		c.SetParamValues("6a30ca5b-bd31-42e8-946f-f9194c7a8981")

		e.ServeHTTP(res, req)
		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "berhasil mendapatkan data", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)

	})
	t.Run("GetPengajuanByItemIdError", func(t *testing.T) {
		var model = PengajuanMockModel{}
		var ctl = NewPengajuanModelController(&model)
		var e = echo.New()

		e.GET("/pengajuans", ctl.GetPengajuanByItemId())
		req := httptest.NewRequest(http.MethodGet, "/pengajuans", bytes.NewReader([]byte(`{"id_items":"6a30ca5b-bd31-42e8-946f-f9194c7a8981"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath(":id") // Perbaiki set path ke "/:id"
		c.SetParamNames("id")
		c.SetParamValues("6a30ca5b-bd31-42e8-946f-f9194c7a8981")

		e.ServeHTTP(res, req)
		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "id invalid", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)

	})

}
func TestGetPengajuanById(t *testing.T) {
	t.Run("GetPengajuanById", func(t *testing.T) {
		var model = PengajuanMockModel{}
		var ctl = NewPengajuanModelController(&model)
		var e = echo.New()

		e.GET("/pengajuan/:id", ctl.GetPengajuanById())
		req := httptest.NewRequest(http.MethodGet, "/pengajuan/:id", bytes.NewReader([]byte(`{"id_items":"6a30ca5b-bd31-42e8-946f-f9194c7a8981"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath("/:id") // Perbaiki set path ke "/:id"
		c.SetParamNames("id")
		c.SetParamValues("6a30ca5b-bd31-42e8-946f-f9194c7a8981")

		e.ServeHTTP(res, req)
		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "berhasil mendapatkan data", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)

	})
	t.Run("GetPengajuanByIdError", func(t *testing.T) {
		var model = PengajuanMockModel{}
		var ctl = NewPengajuanModelController(&model)
		var e = echo.New()

		e.GET("/pengajuan", ctl.GetPengajuanById())
		req := httptest.NewRequest(http.MethodGet, "/pengajuan", bytes.NewReader([]byte(`{"id_items":"6a30ca5b-bd31-42e8-946f-f9194c7a8981"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath("") // Perbaiki set path ke "/:id"
		c.SetParamNames("")
		c.SetParamValues("6a30ca5b-bd31-42e8-946f-f9194c7a8981")

		e.ServeHTTP(res, req)
		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "id invalid", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}(nil)), tmp.Data)

	})
}
