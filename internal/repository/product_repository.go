package repository

import (
	"product-management-service/internal/entity"

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
		Select("products.id, products.name, products.price, products.stock, categories.id as category_id, categories.name as category_name").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("products.id = ?", id).
		First(&product).Error
	return product, err
}
