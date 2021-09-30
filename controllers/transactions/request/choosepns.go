package requests

import "final_project/business/transactions"

type ChoosePnS struct {
	ID               int `json:"id"`
	Payment_MethodID int `json:"payment_methodid"`
	ShipmentID       int `json:"shipmentid"`
}

func (pns *ChoosePnS) ToDomain() transactions.TransactionDomain {
	return transactions.TransactionDomain{
		ID:               pns.ID,
		Payment_MethodID: pns.Payment_MethodID,
		ShipmentID:       pns.ShipmentID,
	}
}
