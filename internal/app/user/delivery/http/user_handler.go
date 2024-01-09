package http

import (
	"superindo-test/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	useCase domain.UserUseCase
}

func NewUserhandler(c *fiber.App, us domain.UserUseCase) {
	handler := &UserHandler{
		useCase: us,
	}

	c.Post("/user", handler.UserRegister)
	c.Post("/login", handler.UserLogin)

}

func (h *UserHandler) UserRegister(c *fiber.Ctx) error {
	dto := new(domain.UserCreateDto)
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

	res, err := h.useCase.Store(dto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message":    err,
			"is_created": res,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":    "success to created user",
		"is_created": res,
	})
}

func (h *UserHandler) UserLogin(c *fiber.Ctx) error {
	dto := new(domain.UserCreateDto)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	validate := validator.New()
	errValidate := validate.Struct(dto)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	res, err := h.useCase.Login(dto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"data":    res,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success to sign in",
		"data":    res,
	})
}
