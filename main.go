package main

import (
	"findnseek/config"
	"findnseek/controller"
	"findnseek/middleware"
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

	var config = config.InitConfig()

	db := model.InitModel(*config)
	model.Migrate(db)

	userModel := model.UserModel{}
	userModel.Init(db)
	userController := controller.UserController{}
	userController.InitUserController(userModel)

	itemModel := model.ItemModel{}
	itemModel.Init(db)
	itemController := controller.ItemController{}
	itemController.InitItemController(itemModel)

	pengajuanModel := model.PengajuanModel{}
	pengajuanModel.Init(db)
	pengajuanController := controller.PengajuanController{}
	pengajuanController.InitPengajuanController(pengajuanModel)

	persetujuanModel := model.PersetujuanModel{}
	persetujuanModel.Init(db)
	persetujuanController := controller.PersetujuanController{}
	persetujuanController.InitPersetujuanController(persetujuanModel)

	routes.RouteUser(e, userController)
	routes.RouteItem(e, itemController)
	routes.RoutePengajuan(e, pengajuanController)
	routes.RoutePersetujuan(e, persetujuanController, itemController)

	e.Logger.Fatal(e.Start(":8000").Error())
	middleware.LoggerMiddleware(e)

}
