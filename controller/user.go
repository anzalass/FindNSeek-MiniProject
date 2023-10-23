package controller

import (
	"findnseek/middleware"
	"findnseek/model"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	MyProfile() echo.HandlerFunc
}

type UserController struct {
	model model.UserInterface
}

func NewUserModelControllerInstance(m model.UserInterface) UserControllerInterface {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) InitUserController(um model.UserModel) {
	uc.model = &um
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.User{}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Invalid user input",
			})
		}

		input.ID = uuid.NewString()
		if input.Name == "" || input.Email == "" || input.Password == "" || input.ID == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Invalid user input",
			})
		}
		res, err := uc.model.Register(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "error register user",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sukses registsasi",
			"data":    res,
		})
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.User{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Invalid user input",
			})
		}

		login, err := uc.model.Login(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "username or password incorrect",
			})
		}

		token, err := middleware.CreateToken(login.ID, login.Name, login.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Error login",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "login berhasil",
			"token":   token,
		})

	}
}

func (uc *UserController) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {

		var token2 = c.Request().Header.Get("Authorization")
		tokenWithoutBearer := strings.TrimPrefix(token2, "Bearer ")

		res, err := middleware.ExtractToken(tokenWithoutBearer)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "bad request",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "berhasil mendapatkan data",
			"data":    res,
		})

	}
}
