package model

import (
	"api-projeto-padrao/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func ModelGetPerson(c *fiber.Ctx) error {
	body := domain.Person{
		Id:    855525455,
		Name:  "Name",
		Age:   56,
		Email: "siau@gmail.com",
	}

	return c.JSON(body)
}
