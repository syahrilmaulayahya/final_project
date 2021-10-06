package products_test

import (
	"context"
	"final_project/business/products"
	"final_project/business/products/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository mocks.Repository
var productService products.UseCase
var productDomain products.ProductDomain
var reviewDomain products.Review_RatingDomain
var productDescriptionDomain products.Product_descriptionDomain
var productTypeDomain products.Product_typeDomain
var sizeDomain products.SizeDomain

func setup() {
	productService = products.NewProductUseCase(&productRepository, time.Hour*1)
	productDomain = products.ProductDomain{
		ID:                  1,
		Code:                "btk-01",
		Name:                "Gamis Biru",
		Price:               50000,
		Picture_url:         "www.google.com",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		Product_typeID:      1,
		Product_type:        productTypeDomain,
		Product_description: productDescriptionDomain,
		Review_Rating:       []products.Review_RatingDomain{},
		Size:                []products.SizeDomain{},
	}
	reviewDomain = products.Review_RatingDomain{
		ID:        1,
		Review:    "Bagus",
		Rating:    4.5,
		UserID:    1,
		ProductID: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	productDescriptionDomain = products.Product_descriptionDomain{
		ProductID:   1,
		Description: "Nyaman Dipakai",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	productTypeDomain = products.Product_typeDomain{
		ID:        1,
		Name:      "Gamis",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	sizeDomain = products.SizeDomain{
		ID:        1,
		ProductID: 1,
		Type:      "Dewasa",
		Size:      "XL",
		Stock:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestGet(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Get", func(t *testing.T) {
		productRepository.On("Get",
			mock.Anything).Return([]products.ProductDomain{}, nil).Once()
		_, err := productService.Get(context.Background())
		assert.Nil(t, err)

	})
}
func TestDetails(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Details", func(t *testing.T) {
		productRepository.On("Details",
			mock.Anything,
			mock.AnythingOfType("int")).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.Details(context.Background(), 1)
		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Empty ID", func(t *testing.T) {
		productRepository.On("Details",
			mock.Anything,
			mock.AnythingOfType("int")).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.Details(context.Background(), 0)
		assert.NotNil(t, err)

	})
}

func TestSearch(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Search", func(t *testing.T) {
		productRepository.On("Search",
			mock.Anything,
			mock.AnythingOfType("string")).Return([]products.ProductDomain{}, nil).Once()
		_, err := productService.Search(context.Background(), "Biru")
		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Empty Search", func(t *testing.T) {
		productRepository.On("Search",
			mock.Anything,
			mock.AnythingOfType("string")).Return([]products.ProductDomain{}, nil).Once()
		_, err := productService.Search(context.Background(), "")
		assert.Nil(t, err)

	})

}

func TestFilterByType(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid FilterByType", func(t *testing.T) {
		productRepository.On("FilterByType",
			mock.Anything,
			mock.AnythingOfType("int")).Return([]products.ProductDomain{}, nil).Once()
		_, err := productService.FilterByType(context.Background(), 1)
		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Empty ID", func(t *testing.T) {
		productRepository.On("FilterByType",
			mock.Anything,
			mock.AnythingOfType("int")).Return([]products.ProductDomain{}, nil).Once()
		_, err := productService.FilterByType(context.Background(), 0)
		assert.NotNil(t, err)

	})
}

func TestUploadProduct(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UploadProduct", func(t *testing.T) {
		productRepository.On("UploadProduct",
			mock.Anything,
			mock.Anything).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UploadProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "btk-01",
			Name:                "Gamis Biru",
			Price:               50000,
			Picture_url:         "www.google.com",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      1,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		})
		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Empty Code", func(t *testing.T) {
		productRepository.On("UploadProduct",
			mock.Anything,
			mock.Anything).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UploadProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "",
			Name:                "Gamis Biru",
			Price:               50000,
			Picture_url:         "www.google.com",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      1,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		})
		assert.NotNil(t, err)

	})
	t.Run("Test Case 3 | Empty Name", func(t *testing.T) {
		productRepository.On("UploadProduct",
			mock.Anything,
			mock.Anything).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UploadProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "btk-01",
			Name:                "",
			Price:               50000,
			Picture_url:         "www.google.com",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      1,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		})
		assert.NotNil(t, err)

	})
	t.Run("Test Case 4 | Empty Price", func(t *testing.T) {
		productRepository.On("UploadProduct",
			mock.Anything,
			mock.Anything).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UploadProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "btk-01",
			Name:                "Gamis Biru",
			Price:               0,
			Picture_url:         "www.google.com",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      1,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		})
		assert.NotNil(t, err)

	})
	t.Run("Test Case 5 | invalidPrice", func(t *testing.T) {
		productRepository.On("UploadProduct",
			mock.Anything,
			mock.Anything).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UploadProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "",
			Name:                "Gamis Biru",
			Price:               -50000,
			Picture_url:         "www.google.com",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      1,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		})
		assert.NotNil(t, err)

	})
	t.Run("Test Case 6 | Empty PictureUrl", func(t *testing.T) {
		productRepository.On("UploadProduct",
			mock.Anything,
			mock.Anything).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UploadProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "btk-01",
			Name:                "Gamis Biru",
			Price:               50000,
			Picture_url:         "",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      1,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		})
		assert.NotNil(t, err)

	})
	t.Run("Test Case 7 | Empty product typeid", func(t *testing.T) {
		productRepository.On("UploadProduct",
			mock.Anything,
			mock.Anything).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UploadProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "btk-01",
			Name:                "Gamis Biru",
			Price:               50000,
			Picture_url:         "www.google.com",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      0,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		})
		assert.NotNil(t, err)

	})
}

