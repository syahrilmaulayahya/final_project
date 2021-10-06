package transactions_test

import (
	"context"
	"final_project/app/middleware"
	"final_project/business/transactions"
	"final_project/business/transactions/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactionRepository mocks.Repository
var transactionService transactions.UseCase
var transactionDomain transactions.TransactionDomain
var shoppingCartDomain transactions.Shopping_CartDomain
var paymentMethodDomain transactions.Payment_MethodDomain
var shipmentDomain transactions.ShipmentDomain
var transactionDetailDomain transactions.Transaction_DetailDomain
var userToken middleware.ConfigJWT

func setup() {
	transactionService = transactions.NewTransactionUseCase(&transactionRepository, time.Hour*1, userToken)
	transactionDomain = transactions.TransactionDomain{
		ID:               1,
		Status:           "Paid",
		UserID:           1,
		Shopping_CartID:  1,
		Total_Qty:        1,
		Total_Price:      100000,
		Payment_MethodID: 1,

		ShipmentID: 1,
	}
	shoppingCartDomain = transactions.Shopping_CartDomain{
		ID:        1,
		UserID:    1,
		ProductID: 1,
		SizeID:    1,
		Quantity:  1,
		Price:     100000,
	}
	paymentMethodDomain = transactions.Payment_MethodDomain{
		ID:   1,
		Name: "OVO",
	}
	shipmentDomain = transactions.ShipmentDomain{
		ID:             1,
		Name:           "JNE",
		Shipment_Type:  "Regular",
		Shipment_Price: 50000,
	}
	transactionDetailDomain = transactions.Transaction_DetailDomain{
		UserID:         1,
		StatusShipment: "Undelivered",
		TransactionID:  1,
		ProductID:      1,
	}
}
func TestCheckout(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Checkout", func(t *testing.T) {
		transactionRepository.On("Checkout",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(transactionDomain, nil).Once()
		_, err := transactionService.Checkout(context.Background(), 1, 1)
		assert.Nil(t, err)
	})
}

func TestChoosePnS(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid ChoosePnS", func(t *testing.T) {
		transactionRepository.On("ChoosePnS",
			mock.Anything,
			mock.Anything).Return(transactionDomain, nil).Once()
		_, err := transactionService.ChoosePnS(context.Background(), transactions.TransactionDomain{
			Payment_MethodID: 1,
			ShipmentID:       1,
		})
		assert.Nil(t, err)
	})
}

func TestPay(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Pay", func(t *testing.T) {
		transactionRepository.On("Pay",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("float64")).Return(transactionDomain, nil).Once()
		_, err := transactionService.Pay(context.Background(), 1, 100000)
		assert.Nil(t, err)
	})
}

func TestAdd(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Add", func(t *testing.T) {
		transactionRepository.On("Add",
			mock.Anything,
			mock.Anything).Return(shoppingCartDomain, nil).Once()
		_, err := transactionService.Add(context.Background(), transactions.Shopping_CartDomain{
			ProductID: 1,
			SizeID:    1,
			Quantity:  1,
		})
		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Empty ProductID", func(t *testing.T) {
		transactionRepository.On("Add",
			mock.Anything,
			mock.Anything).Return(shoppingCartDomain, nil).Once()
		_, err := transactionService.Add(context.Background(), transactions.Shopping_CartDomain{
			ProductID: 0,
			SizeID:    1,
			Quantity:  1,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Empty SizeID", func(t *testing.T) {
		transactionRepository.On("Add",
			mock.Anything,
			mock.Anything).Return(shoppingCartDomain, nil).Once()
		_, err := transactionService.Add(context.Background(), transactions.Shopping_CartDomain{
			ProductID: 1,
			SizeID:    0,
			Quantity:  1,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 4 | Invalid Quantity", func(t *testing.T) {
		transactionRepository.On("Add",
			mock.Anything,
			mock.Anything).Return(shoppingCartDomain, nil).Once()
		_, err := transactionService.Add(context.Background(), transactions.Shopping_CartDomain{
			ProductID: 1,
			SizeID:    1,
			Quantity:  -1,
		})
		assert.NotNil(t, err)
	})
}

func TestDetailSC(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid DetailSC", func(t *testing.T) {
		transactionRepository.On("DetailSC",
			mock.Anything,
			mock.AnythingOfType("int")).Return([]transactions.Shopping_CartDomain{}, nil).Once()
		_, err := transactionService.DetailSC(context.Background(), 1)
		assert.Nil(t, err)
	})
}

