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
		edad2 := controllers.CalcularEdadActorial(fechaNacimiento)

		respuesta := fmt.Sprintf("Edad actual: %v años\n", edad)
		respuesta += fmt.Sprintf("Edad actorial: %v años", edad2)

		return c.SendString(respuesta)
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
