package controllers

import (
	"Project-Management-System/models"
	"Project-Management-System/services"
	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	var project models.ProjectModel
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	createdProject, err := services.CreateProject(project)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(createdProject)
}

func GetProjects(c *fiber.Ctx) error {
	projects, err := services.GetProjects()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(projects)
}

func GetProjectById(c *fiber.Ctx) error {
	id := c.Params("project_id")
	project, err := services.GetProjectById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Project not found")
	}
	return c.JSON(project)
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("project_id")
	var updatedProject models.Project
	if err := c.BodyParser(&updatedProject); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := services.UpdateProject(id, updatedProject); err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Project not found")
	}
	return c.JSON(updatedProject)
}

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("project_id")
	if err := services.DeleteProject(id); err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Project not found")
	}
	return c.SendStatus(fiber.StatusNoContent)
}
