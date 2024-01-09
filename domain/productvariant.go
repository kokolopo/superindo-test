package domain

import (
	"time"

	"gorm.io/gorm"
)

type ProductVariant struct {
	ID          uint           `gorm:"column:id;primaryKey;auto_increment"`
	ProductID   uint           `gorm:"column:product_id;not null"`
	Product     Product        `gorm:"foreignKey:ProductID;references:id"`
	Code        string         `gorm:"column:code;type:varchar(100);not null"`
	Name        string         `gorm:"column:name;type:varchar(100);not null"`
	QTY         uint           `gorm:"column:qty;type:int;not null"`
	Price       float64        `gorm:"column:price;type:float;not null"`
	Active      bool           `gorm:"column:active;not null"`
	CreatedUser string         `gorm:"column:created_user;type:varchar(50);not null"`
	CreatedDate time.Time      `gorm:"column:created_date;not null"`
	UpdatedUser string         `gorm:"column:updated_user;type:varchar(50);not null"`
	UpdatedDate time.Time      `gorm:"column:updated_date;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (m *ProductVariant) TableName() string {
	return "product_variant"
}
