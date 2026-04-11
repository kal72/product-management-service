package converter

import (
	"product-management-service/internal/entity"
	"product-management-service/internal/model"
)

func ProductToResponse(product *entity.Product) *model.ProductResponse {
	return &model.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
	}
}

func ProductDetailToResponse(product *entity.ProductDetail) *model.ProductResponse {
	return &model.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Price:        product.Price,
		Stock:        product.Stock,
		CategoryName: product.CategoryName,
	}
}

func ProductListToResponse(products []entity.ProductDetail) []model.ProductResponse {
	var responses []model.ProductResponse
	for _, product := range products {
		responses = append(responses, model.ProductResponse{
			ID:           product.ID,
			Name:         product.Name,
			Price:        product.Price,
			Stock:        product.Stock,
			CategoryName: product.CategoryName,
		})
	}
	return responses
}
