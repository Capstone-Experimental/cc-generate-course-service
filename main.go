package main

import (
	"cc-generate-course-service/middleware"
	routes "cc-generate-course-service/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Use(middleware.JWTProtected())

	routes.InitRoutes(app)

	app.Listen(":8080")
}
