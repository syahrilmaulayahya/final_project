package products

type ProductUpload struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Total_Stock    uint    `json:"total_stock"`
	Price          float64 `json:"price"`
	Picture_url    string  `json:"picture_url"`
	Product_typeID uint    `json:"product_typeid"`
}

type Product_descriptionUpload struct {
	Description string `json:"description"`
}

type Product_typeUpload struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type SizeUpdate struct {
}
