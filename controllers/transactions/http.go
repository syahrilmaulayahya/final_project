package transactions

import (
	"final_project/app/middleware"
	"final_project/business/transactions"
	"final_project/controllers"

	requests "final_project/controllers/transactions/request"
	"final_project/controllers/transactions/respons"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUseCase transactions.UseCase
}

func NewTransactionController(transactionUseCase transactions.UseCase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transactionUseCase,
	}
}
func (transactioncontroller TransactionController) Add(c echo.Context) error {

	shopping_cart := requests.Shopping_CartAdd{}
	shopping_cart.UserID = middleware.GetClaimsUserId(c)
	shopping_cart.ProductID, _ = strconv.Atoi(c.QueryParam("productid"))
	shopping_cart.SizeID, _ = strconv.Atoi(c.QueryParam("sizeid"))
	shopping_cart.Quantity, _ = strconv.Atoi(c.QueryParam("quantity"))

	ctx := c.Request().Context()
	transaction, err := transactioncontroller.TransactionUseCase.Add(ctx, shopping_cart.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.FromDomain(transaction))
}

func (transactionController TransactionController) DetailSC(c echo.Context) error {
	ctx := c.Request().Context()
	listSC, err := transactionController.TransactionUseCase.DetailSC(ctx, middleware.GetClaimsUserId(c))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, listSC)
}

func (transactionController TransactionController) AddPM(c echo.Context) error {
	var payment_method requests.Payment_MethodAdd
	c.Bind(&payment_method)
	ctx := c.Request().Context()
	payment_methodAdd := payment_method.ToDomain()
	newPayment, err := transactionController.TransactionUseCase.AddPM(ctx, payment_methodAdd)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.PMFromDomain(newPayment))
}

func (transactionController TransactionController) GetPM(c echo.Context) error {
	ctx := c.Request().Context()
	payment, err := transactionController.TransactionUseCase.GetPM(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListPMFromDomain(payment))
}

func (transactionController TransactionController) AddShipment(c echo.Context) error {
	var shipment requests.ShipmentAdd
	c.Bind(&shipment)
	ctx := c.Request().Context()
	shipmentAdd := shipment.ToDomain()
	newShipment, err := transactionController.TransactionUseCase.AddShipment(ctx, shipmentAdd)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, newShipment)
}

func (transactionController TransactionController) GetShipment(c echo.Context) error {
	ctx := c.Request().Context()
	shipment, err := transactionController.TransactionUseCase.GetShipment(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListShipmentFromDomain(shipment))
}
