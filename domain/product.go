package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                uint            `gorm:"column:id;primaryKey;auto_increment"`
	PLU               string          `gorm:"column:plu;type:varchar(100);not null"`
	Name              string          `gorm:"column:name;type:varchar(100);not null"`
	ProductCategoryID uint            `gorm:"column:product_category_id;not null"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID;references:id"`
	Active            bool            `gorm:"column:active;default:true"`
	CreatedUser       string          `gorm:"column:created_user;type:varchar(50);not null"`
	UpdatedUser       string          `gorm:"column:updated_user;type:varchar(50);not null"`
	CreatedDate       time.Time       `gorm:"column:created_date;not null"`
	UpdatedDate       time.Time       `gorm:"column:updated_date;not null"`
	DeletedAt         gorm.DeletedAt  `gorm:"index"`
}

func (m *Product) TableName() string {
	return "product"
}
