package main

import (
	"api-projeto-padrao/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", controller.ControllerGetPerson)

	app.Listen(":3000")
}
