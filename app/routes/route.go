package routes

import (
	"final_project/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/login", cl.UserController.Login)
	e.GET("users", cl.UserController.Get)
	e.POST("users/registers", cl.UserController.Register)
}
