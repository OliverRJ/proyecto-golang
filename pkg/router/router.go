package router

import (
	"github.com/gofiber/fiber/v2"
	"proyecto-golang/application/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("/v1")
	api.Get("/", controllers.Hello)

	pruebaGroup := api.Group("prueba")
	pruebaGroup.Get("/calculo/:edad", controllers.Prueba)

	module := api.Group("/modules")
	module.Post("/", controllers.CreateModule)
}
