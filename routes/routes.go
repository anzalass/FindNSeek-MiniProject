package routes

import (
	"findnseek/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controller.UserController) {
	e.POST("/register", uc.Register())                                    // done
	e.POST("/login", uc.Login())                                          // done
	e.GET("/user", uc.MyProfile(), echojwt.JWT([]byte("anzalasganteng"))) // done
}

func RouteItem(e *echo.Echo, ic controller.ItemController) {
	e.POST("/item", ic.CreateItem(), echojwt.JWT([]byte("anzalasganteng"))) //done
	e.GET("/item/", ic.GetItemsWithPaginationAndSearch())
	e.PUT("/item/:id", ic.UpdateStatusItem(), echojwt.JWT([]byte("anzalasganteng")))    // done
	e.GET("/item/:id", ic.GetItemsByIdWithPengajuanAndPersetujuan())                    // done
	e.PUT("/items/:id", ic.UpdateItems(), echojwt.JWT([]byte("anzalasganteng")))        // done
	e.DELETE("/items/:id", ic.DeleteItemsById(), echojwt.JWT([]byte("anzalasganteng"))) // done
}

func RoutePengajuan(e *echo.Echo, pc controller.PengajuanController) {
	e.POST("/pengajuan/:id", pc.CreatePengajuan(), echojwt.JWT([]byte("anzalasganteng")))
	e.GET("/pengajuans/:id", pc.GetPengajuanByItemId(), echojwt.JWT([]byte("anzalasganteng")))
	e.GET("/pengajuan/:id", pc.GetPengajuanById(), echojwt.JWT([]byte("anzalasganteng")))
}

func RoutePersetujuan(e *echo.Echo, pc controller.PersetujuanController, ic controller.ItemController) {
	e.PUT("/persetujuan/:id", pc.CreatePersetujuan(), echojwt.JWT([]byte("anzalasganteng")))
	e.GET("/persetujuan/:id", pc.GetPersetujuanByID(), echojwt.JWT([]byte("anzalasganteng")))
}
