package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct{}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, id uint, data map[string]any) error {
	return db.Model(new(T)).
		Where("id = ?", id).
		Updates(data).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, id uint) error {
	return db.Delete(new(T), id).Error
}

func (r *Repository[T]) FindByID(db *gorm.DB, id uint) (T, error) {
	var entity T
	err := db.Where("id = ?", id).First(&entity).Error
	return entity, err
}
