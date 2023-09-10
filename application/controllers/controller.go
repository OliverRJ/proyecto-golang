package controllers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func CalcularEdad(fechaNacimiento time.Time) int {
	fechaActual := time.Now()
	edad := fechaActual.Year() - fechaNacimiento.Year()
	if fechaNacimiento.YearDay() > fechaActual.YearDay() {
		edad--
	}
	return edad
}
func CalcularEdadActorial(fechaNacimiento time.Time) int {
	fechaActual := time.Now()
	birthdayActual := time.Date(fechaActual.Year(), fechaNacimiento.Month(), fechaNacimiento.Day(), 0, 0, 0, 0, time.UTC)
	fechaParametro := birthdayActual.AddDate(0, 6, 0)
	edad := fechaActual.Year() - fechaNacimiento.Year()
	param := fechaActual.After(fechaParametro)
	if fechaNacimiento.YearDay() > fechaActual.YearDay() {
		edad--
	}
	if param == true {
		edad++
	}
	return edad
}
