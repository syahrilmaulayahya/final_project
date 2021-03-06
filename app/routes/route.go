package routes

import (
	"final_project/controllers/admins"
	"final_project/controllers/products"
	"final_project/controllers/transactions"
	"final_project/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController        users.UserController
	ProductController     products.ProductController
	TransactionController transactions.TransactionController
	AdminController       admins.AdminController
	JWTMiddleware         middleware.JWTConfig
	JWTAdmin              middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("users/logins", cl.UserController.Login)
	e.GET("users/details/:id", cl.UserController.Details, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.POST("users/registers", cl.UserController.Register)
	e.POST("users/reviews", cl.UserController.UploadReview, middleware.JWTWithConfig(cl.JWTMiddleware))

	e.GET("products", cl.ProductController.Get)
	e.GET("products/details/:id", cl.ProductController.Details)
	e.GET("products/searches", cl.ProductController.Search)
	e.GET("products/filters", cl.ProductController.FilterByType)
	e.POST("products/uploads", cl.ProductController.UploadProduct, middleware.JWTWithConfig(cl.JWTAdmin))
	e.PUT("products/updates", cl.ProductController.UpdateProduct, middleware.JWTWithConfig(cl.JWTAdmin))

	e.POST("products/uploadtypes", cl.ProductController.UploadType, middleware.JWTWithConfig(cl.JWTAdmin))

	e.POST("products/uploadsizes", cl.ProductController.UploadSize, middleware.JWTWithConfig(cl.JWTAdmin))
	e.PUT("products/updatesizes", cl.ProductController.UpdateSize, middleware.JWTWithConfig(cl.JWTAdmin))
	e.PUT("products/updatestocks", cl.ProductController.UpdateStock, middleware.JWTWithConfig(cl.JWTAdmin))

	e.POST("products/uploaddescriptions", cl.ProductController.UploadDescription, middleware.JWTWithConfig(cl.JWTAdmin))
	e.PUT("products/updatedescriptions", cl.ProductController.UpdateDescription, middleware.JWTWithConfig(cl.JWTAdmin))

	e.POST("transactions/addshoppingcarts", cl.TransactionController.Add, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("transactions/details", cl.TransactionController.DetailSC, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.POST("transactions/details/checkouts", cl.TransactionController.Checkout, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.POST("transactions/details/checkouts/pns", cl.TransactionController.ChoosePnS, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.POST("transactions/details/checkouts/pns/pay", cl.TransactionController.Pay, middleware.JWTWithConfig(cl.JWTMiddleware))

	e.POST("transactions/addpayments", cl.TransactionController.AddPM, middleware.JWTWithConfig(cl.JWTAdmin))
	e.GET("transactions/getpayments", cl.TransactionController.GetPM)

	e.POST("transactions/addshipments", cl.TransactionController.AddShipment, middleware.JWTWithConfig(cl.JWTAdmin))
	e.GET("transactions/getshipments", cl.TransactionController.GetShipment)

	e.GET("transactions/transactiondetails", cl.TransactionController.GetTransDetail, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.POST("transactions/transactiondetails/delivered", cl.TransactionController.Delivered, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.POST("transactions/transactiondetails/canceled", cl.TransactionController.Canceled, middleware.JWTWithConfig(cl.JWTMiddleware))

	e.POST("admins/registers", cl.AdminController.Register)
	e.POST("admins/logins", cl.AdminController.Login)
}
