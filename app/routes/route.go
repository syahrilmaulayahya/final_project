package routes

import (
	"final_project/controllers/products"
	"final_project/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController    users.UserController
	ProductController products.ProductController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/login", cl.UserController.Login)
	e.GET("users/details", cl.UserController.Details)
	e.POST("users/registers", cl.UserController.Register)

	e.GET("products", cl.ProductController.Get)
	e.POST("products/uploads", cl.ProductController.UploadProduct)
	e.PUT("products/updates", cl.ProductController.UpdateProduct)

	e.POST("products/uploadtypes", cl.ProductController.UploadType)

	e.POST("products/uploadsizes", cl.ProductController.UploadSize)
	e.PUT("products/updatesizes", cl.ProductController.UpdateSize)
	e.PUT("products/updatestocks", cl.ProductController.UpdateStock)

	e.POST("products/uploaddescriptions", cl.ProductController.UploadDescription)
	e.PUT("products/updatedescriptions", cl.ProductController.UpdateDescription)
}
