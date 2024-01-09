package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"column:id;primaryKey;auto_increment"`
	Username  string         `gorm:"type:varchar(50);unique;not null"`
	Password  string         `gorm:"type:longtext;not null"`
	Role      string         `gorm:"type:varchar(20);not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) TableName() string {
	return "user"
}

type UserCreateDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRepository interface {
	Store(entity *User) error
	FindByUsername(username string) (User, error)
}

type UserUseCase interface {
	Store(dto *UserCreateDto) (bool, error)
	Login(dto *UserCreateDto) (User, error)
}
