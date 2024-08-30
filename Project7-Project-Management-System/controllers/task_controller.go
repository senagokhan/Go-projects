package controllers

import (
	"Project-Management-System/models"
	"Project-Management-System/services"
	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	var task models.TaskModel
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	createdTask, err := services.CreateTask(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(createdTask)
}

func GetTasks(c *fiber.Ctx) error {
	projectId := c.Query("project_id")
	tasks, err := services.GetTasks(projectId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(tasks)
}

func GetTaskById(c *fiber.Ctx) error {
	id := c.Params("task_id")
	tasks, err := services.GetTaskById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Task not found")
	}
	return c.JSON(tasks)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("task_id")
	var updatedTask models.TaskModel
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := services.updateTask(id, updatedTask); err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Task not found")
	}
	return c.JSON(updatedTask)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("task_id")
	if err := services.DeleteTask(id); err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Task not found")
	}
	return c.SendStatus(fiber.StatusNoContent)
}
