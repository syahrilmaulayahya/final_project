package product_controller

import (
	"final_project/configs"
	"final_project/models/products"
	"final_project/models/responses"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProductController(c echo.Context) error {
	products := []products.Product{}
	result := configs.DB.Preload("Review_Rating").Preload("Product_description").Preload("Product_type").Preload("Size").Find(&products)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika mendapatkan data produk dari DB",
				Data:    nil,
			})

		}
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data produk user dari DB",
		Data:    products,
	})

}
