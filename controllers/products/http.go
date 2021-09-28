package products

import (
	"errors"
	"final_project/business/products"
	"final_project/controllers"
	"final_project/controllers/products/requests"
	"final_project/controllers/products/respons"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductUseCase products.UseCase
}

func NewProductController(productUseCase products.UseCase) *ProductController {
	return &ProductController{
		ProductUseCase: productUseCase,
	}
}

func (ProductController ProductController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	product, err := ProductController.ProductUseCase.Get(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListFromDomain(product))
}
func (ProductController ProductController) Details(c echo.Context) error {
	id, fail := strconv.Atoi(c.Param("id"))
	if fail != nil {
		return errors.New("gagal konversi id")
	}
	ctx := c.Request().Context()
	product, err := ProductController.ProductUseCase.Details(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.FromDomain(product))
}
func (ProductController ProductController) Search(c echo.Context) error {
	searchProduct := requests.Product_Search{}
	c.Bind(&searchProduct)
	ctx := c.Request().Context()
	result, err := ProductController.ProductUseCase.Search(ctx, searchProduct.Name)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListFromDomain(result))

}
func (ProductController ProductController) FilterByType(c echo.Context) error {

	filterProduct, fail := strconv.Atoi(c.QueryParam("product_typeid"))
	if fail != nil {
		return errors.New("gagal konversi product type id")
	}

	ctx := c.Request().Context()
	result, err := ProductController.ProductUseCase.FilterByType(ctx, filterProduct)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListFromDomain(result))
}
func (ProductController ProductController) UploadProduct(c echo.Context) error {
	newProduct := requests.ProductUpload{}
	c.Bind(&newProduct)
	uploadProduct := newProduct.ToDomain()
	ctx := c.Request().Context()
	product, err := ProductController.ProductUseCase.UploadProduct(ctx, uploadProduct)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.ProductFromDomain(product))
}
func (ProductController ProductController) UpdateProduct(c echo.Context) error {
	newproduct := requests.ProductUpdate{}
	c.Bind(&newproduct)
	updateProduct := newproduct.ToDomain()
	ctx := c.Request().Context()
	product, err := ProductController.ProductUseCase.UpdateProduct(ctx, updateProduct, newproduct.ID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.ProductFromDomain(product))
}

func (ProductController ProductController) UploadType(c echo.Context) error {
	newProductType := requests.ProductTypeUpload{}
	c.Bind(&newProductType)
	uploadType := newProductType.ToDomain()
	ctx := c.Request().Context()
	productType, err := ProductController.ProductUseCase.UploadType(ctx, uploadType)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.TypeFromDomain(productType))

}

func (ProductController ProductController) UploadSize(c echo.Context) error {
	newSize := requests.SizeUpload{}
	c.Bind(&newSize)
	uploadSize := newSize.ToDomain()
	ctx := c.Request().Context()
	size, err := ProductController.ProductUseCase.UploadSize(ctx, uploadSize)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.SizeFromDomain(size))
}
func (ProductController ProductController) UpdateSize(c echo.Context) error {
	newSize := requests.SizeUpdate{}
	c.Bind(&newSize)
	updateSize := newSize.ToDomain()
	ctx := c.Request().Context()
	size, err := ProductController.ProductUseCase.UpdateSize(ctx, updateSize, newSize.ID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.SizeFromDomain(size))
}

func (ProductController ProductController) UpdateStock(c echo.Context) error {
	newStock := requests.StockUpload{}
	c.Bind(&newStock)
	updateStock := newStock.ToDomain()
	ctx := c.Request().Context()
	stock, err := ProductController.ProductUseCase.UpdateStock(ctx, updateStock.Stock, updateStock.ID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.SizeFromDomain(stock))
}

func (ProductController ProductController) UploadDescription(c echo.Context) error {
	newDescription := requests.Product_descriptionUpload{}
	c.Bind(&newDescription)
	uploadDescription := newDescription.ToDomain()
	ctx := c.Request().Context()
	description, err := ProductController.ProductUseCase.UploadDescription(ctx, uploadDescription)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.Product_descriptionResponse(description))
}

func (ProductController ProductController) UpdateDescription(c echo.Context) error {
	newDescription := requests.Product_descriptionUpload{}
	c.Bind(&newDescription)
	updateDescription := newDescription.ToDomain()
	ctx := c.Request().Context()
	description, err := ProductController.ProductUseCase.UpdateDescription(ctx, updateDescription, updateDescription.ProductID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.Product_descriptionResponse(description))
}
