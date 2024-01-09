package usecase

import (
	"errors"
	"superindo-test/domain"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repository domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) domain.UserUseCase {
	return &userUseCase{r}
}

func (uc *userUseCase) Store(dto *domain.UserCreateDto) (bool, error) {
	//enkripsi password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.MinCost)
	if err != nil {
		return false, err
	}

	user := domain.User{
		Username: dto.Username,
		Password: string(passwordHash),
		Role:     "Customer",
	}

	err = uc.repository.Store(&user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uc *userUseCase) Login(dto *domain.UserCreateDto) (domain.User, error) {
	res, err := uc.repository.FindByUsername(dto.Username)
	if err != nil {
		return res, err
	}

	// cek jika user tidak ada
	if res.ID == 0 {
		return res, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(dto.Password))
	if err != nil {
		return domain.User{}, errors.New("incorrect password")
	}

	return res, nil
}
