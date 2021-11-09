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

// function to create a new user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

// function to get a user by id
func GetUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	if err := db.First(&user, c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

// function to update a user by id
func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	if err := db.First(&user, c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := db.Save(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

// function to delete a user by id
func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	if err := db.First(&user, c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := db.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

// function to create a new article
func CreateArticle(c *fiber.Ctx) error {
	db := database.DB
	article := new(model.Article)

	if err := c.BodyParser(article); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := db.Create(&article).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": article})
}

// function to get all articles
func GetArticles(c *fiber.Ctx) error {
	db := database.DB
	var articles []model.Article

	if err := db.Find(&articles).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": articles})
}

// function to get an article by id
func GetArticle(c *fiber.Ctx) error {
	db := database.DB
	article := new(model.Article)

	if err := db.First(&article, c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": article})
}

// function to update an article by id
func UpdateArticle(c *fiber.Ctx) error {
	db := database.DB
	article := new(model.Article)

	if err := db.First(&article, c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := c.BodyParser(article); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := db.Save(&article).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": article})
}

// function to delete an article by id
func DeleteArticle(c *fiber.Ctx) error {
	db := database.DB
	article := new(model.Article)

	if err := db.First(&article, c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := db.Delete(&article).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": article})
}

// function to find articles by tag
func GetArticlesByTag(c *fiber.Ctx) error {
	db := database.DB
	var articles []model.Article

	if err := db.Where("tag = ?", c.Params("tag")).Find(&articles).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": articles})
}

// function to find articles by author
func GetArticlesByAuthor(c *fiber.Ctx) error {
	db := database.DB
	var articles []model.Article

	if err := db.Where("author = ?", c.Params("author")).Find(&articles).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": articles})
}
