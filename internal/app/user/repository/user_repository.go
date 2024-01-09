package repository

import (
	"superindo-test/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Store(u *domain.User) error {
	if err := r.DB.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User

	err := r.DB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
