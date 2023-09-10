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

func CalcularEdadRelativa(fechaNacimiento time.Time) bool {
	if fechaNacimiento.Month() >= 06 {
		return true
	}
	return false
}

func CalcularEdadRelativaDos(fechaNacimiento time.Time) int {
	fechaActual := time.Now()
	fechaParametro := fechaActual.AddDate(0, 6, 0)
	birthdayActual := time.Date(fechaActual.Year(), fechaNacimiento.Month(), fechaNacimiento.Day(), 0, 0, 0, 0, time.UTC)
	edad := fechaActual.Year() - fechaNacimiento.Year()
	param := fechaParametro.After(birthdayActual)
	if fechaNacimiento.YearDay() > fechaActual.YearDay() {
		edad--
	} else {
		param = false
	}
	if param == true {
		edad++
	}
	return edad
}
