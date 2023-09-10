package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"proyecto-golang/application/entities"
	"strconv"
)

func Prueba(c *fiber.Ctx) error {
	params := c.AllParams()
	var cliente entities.Cliente
	age, err := strconv.ParseInt(params["edad"], 10, 32)
	if err != nil {
		log.Println("Error al convertir la edad:", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)

	}
	cliente.Nombre = "oliver"
	cliente.Edad = age

	//response := "Hello, Prueba, edad: " + strconv.FormatInt(age, 10)
	//return c.SendString("Hello, Prueba, edad: " + strconv.FormatInt(age, 10))
	// edad real de la persona, la fecha en formato "dd/MM/yyyy"
	// edad actuarial: edad, donde se suma un año más si ya pasaron 6 meses de cumpleaños
	//01/01/2003 -- 20
	//edad actuarial = 21

	return c.Status(fiber.StatusOK).JSON(cliente)
}
