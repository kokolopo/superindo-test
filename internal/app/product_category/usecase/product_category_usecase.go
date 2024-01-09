package usecase

import (
	"log"
	"superindo-test/domain"
)

type productCategoryUseCase struct {
	repository domain.ProductCategoryRepository
}

func NewProductCategoryUseCase(r domain.ProductCategoryRepository) domain.ProductCategoryUseCase {
	return &productCategoryUseCase{r}
}

func (uc *productCategoryUseCase) FindAll() ([]domain.ProductCategory, error) {
	res, err := uc.repository.FindAll()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *productCategoryUseCase) Add(dto *domain.CreatedProductCategoryDto) (bool, error) {
	productCategory := domain.ProductCategory{
		Name:        dto.Name,
		Active:      dto.Active,
		CreatedUser: "OPERATOR",
		UpdatedUser: "OPERATOR",
	}

	log.Println(dto.Active)

	res, err := uc.repository.Store(&productCategory)
	if err != nil {
		return res, err
	}

	return res, nil
}
