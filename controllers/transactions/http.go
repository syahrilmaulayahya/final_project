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
