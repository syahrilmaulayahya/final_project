package requests

import "final_project/business/products"

type SizeUpload struct {
	ProductID int    `json:"productid"`
	Type      string `json:"type"`
	Size      string `json:"size"`
	Stock     int    `json:"stock"`
}

func (product_size *SizeUpload) ToDomain() products.SizeDomain {
	return products.SizeDomain{
		ProductID: product_size.ProductID,
		Type:      product_size.Type,
		Size:      product_size.Size,
		Stock:     product_size.Stock,
	}
}
