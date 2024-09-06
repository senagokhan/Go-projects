package controllers

import (
	"Project-Management-System/models"
	"Project-Management-System/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var user models.UserModel
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	user.PasswordHash = string(hashedPassword)

	createdUser, err := services.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func Login(c *fiber.Ctx) error {

	var loginData models.LoginRequest
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user, err := services.GetUserByEmail(loginData.Email)
	if err != nil || user.Email == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginData.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	token, err := services.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"token": token})
}

func GetProfile(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	user, err := services.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user)
}

func UpdateProfile(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	var updatedUser models.UserModel

	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if updatedUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), 10)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		updatedUser.PasswordHash = string(hashedPassword)
	}

	if err := services.UpdateUser(userId, updatedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(updatedUser)
}
