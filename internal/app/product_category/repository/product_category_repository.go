package repository

import (
	"log"
	"superindo-test/domain"

	"gorm.io/gorm"
)

type productCategoryRepository struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) domain.ProductCategoryRepository {
	return &productCategoryRepository{db}
}

func (r *productCategoryRepository) FindAll() ([]domain.ProductCategory, error) {
	var data []domain.ProductCategory
	err := r.db.Order("id desc").Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *productCategoryRepository) Store(entity *domain.ProductCategory) (bool, error) {
	log.Println(entity)
	if err := r.db.Create(&entity).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *productCategoryRepository) Update(u *domain.ProductCategory) (bool, error) {
	err := r.db.Save(&u).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
