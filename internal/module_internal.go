package internal

import (
	"superindo-test/config"
	productCategoryHandler "superindo-test/internal/app/product_category/delivery/http"
	productCategoryRepository "superindo-test/internal/app/product_category/repository"
	productCategoryUseCase "superindo-test/internal/app/product_category/usecase"
	userHandler "superindo-test/internal/app/user/delivery/http"
	userRepository "superindo-test/internal/app/user/repository"
	userUseCase "superindo-test/internal/app/user/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitializeModule(db *gorm.DB, c *fiber.App, cfg config.Config) {
	// user
	ur := userRepository.NewUserRepository(db)
	uuc := userUseCase.NewUserUseCase(ur)
	userHandler.NewUserhandler(c, uuc)

	// product_category
	pcr := productCategoryRepository.NewProductCategoryRepository(db)
	pcuc := productCategoryUseCase.NewProductCategoryUseCase(pcr)
	productCategoryHandler.NewProductCategoryHandler(c, pcuc)
}
