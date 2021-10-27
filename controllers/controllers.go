package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j1mmyson/hackathon_backend/database"
	"github.com/j1mmyson/hackathon_backend/model"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("hello hackathon!@#")
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong!"})
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	if err := db.Find(&users).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": users})
}
