package controller

import (
	"bytes"
	"encoding/json"
	_ "encoding/json"
	"findnseek/middleware"
	_ "findnseek/middleware"
	"findnseek/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type UserMockModel struct{}

func (umm *UserMockModel) Register(data model.User) (*model.User, error) {
	return &data, nil
}

func (umm *UserMockModel) Login(data model.User) (*model.User, error) {
	return &data, nil
}

func TestRegister(t *testing.T) {
	t.Run("Sukses", func(t *testing.T) {
		var model = UserMockModel{}
		var ctl = NewUserModelControllerInstance(&model)

		var e = echo.New()
		e.POST("/register", ctl.Register())
		var req = httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader([]byte(`{"id":"123", "name":"anzlas", "email":"enz@gmail.com", "password":"123"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		fmt.Println(resData)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "sukses registsasi", tmp.Message)
		assert.Equal(t, "anzlas", tmp.Data["name"])

	})
	t.Run("Invalid User Input", func(t *testing.T) {
		var model = UserMockModel{}
		var ctl = NewUserModelControllerInstance(&model)

		var e = echo.New()
		e.POST("/register", ctl.Register())
		var req = httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader([]byte(`{"email":"enz@gmail.com", "password":"123"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		fmt.Println(resData)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid user input", tmp.Message)

	})
	t.Run("Invalid User Input2", func(t *testing.T) {
		var model = UserMockModel{}
		var ctl = NewUserModelControllerInstance(&model)

		var e = echo.New()
		e.POST("/register", ctl.Register())
		var req = httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader([]byte(``)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		fmt.Println(resData)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid user input", tmp.Message)

	})
}

func TestLogin(t *testing.T) {
	t.Run("Sukses login", func(t *testing.T) {

		var model = UserMockModel{}
		var ctl = NewUserModelControllerInstance(&model)

		var e = echo.New()
		e.POST("/login", ctl.Login())
		var req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte(`{"id":"123", "name":"anzalas", "email":"enz@gmail.com", "password":"123"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		token, _ := middleware.CreateToken("123", "anzalas", "enz@gmail.com")

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
			Token   string                 `json:"token"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		fmt.Println(resData)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "login berhasil", tmp.Message)
		assert.Equal(t, token, tmp.Token)

	})
	t.Run("Invalid User Input login", func(t *testing.T) {

		var model = UserMockModel{}
		var ctl = NewUserModelControllerInstance(&model)

		var e = echo.New()
		e.POST("/login", ctl.Login())
		var req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte(`"salah":"salah"`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		// token, _ := middleware.CreateToken("123", "anzalas", "enz@gmail.com")

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
			Token   string                 `json:"token"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		fmt.Println(resData)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid user input", tmp.Message)

	})
}

func TestGetProfile(t *testing.T) {
	t.Run("Sukses get profile", func(t *testing.T) {
		var model = UserMockModel{}
		var ctl = NewUserModelControllerInstance(&model)
		var e = echo.New()

		e.GET("/user", ctl.MyProfile())
		var req = httptest.NewRequest(http.MethodGet, "/user", bytes.NewReader([]byte(``)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		fmt.Println(resData)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "berhasil mendapatkan data", tmp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}{}), tmp.Data)

	})
}
