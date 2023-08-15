package controller

import (
	"api-projeto-padrao/internal/domain"
	"api-projeto-padrao/internal/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func verifyStruct(d domain.Person) (bool, string) {
	if d.Age == 0 {
		err := "idade não está preenchido"
		log.Println(err)
		return false, err
	} else if d.Email == "" {
		err := "email não está preenchido"
		log.Println(err)
		return false, err
	} else if d.Name == "" {
		err := "nome não está preenchido"
		log.Println(err)
		return false, err
	}

	return true, ""

}

func ControllerGetPerson(c *fiber.Ctx) error {
	body := new(domain.Person)

	bodyReq := c.BodyParser(body)

	if bodyReq != nil {
		log.Println("a struct de ctx não foi convertida")
	}

	filledStruct, err := verifyStruct(*body)

	if filledStruct == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return model.ModelGetPerson(c)
}
