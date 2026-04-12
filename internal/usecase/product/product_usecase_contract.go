package product

import (
	"context"

	"product-management-service/internal/model"
)

type ProductUsecaseContract interface {
	Create(ctx context.Context, req *model.CreateProductRequest) (*model.ProductResponse, *model.ErrorData)
	Update(ctx context.Context, id int, req *model.UpdateProductRequest) *model.ErrorData
	Delete(ctx context.Context, id int) *model.ErrorData
	GetDetailByID(ctx context.Context, id int) (*model.ProductResponse, *model.ErrorData)
	List(ctx context.Context, req *model.ProductFilter) ([]model.ProductResponse, *model.PageMetadata, *model.ErrorData)
}
