package controllers

import (
	"github.com/gofiber/fiber/v2"
	"proyecto-golang/application/dto"
	"proyecto-golang/application/entities"
	"proyecto-golang/pkg/database"
)

func CreateModule(c *fiber.Ctx) error {
	var data entities.Module
	if err := c.BodyParser(&data); err != nil {
		resp := dto.NewResponse(0, dto.Error, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	if err := database.Instance.Create(&data).Error; err != nil {
		resp := dto.NewResponse(0, dto.Error, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	resp := dto.NewResponse(0, dto.Message, "Success create module", nil)
	return c.Status(fiber.StatusCreated).JSON(resp)
}
