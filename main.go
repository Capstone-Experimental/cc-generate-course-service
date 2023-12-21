package main

import (
	"cc-generate-course-service/middleware"
	routes "cc-generate-course-service/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// Auth Middleware
	// app.Use(middleware.JWTProtected())
	app.Use(middleware.FirebaseAuth())

	routes.InitRoutes(app)

	app.Listen(":8080")
}
