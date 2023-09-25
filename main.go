package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"proyecto-golang/pkg/database"
	"proyecto-golang/pkg/helpers/configloader"
	"proyecto-golang/pkg/router"
)

func main() {
	app := Setup()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	conf := configloader.ReadConfig()
	database.Connect(conf) /*
		database.Migrate()
	*/
	app.Listen(port)
}

func Setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	router.Setup(app)

	return app
}
