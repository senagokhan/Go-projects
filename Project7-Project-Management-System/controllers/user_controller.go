package controllers

import (
	"Project-Management-System/models"
	"Project-Management-System/services"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	user, err := services.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user)
}

func UpdateUserProfile(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	var updatedUser models.UserModel
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := services.UpdateUser(userId, updatedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(updatedUser)
}
