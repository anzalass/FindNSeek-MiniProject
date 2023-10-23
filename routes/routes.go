package routes

import (
	"findnseek/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, uc controller.UserController) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
	e.GET("/user", uc.MyProfile(), middleware.JWT([]byte("anzalasganteng")))
}

func RouteItem(e *echo.Echo, ic controller.ItemController) {
	e.POST("/item", ic.CreateItem())
	e.GET("/item/:page", ic.GetItemsWithPaginationAndSearch())
}

func RoutePengajuan(e *echo.Echo, pc controller.PengajuanController) {
	e.POST("/pengajuan", pc.CreatePengajuan())
}
