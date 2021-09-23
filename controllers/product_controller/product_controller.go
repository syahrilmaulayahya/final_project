package product_controller

import (
	"errors"
	"final_project/configs"
	"final_project/models/products"
	"final_project/models/responses"
	"net/http"

	"github.com/go-sql-driver/mysql"
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
		Message: "Berhasil mendapatkan data produk  dari DB",
		Data:    products,
	})

}

func UploadProductController(c echo.Context) error {
	var newProduct products.ProductUpload
	var productDB products.Product
	c.Bind(&newProduct)

	productDB.Code = newProduct.Code
	productDB.Name = newProduct.Name
	productDB.Total_Stock = newProduct.Total_Stock
	productDB.Price = newProduct.Price
	productDB.Picture_url = newProduct.Picture_url
	productDB.Product_typeID = newProduct.Product_typeID

	result := configs.DB.Create(&productDB)

	var mysqlErr *mysql.MySQLError
	if result.Error != nil {
		if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.JSON(http.StatusBadRequest, responses.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "code sudah diguanakan",
				Data:    nil,
			})
		} else if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1452 {
			return c.JSON(http.StatusBadRequest, responses.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "tipe produk belum tersedia",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input data product ke DB",
				Data:    nil,
			})
		}

	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "upload product berhasil",
		Data:    productDB,
	})
}
