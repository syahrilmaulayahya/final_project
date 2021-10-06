package requests

import "final_project/business/products"

type SizeUpdate struct {
	ID        int    `json:"id"`
	ProductID int    `json:"productid"`
	Type      string `json:"type"`
	Size      string `json:"size"`
	Stock     int    `json:"stock"`
}

func (sizeUpdate *SizeUpdate) ToDomain() products.SizeDomain {
	return products.SizeDomain{
		ID:        sizeUpdate.ID,
		ProductID: sizeUpdate.ProductID,
		Type:      sizeUpdate.Type,
		Size:      sizeUpdate.Size,
	}
}
