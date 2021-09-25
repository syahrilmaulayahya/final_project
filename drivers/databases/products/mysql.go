package product

import (
	"context"
	"final_project/business/products"

	"gorm.io/gorm"
)

type MysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) products.Repository {
	return &MysqlProductRepository{
		Conn: conn,
	}
}

func (rep MysqlProductRepository) Get(ctx context.Context) ([]products.ProductDomain, error) {
	var product []Product
	result := rep.Conn.Preload("Review_Rating").Preload("Product_description").Preload("Product_type").Preload("Size").Find(&product)

	if result.Error != nil {
		return nil, result.Error
	}
	return ToListDomain(product), nil
}

func (rep MysqlProductRepository) UploadType(ctx context.Context, domain products.Product_typeDomain) (products.Product_typeDomain, error) {
	var newProductType Product_type
	newProductType.Name = domain.Name
	result := rep.Conn.Create(&newProductType)
	if result.Error != nil {
		return products.Product_typeDomain{}, result.Error
	}
	return products.Product_typeDomain(newProductType), nil

}
