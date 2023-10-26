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

type ItemMockModel struct {
}

func (im *ItemMockModel) CreateItem(data model.Item) (*model.Item, error) {
	return &data, nil
}
func (im *ItemMockModel) UpdateStatusItem(id string) error {
	return nil
}
func (im *ItemMockModel) GetItemsWithPaginationAndSearch(judul string, kategori string, page int, perPage int) []model.Item {
	data := []model.Item{}
	return data
}
func (im *ItemMockModel) GetItemsByID(id string) (*model.Item, error) {
	data := model.Item{}
	return &data, nil
}
func (im *ItemMockModel) GetPersetujuanByID(id string) (*model.Persetujuan, error) {
	data := model.Persetujuan{}
	return &data, nil
}
func (im *ItemMockModel) GetPengajuanByItemId(id string) ([]model.Pengajuan, error) {
	data := []model.Pengajuan{}
	return data, nil
}
func (im *ItemMockModel) UpdateItemsById(data model.Item) (*model.Item, error) {
	return &data, nil
}
func (im *ItemMockModel) DeleteItemsById(id string) error {
	return nil
}

func TestCreateItem(t *testing.T) {
	t.Run("sukse create item", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		var fakeFileContents = []byte("file.jpg")
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

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

		e.POST("/item", ctl.CreateItem())
		req := httptest.NewRequest(http.MethodPost, "/item", body)
		req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)
		var res = httptest.NewRecorder()
		e.NewContext(req, res)
		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "sukses membuat item", tmp.Message)
		assert.Equal(t, "abcd123", tmp.Data["judul"])

	})
	t.Run("invalid input create input", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		var fakeFileContents = []byte("file.jpg")
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		writer.WriteField("id_user", "abcd123")
		writer.WriteField("judul", "abcd123")
		writer.WriteField("kategori", "abcd123")
		writer.WriteField("tanggal", "abcd123")
		writer.WriteField("lokasi", "abcd123")
		writer.WriteField("foto", "abcd123")
		writer.WriteField("no_hp", "abcd123")
		writer.WriteField("email", "abcd123")
		writer.WriteField("alamat", "abcd123")
		// writer.WriteField("deskripsi", "abcd123")
		part, _ := writer.CreateFormFile("file", "file.jpg")
		part.Write(fakeFileContents)
		writer.Close()

		e.POST("/item", ctl.CreateItem())
		req := httptest.NewRequest(http.MethodPost, "/item", body)
		req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)
		var res = httptest.NewRecorder()
		e.NewContext(req, res)
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
		assert.Equal(t, "invalid input", tmp.Message)

	})
}

func TestGetItemsWithPaginationAndSearch(t *testing.T) {
	t.Run("invalid page parameter", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		uri := "/items"
		e.GET(uri, ctl.GetItemsWithPaginationAndSearch())
		var req = httptest.NewRequest(http.MethodGet, uri, nil)
		var res = httptest.NewRecorder()
		e.NewContext(req, res)
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
		assert.Equal(t, "invalid page parameter", tmp.Message)

	})
}

func TestUpdateItem(t *testing.T) {
	t.Run("sukses update items", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		var fakeFileContents = []byte("file.jpg")
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

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

		uri := "/items/:id"
		e.PUT(uri, ctl.UpdateItems())
		var req = httptest.NewRequest(http.MethodPut, uri, body)
		req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)
		var res = httptest.NewRecorder()

		c := e.NewContext(req, res)
		c.SetPath(":id")
		c.SetParamNames("id")
		c.SetParamValues("6a30ca5b-bd31-42e8-946f-f9194c7a8981")

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp ResponData
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "", tmp.Message)

	})
}
func TestUpdateStatusItem(t *testing.T) {
	t.Run("sukses update status items", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		uri := "/item/:id"
		e.PUT(uri, ctl.UpdateStatusItem())
		var req = httptest.NewRequest(http.MethodPut, uri, nil)

		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)
		var res = httptest.NewRecorder()

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
		assert.Equal(t, "", tmp.Message)

	})
	t.Run("gagal update status items", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		uri := "/item"
		e.PUT(uri, ctl.UpdateStatusItem())
		var req = httptest.NewRequest(http.MethodPut, uri, nil)

		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)
		var res = httptest.NewRecorder()

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

	})
}

func TestGetItemsByIdWithPengajuanAndPersetujuan(t *testing.T) {
	t.Run("berhasill", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		uri := "/item/:id"
		e.GET(uri, ctl.GetItemsByIdWithPengajuanAndPersetujuan())
		var req = httptest.NewRequest(http.MethodGet, uri, nil)

		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)
		var res = httptest.NewRecorder()

		c := e.NewContext(req, res)
		c.SetPath(":id")
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
		assert.Equal(t, "sukses mendapatkan item", tmp.Message)
	})
	t.Run("gagalll", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		uri := "/item"
		e.GET(uri, ctl.GetItemsByIdWithPengajuanAndPersetujuan())
		var req = httptest.NewRequest(http.MethodGet, uri, nil)

		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)
		var res = httptest.NewRecorder()

		e.NewContext(req, res)
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
		assert.Equal(t, "id tidak ditemukan", tmp.Message)
	})
}

func TestDeleteItemsByID(t *testing.T) {
	t.Run("delete items", func(t *testing.T) {
		var mdl = ItemMockModel{}
		var ctl = NewItemModelControllerInstance(&mdl)
		var e = echo.New()

		uri := "/items/:id"
		e.PUT(uri, ctl.DeleteItemsById())
		var req = httptest.NewRequest(http.MethodPut, uri, nil)
		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk3OTk4MDYwLCJpZCI6IjY3YTYyNGU3LTI2NTAtNGQwMy05N2YzLTNkZTMxMTY5OTlhZSIsIm5hbWUiOiJhbnphbGFzIn0.E34_6FqHU0JlrNhaAGAhQRBn0HjEPBr8dUlMy4GiFpQ"
		req.Header.Set(echo.HeaderAuthorization, token)
		var res = httptest.NewRecorder()

		c := e.NewContext(req, res)
		c.SetPath(":id")
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
		assert.Equal(t, "ini bukan barang milikmu, tidak boleh menghapusnya", tmp.Message)
	})
}
