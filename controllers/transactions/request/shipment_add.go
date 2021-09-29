package requests

import "final_project/business/transactions"

type ShipmentAdd struct {
	Name           string  `json:"name"`
	Shipment_Type  string  `json:"shipment_type"`
	Shipment_Price float64 `json:"shipment_price"`
}

func (shipmentAdd *ShipmentAdd) ToDomain() transactions.ShipmentDomain {
	return transactions.ShipmentDomain{
		Name:           shipmentAdd.Name,
		Shipment_Type:  shipmentAdd.Shipment_Type,
		Shipment_Price: shipmentAdd.Shipment_Price,
	}
}
