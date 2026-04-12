package product

import (
	"context"
	"fmt"
	"time"

	"product-management-service/internal/entity"
	"product-management-service/internal/model"
	"product-management-service/internal/model/converter"
	"product-management-service/internal/repository"
	"product-management-service/internal/utils/errorhandler"
	"product-management-service/internal/utils/pagination"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductUsecase struct {
	db          *gorm.DB
	productRepo *repository.ProductRepository
	redisRepo   *repository.RedisRepository
	validator   *validator.Validate
}

func NewProductUsecase(
	db *gorm.DB,
	productRepo *repository.ProductRepository,
	redisRepo *repository.RedisRepository,
	validator *validator.Validate,
) ProductUsecaseContract {
	return &ProductUsecase{
		db:          db,
		productRepo: productRepo,
		redisRepo:   redisRepo,
		validator:   validator,
	}
}

func (uc *ProductUsecase) Create(ctx context.Context, req *model.CreateProductRequest) (*model.ProductResponse, *model.ErrorData) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, errorhandler.ErrorInvalidRequest(err)
	}

	product := &entity.Product{
		Name:       req.Name,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}

	if err := uc.productRepo.Create(uc.db, product); err != nil {
		return nil, errorhandler.ErrorDB(err)
	}

	_ = uc.redisRepo.DeleteByPrefix(ctx, "product:list:")

	return converter.ProductToResponse(product), nil
}

func (uc *ProductUsecase) Update(ctx context.Context, id int, req *model.UpdateProductRequest) *model.ErrorData {
	if err := uc.validator.Struct(req); err != nil {
		return errorhandler.ErrorInvalidRequest(err)
	}

	if _, err := uc.productRepo.FindByID(uc.db, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errorhandler.ErrorNotFound(err)
		}
		return errorhandler.ErrorDB(err)
	}

	updateData := map[string]interface{}{
		"name":        req.Name,
		"price":       req.Price,
		"stock":       req.Stock,
		"category_id": req.CategoryID,
	}

	if err := uc.productRepo.Update(uc.db, id, updateData); err != nil {
		return errorhandler.ErrorDB(err)
	}

	_ = uc.redisRepo.DeleteByPrefix(ctx, "product:list:")

	return nil
}

func (uc *ProductUsecase) Delete(ctx context.Context, id int) *model.ErrorData {
	if _, err := uc.productRepo.FindByID(uc.db, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errorhandler.ErrorNotFound(err)
		}
		return errorhandler.ErrorDB(err)
	}

	if err := uc.productRepo.Delete(uc.db, id); err != nil {
		return errorhandler.ErrorDB(err)
	}

	_ = uc.redisRepo.DeleteByPrefix(ctx, "product:list:")

	return nil
}

func (uc *ProductUsecase) GetDetailByID(ctx context.Context, id int) (*model.ProductResponse, *model.ErrorData) {
	product, err := uc.productRepo.GetByIDWithCategory(uc.db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorhandler.ErrorNotFound(err)
		}
		return nil, errorhandler.ErrorDB(err)
	}

	return converter.ProductDetailToResponse(&product), nil
}

func (uc *ProductUsecase) List(ctx context.Context, req *model.ProductFilter) ([]model.ProductResponse, *model.PageMetadata, *model.ErrorData) {
	cacheKey := fmt.Sprintf("product:list:s_%s:c_%d:p_%d:sz_%d:sb_%s:so_%s", req.Search, req.CategoryID, req.Page, req.Size, req.SortBy, req.SortOrder)
	cacheData := model.CacheData[[]model.ProductResponse]{}

	if err := uc.redisRepo.Get(ctx, cacheKey, &cacheData); err == nil {
		return cacheData.Data, cacheData.Pages, nil
	}

	limit, offset := pagination.CalculateLimitOffset(req.Page, req.Size)

	products, total, err := uc.productRepo.FindWithFilter(uc.db, limit, offset, *req)
	if err != nil {
		return nil, nil, errorhandler.ErrorDB(err)
	}

	pages := pagination.CalculatePage(total, req.Size, req.Page)
	productResp := converter.ProductListToResponse(products)

	cacheData = model.CacheData[[]model.ProductResponse]{
		Data:  productResp,
		Pages: pages,
	}

	uc.redisRepo.Set(ctx, cacheKey, cacheData, 5*time.Minute)

	return productResp, pages, nil
}
