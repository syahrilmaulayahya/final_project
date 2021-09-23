package routes

import (
	"final_project/constans"
	"final_project/controllers/product_controller"
	"final_project/controllers/user_controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	jwt := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(constans.SECRET_JWT),
	})
	ev1 := e.Group("api/v1/users")
	ev1.GET("", user_controller.GetUserController, jwt)
	ev1.POST("/registers", user_controller.UserRegisterController)
	ev1.POST("/login", user_controller.LoginController)
	ev1.GET("/address", user_controller.GetAddressController, jwt)
	ev1.PUT("/updates", user_controller.UpdateController, jwt)
	ev2 := e.Group("api/v1/products")
	ev2.GET("", product_controller.GetProductController)

	return e
}
