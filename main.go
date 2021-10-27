package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j1mmyson/hackathon_backend/controllers"
	"github.com/j1mmyson/hackathon_backend/database"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", controllers.Hello)
	app.Get("/ping", controllers.HealthCheck)
	app.Get("/api/users", controllers.GetUsers)

	app.Listen(":3000")
}
