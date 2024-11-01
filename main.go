package main

import(
"github.com/gofiber/fiber/v2"
"backend/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.Test)

	app.Listen(":3000")
}