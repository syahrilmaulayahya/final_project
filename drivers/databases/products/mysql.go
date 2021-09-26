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

func (rep MysqlProductRepository) UploadProduct(ctx context.Context, productdomain products.ProductDomain) (products.ProductDomain, error) {
	var newProduct Product
	newProduct.Code = productdomain.Code
	newProduct.Name = productdomain.Name
	newProduct.Price = productdomain.Price
	newProduct.Picture_url = productdomain.Picture_url
	newProduct.Product_typeID = productdomain.Product_typeID
	result := rep.Conn.Create(&newProduct)
	if result.Error != nil {
		return products.ProductDomain{}, result.Error
	}
	return newProduct.ToDomain(), nil
}

func (rep MysqlProductRepository) UploadSize(ctx context.Context, sizedomain products.SizeDomain) (products.SizeDomain, error) {
	var newSize Size

	newSize.ProductID = sizedomain.ProductID
	newSize.Type = sizedomain.Type
	newSize.Size = sizedomain.Size
	newSize.Stock = sizedomain.Stock

	result := rep.Conn.Create(&newSize)
	if result.Error != nil {
		return products.SizeDomain{}, result.Error
	}
	return products.SizeDomain(newSize), nil

}
func (rep MysqlProductRepository) UpdateStock(ctx context.Context, stock, id int) (products.SizeDomain, error) {
	var size Size
	size.Stock = stock
	result := rep.Conn.First(&size, "id = ?", id).Table("sizes").Where("id= ?", id).Updates(map[string]interface{}{"stock": size.Stock})
	if result.Error != nil {
		return products.SizeDomain{}, result.Error
	}
	return products.SizeDomain(size), nil
}

func (rep MysqlProductRepository) UpdateProduct(ctx context.Context, domain products.ProductDomain, id int) (products.ProductDomain, error) {
	var product Product
	result := rep.Conn.First(&product, "id = ?", id)
	product.Code = domain.Code
	product.Name = domain.Name
	product.Price = domain.Price
	product.Picture_url = domain.Picture_url
	result.Save(&product)
	if result.Error != nil {
		return products.ProductDomain{}, result.Error
	}
	return product.ToDomain(), nil

}