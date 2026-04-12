package repository

import (
	"fmt"
	"product-management-service/internal/entity"
	"product-management-service/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository[entity.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetByIDWithCategory(db *gorm.DB, id int) (entity.ProductDetail, error) {
	var product entity.ProductDetail
	err := db.Model(&entity.Product{}).
		Select("products.id, products.name, products.price, products.stock, products.created_at, categories.id as category_id, categories.name as category_name").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("products.id = ?", id).
		First(&product).Error
	return product, err
}

func (r *ProductRepository) FindWithFilter(db *gorm.DB, limit, offset int, filter model.ProductFilter) ([]entity.ProductDetail, int64, error) {
	var products []entity.ProductDetail
	var total int64

	query := db.Model(&entity.Product{}).
		Select("products.id, products.name, products.price, products.stock, products.created_at, categories.id as category_id, categories.name as category_name")

	if filter.Search != "" {
		query = query.Where("products.name LIKE ? OR products.id = ?", "%"+filter.Search+"%", filter.Search)
	}

	if filter.CategoryID > 0 {
		query = query.Where("products.category_id = ?", filter.CategoryID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if filter.SortBy != "" {
		query = query.Order(fmt.Sprintf("%s %s", filter.SortBy, filter.SortOrder))
	}

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	err = query.Joins("JOIN categories ON categories.id = products.category_id").Find(&products).Error
	return products, total, err
}
