package routes

import (
	"cc-generate-course-service/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	routes := app.Group("/api/v1/course")

	routes.Post("/generate", handler.GenerateCourseHandler) // Generate Course
}
