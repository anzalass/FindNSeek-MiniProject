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
