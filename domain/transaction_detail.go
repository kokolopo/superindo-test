package domain

import (
	"time"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	ID               uint           `gorm:"column:id;primaryKey;auto_increment"`
	TransactionID    uint           `gorm:"column:transaction_id;not null"`
	Transaction      Transaction    `gorm:"foreignKey:TransactionID;references:id"`
	ProductVariantID uint           `gorm:"column:product_variant_id;not null"`
	ProductVariant   ProductVariant `gorm:"foreignKey:ProductVariantID;references:id"`
	Price            float64        `gorm:"column:price;type:float;not null"`
	QTY              uint           `gorm:"column:qty;type:int;not null"`
	Subtotal         int64          `gorm:"column:subtotal;type:int;not null"`
	Active           bool           `gorm:"column:active;not null"`
	CreatedUser      string         `gorm:"column:created_user;not null"`
	CreatedDate      time.Time      `gorm:"column:created_date;not null"`
	UpdatedUser      string         `gorm:"column:updated_user;not null"`
	UpdatedDate      time.Time      `gorm:"column:updated_date;not null"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (m *TransactionDetail) TableName() string {
	return "transaction_detail"
}