func TestAddPM(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid AddPM", func(t *testing.T) {
		transactionRepository.On("AddPM",
			mock.Anything,
			mock.Anything).Return(paymentMethodDomain, nil).Once()
		_, err := transactionService.AddPM(context.Background(), transactions.Payment_MethodDomain{
			Name: "OVO",
		})
		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Empty Name", func(t *testing.T) {
		transactionRepository.On("AddPM",
			mock.Anything,
			mock.Anything).Return(paymentMethodDomain, nil).Once()
		_, err := transactionService.AddPM(context.Background(), transactions.Payment_MethodDomain{})
		assert.NotNil(t, err)
	})
}
func TestGetPM(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid GetPM", func(t *testing.T) {
		transactionRepository.On("GetPM",
			mock.Anything).Return([]transactions.Payment_MethodDomain{}, nil).Once()
		_, err := transactionService.GetPM(context.Background())
		assert.Nil(t, err)
	})

}

func TestAddShipment(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid AddShipment", func(t *testing.T) {
		transactionRepository.On("AddShipment",
			mock.Anything,
			mock.Anything).Return(shipmentDomain, nil).Once()
		_, err := transactionService.AddShipment(context.Background(), transactions.ShipmentDomain{
			Name:           "JNE",
			Shipment_Type:  "Regular",
			Shipment_Price: 50000,
		})
		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Empty Name", func(t *testing.T) {
		transactionRepository.On("AddShipment",
			mock.Anything,
			mock.Anything).Return(shipmentDomain, nil).Once()
		_, err := transactionService.AddShipment(context.Background(), transactions.ShipmentDomain{
			Name:           "",
			Shipment_Type:  "Regular",
			Shipment_Price: 50000,
		})
		assert.Nil(t, err)
	})
	t.Run("Test Case 3 | Empty Shipment Type", func(t *testing.T) {
		transactionRepository.On("AddShipment",
			mock.Anything,
			mock.Anything).Return(shipmentDomain, nil).Once()
		_, err := transactionService.AddShipment(context.Background(), transactions.ShipmentDomain{
			Name:           "JNE",
			Shipment_Type:  "",
			Shipment_Price: 50000,
		})
		assert.Nil(t, err)
	})
	t.Run("Test Case 4 | Empty Shipment Price", func(t *testing.T) {
		transactionRepository.On("AddShipment",
			mock.Anything,
			mock.Anything).Return(shipmentDomain, nil).Once()
		_, err := transactionService.AddShipment(context.Background(), transactions.ShipmentDomain{
			Name:          "JNE",
			Shipment_Type: "Regular",
		})
		assert.Nil(t, err)
	})

}

func TestGetShipment(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid GetShipment", func(t *testing.T) {
		transactionRepository.On("GetShipment",
			mock.Anything).Return([]transactions.ShipmentDomain{}, nil).Once()
		_, err := transactionService.GetShipment(context.Background())
		assert.Nil(t, err)
	})

}

func TestGetTransDetail(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid GetTransDetail", func(t *testing.T) {
		transactionRepository.On("GetTransDetail",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(transactionDetailDomain, transactionDomain, shoppingCartDomain, nil).Once()
		_, _, _, err := transactionService.GetTransDetail(context.Background(), 1, 1)
		assert.Nil(t, err)
	})
}
func TestDelivered(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Delivered", func(t *testing.T) {
		transactionRepository.On("Delivered",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(transactionDetailDomain, nil).Once()
		_, err := transactionService.Delivered(context.Background(), 1, 1)
		assert.Nil(t, err)
	})
}

func TestCanceled(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Canceled", func(t *testing.T) {
		transactionRepository.On("Canceled",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(transactionDetailDomain, nil).Once()
		_, err := transactionService.Canceled(context.Background(), 1, 1)
		assert.Nil(t, err)
	})
}
