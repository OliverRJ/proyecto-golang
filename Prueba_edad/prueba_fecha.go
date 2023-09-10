package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"proyecto-golang/application/controllers"
	"time"
)

func main() {
	app := fiber.New()

	app.Post("/obtenerFecha", func(c *fiber.Ctx) error {
		fechaRecuperada := c.FormValue("fecha")
		fechaNacimiento, err := time.Parse("02-01-2006", fechaRecuperada)
		if err != nil {
			return c.SendString("Fecha de nacimiento no válida")
		}

		edad := controllers.CalcularEdad(fechaNacimiento)
		agregarEdad := controllers.CalcularEdadRelativa(fechaNacimiento)
		edad2 := controllers.CalcularEdadRelativaDos(fechaNacimiento)

		respuesta := fmt.Sprintf("Edad actual: %v años\n", edad)
		if agregarEdad {
			edad++
		}
		respuesta += fmt.Sprintf("Edad relativa: %v años\n", edad)
		respuesta += fmt.Sprintf("Edad relativa 2: %v años", edad2)

		return c.SendString(respuesta)
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
