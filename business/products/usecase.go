package products

import (
	"context"
	"errors"
	"time"
)

type ProductUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewProductUseCase(repo Repository, timeOut time.Duration) UseCase {
	return &ProductUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
	}
}

func (uc *ProductUseCase) Get(ctx context.Context) ([]ProductDomain, error) {
	product, err := uc.Repo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (uc *ProductUseCase) UploadProduct(ctx context.Context, productdomain ProductDomain) (ProductDomain, error) {
	if productdomain.Code == "" {
		return ProductDomain{}, errors.New("product code is empty")
	}
	if productdomain.Name == "" {
		return ProductDomain{}, errors.New("product name is empty")
	}
	if productdomain.Price == 0 {
		return ProductDomain{}, errors.New("price is empty")
	}
	if productdomain.Price < 0 {
		return ProductDomain{}, errors.New("invalid price")
	}
	if productdomain.Picture_url == "" {
		return ProductDomain{}, errors.New("picture is empty")
	}
	if productdomain.Product_typeID == 0 {
		return ProductDomain{}, errors.New("product type id is empty")
	}
	product, err := uc.Repo.UploadProduct(ctx, productdomain)
	if err != nil {
		return ProductDomain{}, err
	}
	return product, nil
}
func (uc *ProductUseCase) UpdateProduct(ctx context.Context, domain ProductDomain, id int) (ProductDomain, error) {
	if domain.Price < 0 {
		return ProductDomain{}, errors.New("invalid price")
	}
	updateProduct, err := uc.Repo.UpdateProduct(ctx, domain, id)
	if err != nil {
		return ProductDomain{}, nil
	}
	return updateProduct, nil
}

func (uc *ProductUseCase) UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error) {
	if domain.Name == "" {
		return Product_typeDomain{}, errors.New("product type name is empty")
	}
	productType, err := uc.Repo.UploadType(ctx, domain)
	if err != nil {
		return Product_typeDomain{}, err
	}
	return productType, nil
}

func (uc *ProductUseCase) UploadSize(ctx context.Context, sizedomain SizeDomain) (SizeDomain, error) {
	if sizedomain.ProductID == 0 {
		return SizeDomain{}, errors.New("product id is empty")
	}
	if sizedomain.Type == "" {
		return SizeDomain{}, errors.New("size type is empty")
	}
	if sizedomain.Size == "" {
		return SizeDomain{}, errors.New("size is empty")
	}
	if sizedomain.Stock < 0 {
		return SizeDomain{}, errors.New("stock is empty")
	}
	size, err := uc.Repo.UploadSize(ctx, sizedomain)
	if err != nil {
		return SizeDomain{}, err
	}
	return size, nil
}
func (uc *ProductUseCase) UpdateSize(ctx context.Context, sizedomain SizeDomain, id int) (SizeDomain, error) {
	if sizedomain.Type == "" {
		return SizeDomain{}, errors.New("size type is empty")
	}
	if sizedomain.Size == "" {
		return SizeDomain{}, errors.New("size is empty")
	}
	updateSize, err := uc.Repo.UpdateSize(ctx, sizedomain, id)
	if err != nil {
		return SizeDomain{}, err
	}
	return updateSize, nil
}
func (uc *ProductUseCase) UpdateStock(ctx context.Context, stock, id int) (SizeDomain, error) {
	if stock < 0 {
		return SizeDomain{}, errors.New("invalid quantity")
	}
	updateStock, err := uc.Repo.UpdateStock(ctx, stock, id)
	if err != nil {
		return SizeDomain{}, err
	}
	return updateStock, nil
}

func (uc *ProductUseCase) UploadDescription(ctx context.Context, domain Product_descriptionDomain) (Product_descriptionDomain, error) {
	if domain.Description == "" {
		return Product_descriptionDomain{}, errors.New("product description is empty")
	}
	uploadDescription, err := uc.Repo.UploadDescription(ctx, domain)
	if err != nil {
		return Product_descriptionDomain{}, nil
	}
	return uploadDescription, nil
}

func (uc *ProductUseCase) UpdateDescription(ctx context.Context, domain Product_descriptionDomain, id int) (Product_descriptionDomain, error) {
	updateDescription, err := uc.Repo.UpdateDescription(ctx, domain, id)
	if err != nil {
		return Product_descriptionDomain{}, nil
	}
	return updateDescription, nil
}
