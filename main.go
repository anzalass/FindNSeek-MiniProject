package main

import (
	"findnseek/controller"
	"findnseek/model"
	"findnseek/routes"
	_ "fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Find N Seek")
	})

	db := model.InitModel()
	model.Migrate(db)

	userModel := model.UserModel{}
	userModel.Init(db)
	userController := controller.UserController{}
	userController.InitUserController(userModel)

	routes.RouteUser(e, userController)

	e.Logger.Fatal(e.Start(":8000").Error())

}
