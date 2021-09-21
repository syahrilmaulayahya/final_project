package routes

import (
	"final_project/controllers/user_controller"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/users")
	ev1.GET("", user_controller.GetUserController)
	ev1.POST("/registers", user_controller.UserRegisterController)

	return e
}
