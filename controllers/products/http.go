package products

import (
	"final_project/business/products"
	"final_project/controllers"
	"final_project/controllers/products/requests"
	"final_project/controllers/products/respons"
	"net/http"

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
