package http

import (
	"superindo-test/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductCategoryHandler struct {
	usecase domain.ProductCategoryUseCase
}

func NewProductCategoryHandler(c *fiber.App, uc domain.ProductCategoryUseCase) {
	handler := &ProductCategoryHandler{
		usecase: uc,
	}

	c.Get("/product-categories", handler.FindAllCategory)
	c.Post("/product-categories", handler.AddProductCategory)
}

func (h *ProductCategoryHandler) FindAllCategory(c *fiber.Ctx) error {
	res, err := h.usecase.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to get data",
			"data":    res,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success get data",
		"data":    res,
	})
}

func (h *ProductCategoryHandler) AddProductCategory(c *fiber.Ctx) error {
	dto := new(domain.CreatedProductCategoryDto)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(dto)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	res, err := h.usecase.Add(dto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to add",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":    "success to add",
		"is_created": res,
	})
}
