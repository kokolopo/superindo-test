package domain

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID          uint           `gorm:"column:id;primaryKey;auto_increment"`
	Name        string         `gorm:"column:name;type:varchar(50);not null"`
	Active      bool           `gorm:"not null"`
	CreatedUser string         `gorm:"column:created_user;type:varchar(50);not null"`
	UpdatedUser string         `gorm:"column:updated_user;type:varchar(50);not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *ProductCategory) TableName() string {
	return "product_category"
}

type CreatedProductCategoryDto struct {
	Name   string `json:"name" validate:"required"`
	Active bool   `json:"active"`
}

type ProductCategoryRepository interface {
	Store(entity *ProductCategory) (bool, error)
	FindAll() ([]ProductCategory, error)
}

type ProductCategoryUseCase interface {
	Add(*CreatedProductCategoryDto) (bool, error)
	FindAll() ([]ProductCategory, error)
}
