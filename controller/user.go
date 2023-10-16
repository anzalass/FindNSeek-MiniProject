package controller

import (
	"findnseek/middleware"
	"findnseek/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
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
		var res = uc.model.Register(input)
		if res == nil {
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

		token, err := middleware.CreateToken(login.ID, login.Name)
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
		// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1MjIwOTMsImlkIjoyLCJuYW1lIjoiZ2VtcGFyIn0.Z1tP2hSrZy9nzxPFbS4panekZZnOHraQK0ycbf7B2vM"
		var token2 = c.Get("user")
		// var token = c.Request().Header

		// res, err := middleware.ExtractToken(token2)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, map[string]any{
		// 		"message": "bad request",
		// 	})
		// }

		// name := res["name"].(string)

		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			// "token":   res,
			"token2": token2,
		})

	}
}