func TestUpdateProduct(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UpdateProduct", func(t *testing.T) {
		productRepository.On("UpdateProduct",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("int")).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UpdateProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "btk-01",
			Name:                "Gamis Biru",
			Price:               50000,
			Picture_url:         "www.google.com",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      1,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		}, 1)
		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Invalid Price", func(t *testing.T) {
		productRepository.On("UpdateProduct",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("int")).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UpdateProduct(context.Background(), products.ProductDomain{
			ID:                  1,
			Code:                "btk-01",
			Name:                "Gamis Biru",
			Price:               -50000,
			Picture_url:         "www.google.com",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Product_typeID:      1,
			Product_type:        productTypeDomain,
			Product_description: productDescriptionDomain,
			Review_Rating:       []products.Review_RatingDomain{},
			Size:                []products.SizeDomain{},
		}, 1)
		assert.NotNil(t, err)

	})
}

func TestUploadType(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UploadType", func(t *testing.T) {
		productRepository.On("UploadType",
			mock.Anything,
			mock.Anything).Return(products.Product_typeDomain{}, nil).Once()
		_, err := productService.UploadType(context.Background(), products.Product_typeDomain{
			ID:        1,
			Name:      "Gamis",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Empty Name", func(t *testing.T) {
		productRepository.On("UploadType",
			mock.Anything,
			mock.Anything).Return(products.Product_typeDomain{}, nil).Once()
		_, err := productService.UploadType(context.Background(), products.Product_typeDomain{
			ID:        1,
			Name:      "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		assert.NotNil(t, err)

	})
}

func TestUploadSize(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UploadSize", func(t *testing.T) {
		productRepository.On("UploadSize",
			mock.Anything,
			mock.Anything).Return(products.SizeDomain{}, nil).Once()
		_, err := productService.UploadSize(context.Background(), products.SizeDomain{
			ID:        1,
			ProductID: 1,
			Type:      "Dewasa",
			Size:      "XL",
			Stock:     100,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		assert.Nil(t, err)
		assert.Equal(t, "Dewasa", sizeDomain.Type)

	})
	t.Run("Test Case 2 | Invalid Type", func(t *testing.T) {
		productRepository.On("UploadSize",
			mock.Anything,
			mock.Anything).Return(products.SizeDomain{}, nil).Once()
		_, err := productService.UploadSize(context.Background(), products.SizeDomain{
			ID:        1,
			ProductID: 1,
			Type:      "Adult",
			Size:      "XL",
			Stock:     100,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		assert.NotNil(t, err)

	})

}
func TestUpdateSize(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UpdateSize", func(t *testing.T) {
		productRepository.On("UpdateSize",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("int")).Return(products.SizeDomain{}, nil).Once()
		_, err := productService.UpdateSize(context.Background(), products.SizeDomain{
			ID:        1,
			ProductID: 1,
			Type:      "Dewasa",
			Size:      "XL",
			Stock:     100,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, 1)
		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Invalid Type", func(t *testing.T) {
		productRepository.On("UpdateSize",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("int")).Return(products.SizeDomain{}, nil).Once()
		_, err := productService.UpdateSize(context.Background(), products.SizeDomain{
			ID:        1,
			ProductID: 1,
			Type:      "Adult",
			Size:      "XL",
			Stock:     100,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, 1)
		assert.NotNil(t, err)
	})

}
func TestUpdateStock(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UpdateStock", func(t *testing.T) {
		productRepository.On("UpdateStock",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(products.SizeDomain{}, nil).Once()
		_, err := productService.UpdateStock(context.Background(), 100, 1)
		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Invalid Stock", func(t *testing.T) {
		productRepository.On("UpdateStock",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(products.SizeDomain{}, nil).Once()
		_, err := productService.UpdateStock(context.Background(), -100, 1)
		assert.NotNil(t, err)
	})

}
func TestUploadDescription(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UploadDescription", func(t *testing.T) {
		productRepository.On("UploadDescription",
			mock.Anything,
			mock.Anything).Return(products.Product_descriptionDomain{}, nil).Once()
		_, err := productService.UploadDescription(context.Background(), products.Product_descriptionDomain{
			ProductID:   1,
			Description: "Nyaman Dipakai",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Invalid UploadDescription", func(t *testing.T) {
		productRepository.On("UploadDescription",
			mock.Anything,
			mock.Anything).Return(products.Product_descriptionDomain{}, nil).Once()
		_, err := productService.UploadDescription(context.Background(), products.Product_descriptionDomain{
			ProductID:   1,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		assert.NotNil(t, err)

	})

}
func TestUpdateDescription(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UpdateDescription", func(t *testing.T) {
		productRepository.On("UpdateDescription",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("int")).Return(products.Product_descriptionDomain{}, nil).Once()
		_, err := productService.UpdateDescription(context.Background(), products.Product_descriptionDomain{
			ProductID:   1,
			Description: "Bagus",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}, 1)
		assert.Nil(t, err)

	})
}
