package domain

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint           `gorm:"column:id;primaryKey;auto_increment"`
	TransactionNo string         `gorm:"column:transaction_no;type:varchar(100);not null"`
	TotalAmount   int64          `gorm:"column:total_amount;type:int;not null"`
	Active        bool           `gorm:"column:active;not null"`
	CreatedUser   string         `gorm:"column:created_user;type:varchar(50);not null"`
	CreatedDate   time.Time      `gorm:"column:created_date;not null"`
	UpdatedUser   string         `gorm:"column:updated_user;type:varchar(50);not null"`
	UpdatedDate   time.Time      `gorm:"column:updated_date;not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (m *Transaction) TableName() string {
	return "transaction"
}
